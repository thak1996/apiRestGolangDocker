package handlers

import (
    "log"
    "net/http"
    "time"
)

// LoggingMiddleware registra cada requisição recebida e seu tempo de resposta
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Recebendo %s %s de %s", r.Method, r.URL.Path, r.RemoteAddr)

        // Capture o status de resposta
        rw := &responseWriter{w, http.StatusOK}
        next.ServeHTTP(rw, r)

        log.Printf("Respondido %s %s com status %d em %v", r.Method, r.URL.Path, rw.statusCode, time.Since(start))
    })
}

// responseWriter é um wrapper para capturar o código de status de resposta
type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}
