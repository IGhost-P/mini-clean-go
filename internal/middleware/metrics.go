package middleware

import (
    "net/http"
    "strconv"
    "time"

    "github.com/IGhost-p/mini-clean-go/internal/metrics"
)

func MetricsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // 응답 캡처를 위한 래퍼
        wrapped := wrapResponseWriter(w)
        
        // 요청 처리
        next(wrapped, r)
        
        // 메트릭 기록
        duration := time.Since(start).Seconds()
        status := strconv.Itoa(wrapped.status)
        
        metrics.HttpRequestsTotal.WithLabelValues(r.Method, r.URL.Path, status).Inc()
        metrics.HttpRequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
    }
}

type responseWriter struct {
    http.ResponseWriter
    status int
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
    return &responseWriter{w, http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.status = code
    rw.ResponseWriter.WriteHeader(code)
}