package main

import (
	"fmt"
	"net/http"
	"os"
)

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := os.ReadFile("metrics.html")
	if err != nil {
		http.Error(w, "Could not read metrics.html", http.StatusInternalServerError)
		return
	}

	updatedContent := fmt.Sprintf(string(htmlContent), cfg.fileserverHits.Load())

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(updatedContent))
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
