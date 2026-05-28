package runtime_stats

import (
	"encoding/json"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/IrineSistiana/mosdns/v5/pkg/query_context"
	"github.com/miekg/dns"
)

const (
	defaultLogSize = 2000
	secondBuckets  = 60
	hourBuckets    = 24
	dayBuckets     = 30
)

type QueryRecord struct {
	Time      time.Time `json:"time"`
	Client    string    `json:"client,omitempty"`
	QName     string    `json:"qname"`
	QType     uint16    `json:"qtype"`
	QClass    uint16    `json:"qclass"`
	RCode     int       `json:"rcode"`
	ElapsedMS int64     `json:"elapsed_ms"`
	Error     string    `json:"error,omitempty"`
}

type CountPoint struct {
	Time  time.Time `json:"time"`
	Count uint64    `json:"count"`
}

type NamedCount struct {
	Name  string `json:"name"`
	Count uint64 `json:"count"`
}

type Overview struct {
	TotalQueries uint64  `json:"total_queries"`
	QPS          float64 `json:"qps"`
	AvgLatencyMS float64 `json:"avg_latency_ms"`
}

type Snapshot struct {
	Overview Overview     `json:"overview"`
	Hourly   []CountPoint `json:"hourly"`
	Daily    []CountPoint `json:"daily"`
	Clients  []NamedCount `json:"clients"`
	Domains  []NamedCount `json:"domains"`
}

type CacheStats struct {
	Tag          string  `json:"tag"`
	QueryTotal   uint64  `json:"query_total"`
	HitTotal     uint64  `json:"hit_total"`
	LazyHitTotal uint64  `json:"lazy_hit_total"`
	Size         int     `json:"size"`
	HitRate      float64 `json:"hit_rate"`
}

type CacheStatsProvider interface {
	CacheStats() CacheStats
}

type UpstreamStats struct {
	ForwardTag    string    `json:"forward_tag"`
	Tag           string    `json:"tag"`
	Addr          string    `json:"addr"`
	Status        string    `json:"status"`
	QueryTotal    uint64    `json:"query_total"`
	SuccessTotal  uint64    `json:"success_total"`
	ErrorTotal    uint64    `json:"error_total"`
	InFlight      int64     `json:"inflight"`
	LastLatencyMS int64     `json:"last_latency_ms"`
	AvgLatencyMS  float64   `json:"avg_latency_ms"`
	SuccessRate   float64   `json:"success_rate"`
	LastSuccessAt time.Time `json:"last_success_at,omitempty"`
	LastErrorAt   time.Time `json:"last_error_at,omitempty"`
	LastError     string    `json:"last_error,omitempty"`
}

type UpstreamStatsProvider interface {
	UpstreamStats() []UpstreamStats
}

type Stats struct {
	totalQueries atomic.Uint64
	totalLatency atomic.Int64

	mu      sync.RWMutex
	seconds [secondBuckets]countBucket
	hours   [hourBuckets]countBucket
	days    [dayBuckets]countBucket
	clients map[string]uint64
	domains map[string]uint64
	logs    []QueryRecord
	logHead int
	logFull bool

	subscribers map[chan QueryRecord]struct{}
	redis       *RedisSink
}

type countBucket struct {
	stamp int64
	count uint64
}

func New(logSize int, redis RedisConfig) *Stats {
	if logSize <= 0 {
		logSize = defaultLogSize
	}
	s := &Stats{
		clients:     make(map[string]uint64),
		domains:     make(map[string]uint64),
		logs:        make([]QueryRecord, logSize),
		subscribers: make(map[chan QueryRecord]struct{}),
	}
	if redis.Enabled {
		s.redis = NewRedisSink(redis)
	}
	return s
}

func (s *Stats) EnableRedis(cfg RedisConfig) {
	if !cfg.Enabled {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.redis != nil {
		return
	}
	s.redis = NewRedisSink(cfg)
}

func (s *Stats) Close() error {
	if s.redis != nil {
		return s.redis.Close()
	}
	return nil
}

func (s *Stats) RecordQuery(qCtx *query_context.Context, resp *dns.Msg, execErr error) {
	now := time.Now()
	question := qCtx.QQuestion()
	r := QueryRecord{
		Time:      now,
		QName:     question.Name,
		QType:     question.Qtype,
		QClass:    question.Qclass,
		RCode:     dns.RcodeServerFailure,
		ElapsedMS: time.Since(qCtx.StartTime()).Milliseconds(),
	}
	if clientAddr := qCtx.ServerMeta.ClientAddr; clientAddr.IsValid() {
		r.Client = clientAddr.String()
	}
	if resp != nil {
		r.RCode = resp.Rcode
	}
	if execErr != nil {
		r.Error = execErr.Error()
	}

	s.record(r)
}

func (s *Stats) record(r QueryRecord) {
	s.totalQueries.Add(1)
	s.totalLatency.Add(r.ElapsedMS)
	sec := r.Time.Unix()
	hour := r.Time.Truncate(time.Hour).Unix()
	day := r.Time.Truncate(24 * time.Hour).Unix()

	s.mu.Lock()
	s.bumpBucket(s.seconds[:], sec)
	s.bumpBucket(s.hours[:], hour)
	s.bumpBucket(s.days[:], day)
	if r.Client != "" {
		s.clients[r.Client]++
	}
	if r.QName != "" {
		s.domains[r.QName]++
	}
	s.logs[s.logHead] = r
	s.logHead = (s.logHead + 1) % len(s.logs)
	if s.logHead == 0 {
		s.logFull = true
	}
	for ch := range s.subscribers {
		select {
		case ch <- r:
		default:
		}
	}
	s.mu.Unlock()

	if s.redis != nil {
		s.redis.EnqueueQuery(r)
	}
}

func (s *Stats) bumpBucket(buckets []countBucket, stamp int64) {
	i := int(stamp % int64(len(buckets)))
	if buckets[i].stamp != stamp {
		buckets[i] = countBucket{stamp: stamp}
	}
	buckets[i].count++
}

func (s *Stats) Overview() Overview {
	total := s.totalQueries.Load()
	avg := 0.0
	if total > 0 {
		avg = float64(s.totalLatency.Load()) / float64(total)
	}
	now := time.Now().Unix()
	var recent uint64
	s.mu.RLock()
	for _, b := range s.seconds {
		if b.stamp > now-10 && b.stamp <= now {
			recent += b.count
		}
	}
	s.mu.RUnlock()
	return Overview{TotalQueries: total, QPS: float64(recent) / 10, AvgLatencyMS: avg}
}

func (s *Stats) Snapshot(topN int) Snapshot {
	if topN <= 0 {
		topN = 10
	}
	overview := s.Overview()
	s.mu.RLock()
	defer s.mu.RUnlock()
	return Snapshot{
		Overview: overview,
		Hourly:   bucketPoints(s.hours[:], time.Hour),
		Daily:    bucketPoints(s.days[:], 24*time.Hour),
		Clients:  topCounts(s.clients, topN),
		Domains:  topCounts(s.domains, topN),
	}
}

func (s *Stats) RecentLogs(limit int) []QueryRecord {
	s.mu.RLock()
	defer s.mu.RUnlock()
	max := s.logHead
	if s.logFull {
		max = len(s.logs)
	}
	if limit <= 0 || limit > max {
		limit = max
	}
	out := make([]QueryRecord, 0, limit)
	for i := 0; i < limit; i++ {
		idx := s.logHead - 1 - i
		if idx < 0 {
			idx += len(s.logs)
		}
		out = append(out, s.logs[idx])
	}
	return out
}

func (s *Stats) Subscribe() (chan QueryRecord, func()) {
	ch := make(chan QueryRecord, 64)
	s.mu.Lock()
	s.subscribers[ch] = struct{}{}
	s.mu.Unlock()
	return ch, func() {
		s.mu.Lock()
		delete(s.subscribers, ch)
		close(ch)
		s.mu.Unlock()
	}
}

func (s *Stats) RedisStatus() RedisStatus {
	if s.redis == nil {
		return RedisStatus{Enabled: false}
	}
	return s.redis.Status()
}

func EncodeSSE(v any) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	out := make([]byte, 0, len(b)+8)
	out = append(out, "data: "...)
	out = append(out, b...)
	out = append(out, "\n\n"...)
	return out, nil
}

func bucketPoints(buckets []countBucket, step time.Duration) []CountPoint {
	points := make([]CountPoint, 0, len(buckets))
	for _, b := range buckets {
		if b.stamp == 0 {
			continue
		}
		points = append(points, CountPoint{Time: time.Unix(b.stamp, 0), Count: b.count})
	}
	sort.Slice(points, func(i, j int) bool { return points[i].Time.Before(points[j].Time) })
	return points
}

func topCounts(m map[string]uint64, n int) []NamedCount {
	out := make([]NamedCount, 0, len(m))
	for k, v := range m {
		out = append(out, NamedCount{Name: k, Count: v})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Count > out[j].Count })
	if len(out) > n {
		out = out[:n]
	}
	return out
}
