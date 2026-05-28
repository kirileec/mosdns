package webui

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/IrineSistiana/mosdns/v5/coremain"
	"github.com/IrineSistiana/mosdns/v5/pkg/query_context"
	"github.com/IrineSistiana/mosdns/v5/pkg/runtime_stats"
	"github.com/IrineSistiana/mosdns/v5/plugin/executable/sequence"
	"github.com/go-chi/chi/v5"
	"github.com/miekg/dns"
)

const PluginType = "webui"

func init() {
	coremain.RegNewPluginFunc(PluginType, Init, func() any { return new(Args) })
}

type Args struct {
	StaticDir string                    `yaml:"static_dir"`
	IndexFile string                    `yaml:"index_file"`
	Username  string                    `yaml:"username"`
	Password  string                    `yaml:"password"`
	LogSize   int                       `yaml:"log_size"`
	Redis     runtime_stats.RedisConfig `yaml:"redis"`
}

func (a *Args) init() {
	if a.StaticDir == "" {
		a.StaticDir = "./webui"
	}
	if a.IndexFile == "" {
		a.IndexFile = "index.html"
	}
}

type WebUI struct {
	bp            *coremain.BP
	args          *Args
	sessionSecret string
}

func Init(bp *coremain.BP, args any) (any, error) {
	a := args.(*Args)
	a.init()
	bp.M().GetRuntimeStats().EnableRedis(a.Redis)
	p := &WebUI{bp: bp, args: a, sessionSecret: newSessionSecret()}
	bp.RegAPI(p.api())
	return p, nil
}

func (p *WebUI) api() *chi.Mux {
	r := chi.NewRouter()
	r.Use(p.auth)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/session", p.handleSession)
		r.Post("/session", p.handleLogin)
		r.Delete("/session", p.handleLogout)
		r.Get("/overview", p.handleOverview)
		r.Get("/stats/hourly", p.handleHourly)
		r.Get("/stats/daily", p.handleDaily)
		r.Get("/stats/clients", p.handleClients)
		r.Get("/stats/domains/top", p.handleDomains)
		r.Get("/query-log", p.handleQueryLogs)
		r.Get("/query-log/recent", p.handleRecentLogs)
		r.Get("/query-log/stream", p.handleLogStream)
		r.Get("/cache", p.handleCache)
		r.Post("/cache/refresh", p.handleRefreshCache)
		r.Get("/upstreams", p.handleUpstreams)
		r.Get("/config", p.handleGetConfig)
		r.Put("/config", p.handlePutConfig)
		r.Post("/config/validate", p.handleValidateConfig)
		r.Post("/config/reload", p.handleReload)
		r.Get("/storage/status", p.handleStorageStatus)
	})
	r.NotFound(p.serveStatic)
	r.Get("/*", p.serveStatic)
	r.Get("/", p.serveStatic)
	return r
}

func (p *WebUI) auth(next http.Handler) http.Handler {
	if p.args.Username == "" && p.args.Password == "" {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if !isAPIRequest(req) || isSessionEndpoint(req) {
			next.ServeHTTP(w, req)
			return
		}
		if !p.validSession(req) {
			writeJSONStatus(w, http.StatusUnauthorized, map[string]any{"ok": false, "error": "unauthorized"})
			return
		}
		next.ServeHTTP(w, req)
	})
}

func isAPIRequest(req *http.Request) bool {
	return strings.Contains(req.URL.Path, "/api/") || strings.HasPrefix(req.URL.Path, "/api/")
}

func isSessionEndpoint(req *http.Request) bool {
	path := strings.TrimSuffix(req.URL.Path, "/")
	return path == "/api/v1/session" || strings.HasSuffix(path, "/api/v1/session")
}

func (p *WebUI) handleSession(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, map[string]any{"authenticated": p.validSession(req)})
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p *WebUI) handleLogin(w http.ResponseWriter, req *http.Request) {
	var in loginReq
	if err := json.NewDecoder(io.LimitReader(req.Body, 1<<20)).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if subtle.ConstantTimeCompare([]byte(in.Username), []byte(p.args.Username)) != 1 || subtle.ConstantTimeCompare([]byte(in.Password), []byte(p.args.Password)) != 1 {
		writeJSONStatus(w, http.StatusUnauthorized, map[string]any{"ok": false, "error": "invalid credentials"})
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "mosdns_webui_session",
		Value:    p.sessionValue(),
		Path:     "/",
		MaxAge:   int((12 * time.Hour).Seconds()),
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	})
	writeJSON(w, map[string]any{"ok": true})
}

func (p *WebUI) handleLogout(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "mosdns_webui_session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	})
	writeJSON(w, map[string]any{"ok": true})
}

func (p *WebUI) validSession(req *http.Request) bool {
	if p.args.Username == "" && p.args.Password == "" {
		return true
	}
	cookie, err := req.Cookie("mosdns_webui_session")
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(cookie.Value), []byte(p.sessionValue())) == 1
}

func (p *WebUI) sessionValue() string {
	sum := sha256.Sum256([]byte(p.args.Username + ":" + p.args.Password + ":" + p.sessionSecret))
	return hex.EncodeToString(sum[:])
}

func newSessionSecret() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(b)
}

func (p *WebUI) handleOverview(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, p.bp.M().GetRuntimeStats().Overview())
}

func (p *WebUI) handleHourly(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, p.bp.M().GetRuntimeStats().Snapshot(topN(req)).Hourly)
}

func (p *WebUI) handleDaily(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, p.bp.M().GetRuntimeStats().Snapshot(topN(req)).Daily)
}

func (p *WebUI) handleClients(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, p.bp.M().GetRuntimeStats().Snapshot(topN(req)).Clients)
}

func (p *WebUI) handleDomains(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, p.bp.M().GetRuntimeStats().Snapshot(topN(req)).Domains)
}

func (p *WebUI) handleRecentLogs(w http.ResponseWriter, req *http.Request) {
	limit, _ := strconv.Atoi(req.URL.Query().Get("limit"))
	writeJSON(w, p.bp.M().GetRuntimeStats().RecentLogs(limit))
}

func (p *WebUI) handleQueryLogs(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))
	logQuery := runtime_stats.LogQuery{
		Limit:    limit,
		Offset:   offset,
		Search:   query.Get("search"),
		Protocol: query.Get("protocol"),
	}
	if rcode := query.Get("rcode"); rcode != "" {
		v, err := strconv.Atoi(rcode)
		if err != nil {
			http.Error(w, "invalid rcode", http.StatusBadRequest)
			return
		}
		logQuery.RCode = &v
	}
	writeJSON(w, p.bp.M().GetRuntimeStats().QueryLogs(logQuery))
}

func (p *WebUI) handleLogStream(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}
	ch, unsubscribe := p.bp.M().GetRuntimeStats().Subscribe()
	defer unsubscribe()
	for {
		select {
		case r := <-ch:
			b, err := runtime_stats.EncodeSSE(r)
			if err == nil {
				_, _ = w.Write(b)
				flusher.Flush()
			}
		case <-req.Context().Done():
			return
		}
	}
}

func (p *WebUI) handleCache(w http.ResponseWriter, req *http.Request) {
	out := make([]runtime_stats.CacheStats, 0)
	for _, plugin := range p.bp.M().GetPlugins() {
		if provider, ok := plugin.(runtime_stats.CacheStatsProvider); ok {
			out = append(out, provider.CacheStats())
		}
	}
	writeJSON(w, out)
}

type refreshCacheReq struct {
	Domain string `json:"domain"`
	QType  string `json:"qtype"`
	Entry  string `json:"entry"`
}

type refreshCacheResult struct {
	QType   string   `json:"qtype"`
	RCode   int      `json:"rcode"`
	Answers []string `json:"answers,omitempty"`
	TTLs    []uint32 `json:"ttls,omitempty"`
	Error   string   `json:"error,omitempty"`
}

func (p *WebUI) handleRefreshCache(w http.ResponseWriter, req *http.Request) {
	var in refreshCacheReq
	if err := json.NewDecoder(io.LimitReader(req.Body, 1<<20)).Decode(&in); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	domain := strings.TrimSpace(in.Domain)
	if domain == "" {
		http.Error(w, "domain is required", http.StatusBadRequest)
		return
	}
	domain = dns.Fqdn(domain)
	entryTag := strings.TrimSpace(in.Entry)
	if entryTag == "" {
		entryTag = "main_sequence"
	}
	entry, ok := p.bp.M().GetPlugin(entryTag).(sequence.Executable)
	if !ok {
		http.Error(w, "entry executable not found: "+entryTag, http.StatusBadRequest)
		return
	}

	qtypes, err := refreshQTypes(in.QType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	refreshers := make([]runtime_stats.CacheRefresher, 0)
	for _, plugin := range p.bp.M().GetPlugins() {
		if refresher, ok := plugin.(runtime_stats.CacheRefresher); ok {
			refreshers = append(refreshers, refresher)
		}
	}
	if len(refreshers) == 0 {
		http.Error(w, "cache plugin not found", http.StatusBadRequest)
		return
	}

	results := make([]refreshCacheResult, 0, len(qtypes))
	removed := 0
	for _, qtype := range qtypes {
		q := new(dns.Msg)
		q.SetQuestion(domain, qtype)
		for _, refresher := range refreshers {
			removed += refresher.RemoveCache(q.Copy())
		}

		qCtx := query_context.NewContext(q)
		ctx, cancel := context.WithTimeout(req.Context(), 8*time.Second)
		err := entry.Exec(ctx, qCtx)
		cancel()

		result := refreshCacheResult{QType: dns.TypeToString[qtype], RCode: dns.RcodeServerFailure}
		if result.QType == "" {
			result.QType = strconv.Itoa(int(qtype))
		}
		if resp := qCtx.R(); resp != nil {
			result.RCode = resp.Rcode
			result.Answers, result.TTLs = summarizeAnswers(resp)
		}
		if err != nil {
			result.Error = err.Error()
		}
		results = append(results, result)
	}

	writeJSON(w, map[string]any{
		"ok":      true,
		"domain":  domain,
		"entry":   entryTag,
		"removed": removed,
		"results": results,
	})
}

func refreshQTypes(s string) ([]uint16, error) {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "", "ALL":
		return []uint16{dns.TypeA, dns.TypeAAAA}, nil
	case "A":
		return []uint16{dns.TypeA}, nil
	case "AAAA":
		return []uint16{dns.TypeAAAA}, nil
	case "HTTPS":
		return []uint16{dns.TypeHTTPS}, nil
	default:
		return nil, fmt.Errorf("unsupported qtype %q", s)
	}
}

func summarizeAnswers(resp *dns.Msg) ([]string, []uint32) {
	answers := make([]string, 0, len(resp.Answer))
	ttls := make([]uint32, 0, len(resp.Answer))
	for _, rr := range resp.Answer {
		ttls = append(ttls, rr.Header().Ttl)
		switch v := rr.(type) {
		case *dns.A:
			answers = append(answers, v.A.String())
		case *dns.AAAA:
			answers = append(answers, v.AAAA.String())
		case *dns.CNAME:
			answers = append(answers, v.Target)
		case *dns.PTR:
			answers = append(answers, v.Ptr)
		case *dns.MX:
			answers = append(answers, v.Mx)
		case *dns.NS:
			answers = append(answers, v.Ns)
		case *dns.TXT:
			answers = append(answers, strings.Join(v.Txt, " "))
		default:
			answers = append(answers, rr.String())
		}
	}
	return answers, ttls
}

func (p *WebUI) handleUpstreams(w http.ResponseWriter, req *http.Request) {
	out := make([]runtime_stats.UpstreamStats, 0)
	for _, plugin := range p.bp.M().GetPlugins() {
		if provider, ok := plugin.(runtime_stats.UpstreamStatsProvider); ok {
			out = append(out, provider.UpstreamStats()...)
		}
	}
	writeJSON(w, out)
}

func (p *WebUI) handleGetConfig(w http.ResponseWriter, req *http.Request) {
	configFile := p.bp.M().ConfigFile()
	if configFile == "" {
		http.Error(w, "config file is unknown", http.StatusNotFound)
		return
	}
	b, err := os.ReadFile(configFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = w.Write(b)
}

func (p *WebUI) handlePutConfig(w http.ResponseWriter, req *http.Request) {
	configFile := p.bp.M().ConfigFile()
	if configFile == "" {
		http.Error(w, "config file is unknown", http.StatusNotFound)
		return
	}
	b, err := io.ReadAll(io.LimitReader(req.Body, 4<<20))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if old, err := os.ReadFile(configFile); err == nil {
		_ = os.WriteFile(configFile+".bak", old, 0o600)
	}
	if err := os.WriteFile(configFile, b, 0o600); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, map[string]any{"ok": true, "backup": configFile + ".bak"})
}

func (p *WebUI) handleValidateConfig(w http.ResponseWriter, req *http.Request) {
	configFile := p.bp.M().ConfigFile()
	if configFile == "" {
		http.Error(w, "config file is unknown", http.StatusNotFound)
		return
	}
	validateFile := configFile
	if req.Body != nil {
		b, err := io.ReadAll(io.LimitReader(req.Body, 4<<20))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if len(b) > 0 {
			f, err := os.CreateTemp(filepath.Dir(configFile), ".mosdns-config-*.yaml")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			validateFile = f.Name()
			_, writeErr := f.Write(b)
			closeErr := f.Close()
			defer os.Remove(validateFile)
			if writeErr != nil {
				http.Error(w, writeErr.Error(), http.StatusInternalServerError)
				return
			}
			if closeErr != nil {
				http.Error(w, closeErr.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
	if err := coremain.ValidateConfig(validateFile); err != nil {
		writeJSONStatus(w, http.StatusBadRequest, map[string]any{"ok": false, "error": err.Error()})
		return
	}
	writeJSON(w, map[string]any{"ok": true})
}

func (p *WebUI) handleReload(w http.ResponseWriter, req *http.Request) {
	writeJSONStatus(w, http.StatusNotImplemented, map[string]any{
		"supported": false,
		"message":   "config reload is not implemented yet",
	})
}

func (p *WebUI) handleStorageStatus(w http.ResponseWriter, req *http.Request) {
	writeJSON(w, p.bp.M().GetRuntimeStats().RedisStatus())
}

func (p *WebUI) serveStatic(w http.ResponseWriter, req *http.Request) {
	path := strings.TrimPrefix(req.URL.Path, "/")
	path = strings.TrimPrefix(path, "plugins/"+p.bp.Tag()+"/")
	if path == "" || strings.HasPrefix(path, "api/") {
		path = p.args.IndexFile
	}
	fullPath := filepath.Join(p.args.StaticDir, filepath.Clean(path))
	if st, err := os.Stat(fullPath); err != nil || st.IsDir() {
		fullPath = filepath.Join(p.args.StaticDir, p.args.IndexFile)
	}
	if _, err := os.Stat(fullPath); err != nil {
		http.Error(w, fmt.Sprintf("webui static file not found: %s", fullPath), http.StatusNotFound)
		return
	}
	http.ServeFile(w, req, fullPath)
}

func writeJSON(w http.ResponseWriter, v any) {
	writeJSONStatus(w, http.StatusOK, v)
}

func writeJSONStatus(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func topN(req *http.Request) int {
	n, _ := strconv.Atoi(req.URL.Query().Get("limit"))
	if n <= 0 {
		return 10
	}
	return n
}
