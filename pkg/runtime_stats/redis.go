package runtime_stats

import (
	"context"
	"encoding/json"
	"sync"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Enabled       bool   `yaml:"enabled"`
	Addr          string `yaml:"addr"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	DB            int    `yaml:"db"`
	KeyPrefix     string `yaml:"key_prefix"`
	FlushInterval int    `yaml:"flush_interval"`
	QueueSize     int    `yaml:"queue_size"`
	LogSize       int    `yaml:"log_size"`
	RetentionDays int    `yaml:"retention_days"`
}

type RedisStatus struct {
	Enabled       bool      `json:"enabled"`
	QueueLen      int       `json:"queue_len"`
	Dropped       uint64    `json:"dropped"`
	LastWriteAt   time.Time `json:"last_write_at,omitempty"`
	LastError     string    `json:"last_error,omitempty"`
	LastErrorAt   time.Time `json:"last_error_at,omitempty"`
	FlushInterval int       `json:"flush_interval"`
}

type RedisSink struct {
	cfg    RedisConfig
	client *redis.Client
	queue  chan QueryRecord
	done   chan struct{}
	wg     sync.WaitGroup

	dropped atomic.Uint64
	mu      sync.RWMutex
	status  RedisStatus
}

func NewRedisSink(cfg RedisConfig) *RedisSink {
	if cfg.Addr == "" {
		cfg.Addr = "127.0.0.1:6379"
	}
	if cfg.KeyPrefix == "" {
		cfg.KeyPrefix = "mosdns:"
	}
	if cfg.FlushInterval <= 0 {
		cfg.FlushInterval = 5
	}
	if cfg.QueueSize <= 0 {
		cfg.QueueSize = 4096
	}
	if cfg.LogSize <= 0 {
		cfg.LogSize = 2000
	}
	if cfg.RetentionDays <= 0 {
		cfg.RetentionDays = 30
	}
	s := &RedisSink{
		cfg: cfg,
		client: redis.NewClient(&redis.Options{
			Addr:     cfg.Addr,
			Username: cfg.Username,
			Password: cfg.Password,
			DB:       cfg.DB,
		}),
		queue: make(chan QueryRecord, cfg.QueueSize),
		done:  make(chan struct{}),
		status: RedisStatus{
			Enabled:       true,
			FlushInterval: cfg.FlushInterval,
		},
	}
	s.wg.Add(1)
	go s.run()
	return s
}

func (s *RedisSink) EnqueueQuery(r QueryRecord) {
	select {
	case s.queue <- r:
	default:
		s.dropped.Add(1)
	}
}

func (s *RedisSink) Status() RedisStatus {
	s.mu.RLock()
	st := s.status
	s.mu.RUnlock()
	st.QueueLen = len(s.queue)
	st.Dropped = s.dropped.Load()
	return st
}

func (s *RedisSink) Close() error {
	close(s.done)
	s.wg.Wait()
	return s.client.Close()
}

func (s *RedisSink) run() {
	defer s.wg.Done()
	ticker := time.NewTicker(time.Duration(s.cfg.FlushInterval) * time.Second)
	defer ticker.Stop()
	batch := make([]QueryRecord, 0, 256)
	flush := func() {
		if len(batch) == 0 {
			return
		}
		if err := s.flush(batch); err != nil {
			s.setError(err)
		} else {
			s.mu.Lock()
			s.status.LastWriteAt = time.Now()
			s.status.LastError = ""
			s.mu.Unlock()
		}
		batch = batch[:0]
	}
	for {
		select {
		case r := <-s.queue:
			batch = append(batch, r)
			if len(batch) >= 256 {
				flush()
			}
		case <-ticker.C:
			flush()
		case <-s.done:
			flush()
			return
		}
	}
}

func (s *RedisSink) flush(batch []QueryRecord) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pipe := s.client.Pipeline()
	for _, r := range batch {
		prefix := s.cfg.KeyPrefix
		pipe.Incr(ctx, prefix+"queries:total")
		if r.Client != "" {
			pipe.HIncrBy(ctx, prefix+"clients:"+r.Time.Format("20060102"), r.Client, 1)
		}
		if r.QName != "" {
			pipe.HIncrBy(ctx, prefix+"domains:"+r.Time.Format("20060102"), r.QName, 1)
		}
		pipe.HIncrBy(ctx, prefix+"stats:hourly:"+r.Time.Format("2006010215"), "total", 1)
		pipe.HIncrBy(ctx, prefix+"stats:daily:"+r.Time.Format("20060102"), "total", 1)
		b, _ := json.Marshal(r)
		pipe.LPush(ctx, prefix+"query_logs", string(b))
		pipe.LTrim(ctx, prefix+"query_logs", 0, int64(s.cfg.LogSize-1))
	}
	if s.cfg.RetentionDays > 0 {
		exp := time.Duration(s.cfg.RetentionDays) * 24 * time.Hour
		for _, r := range batch {
			prefix := s.cfg.KeyPrefix
			pipe.Expire(ctx, prefix+"clients:"+r.Time.Format("20060102"), exp)
			pipe.Expire(ctx, prefix+"domains:"+r.Time.Format("20060102"), exp)
			pipe.Expire(ctx, prefix+"stats:hourly:"+r.Time.Format("2006010215"), exp)
			pipe.Expire(ctx, prefix+"stats:daily:"+r.Time.Format("20060102"), exp)
		}
	}
	_, err := pipe.Exec(ctx)
	return err
}

func (s *RedisSink) setError(err error) {
	s.mu.Lock()
	s.status.LastError = err.Error()
	s.status.LastErrorAt = time.Now()
	s.mu.Unlock()
}
