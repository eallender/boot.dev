package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type apiConfig struct {
	fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(writer, request)
	})
}

func (cfg *apiConfig) serverMetrics(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(200)
	writer.Write([]byte(fmt.Sprintf("Hits: %d", cfg.fileserverHits.Load())))
}

func (cfg *apiConfig) resetMetrics(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	cfg.fileserverHits.Store(0)
}

func healthz(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/plain; charset=utf-8")
	writer.WriteHeader(200)
	writer.Write([]byte("OK"))
}

func main() {
	serveMux := http.NewServeMux()
	apiCfg := apiConfig{fileserverHits: atomic.Int32{}}
	serveMux.Handle("/app/", http.StripPrefix("/app", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(".")))))
	serveMux.HandleFunc("GET /healthz", healthz)
	serveMux.HandleFunc("GET /metrics", apiCfg.serverMetrics)
	serveMux.HandleFunc("POST /reset", apiCfg.resetMetrics)

	server := http.Server{}
	server.Addr = ":8080"
	server.Handler = serveMux

	server.ListenAndServe()
}
