package middleware

import (
    "net/http"
    "time"

    "github.com/IGhost-p/mini-clean-go/internal/logger"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // 요청 로깅
        logger.Info("Incoming request", map[string]interface{}{
            "method": r.Method,
            "path":   r.URL.Path,
            "remote": r.RemoteAddr,
        })

        next(w, r)

        // 응답 로깅
        logger.Info("Request completed", map[string]interface{}{
            "method":   r.Method,
            "path":     r.URL.Path,
            "duration": time.Since(start).String(),
        })
    }
}