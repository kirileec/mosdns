package webui

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/IrineSistiana/mosdns/v5/coremain"
	"github.com/IrineSistiana/mosdns/v5/pkg/runtime_stats"
	"github.com/go-chi/chi/v5"
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
	bp   *coremain.BP
	args *Args
}

func Init(bp *coremain.BP, args any) (any, error) {
	a := args.(*Args)
	a.init()
	bp.M().GetRuntimeStats().EnableRedis(a.Redis)
	p := &WebUI{bp: bp, args: a}
	bp.RegAPI(p.api())
	return p, nil
}

func (p *WebUI) api() *chi.Mux {
	r := chi.NewRouter()
	r.Use(p.auth)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/overview", p.handleOverview)
		r.Get("/stats/hourly", p.handleHourly)
		r.Get("/stats/daily", p.handleDaily)
		r.Get("/stats/clients", p.handleClients)
		r.Get("/stats/domains/top", p.handleDomains)
		r.Get("/query-log/recent", p.handleRecentLogs)
		r.Get("/query-log/stream", p.handleLogStream)
		r.Get("/cache", p.handleCache)
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
		u, pass, ok := req.BasicAuth()
		if !ok || u != p.args.Username || pass != p.args.Password {
			w.Header().Set("WWW-Authenticate", `Basic realm="mosdns webui"`)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, req)
	})
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
	if err := coremain.ValidateConfig(configFile); err != nil {
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
