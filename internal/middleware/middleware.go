package middleware

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"testProject/pkg/logger"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	size   int64
}

func (rw *responseWriter) WriteHeader(status int) {
	rw.status = status
	rw.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	size, err := rw.ResponseWriter.Write(b)
	rw.size += int64(size)
	return size, err
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}

		ctx := context.WithValue(r.Context(), "requestID", generateRequestID())
		r = r.WithContext(ctx)

		defer func() {
			logger.Info("HTTP Request", map[string]interface{}{
				"method":    r.Method,
				"path":      r.URL.Path,
				"status":    rw.status,
				"size":      rw.size,
				"duration":  time.Since(start).String(),
				"requestID": ctx.Value("requestID"),
				"userAgent": r.UserAgent(),
				"remoteIP":  r.RemoteAddr,
			})
		}()

		next.ServeHTTP(rw, r)
	})
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error(fmt.Errorf("%v", err), "panic recovered", map[string]interface{}{
					"path":      r.URL.Path,
					"method":    r.Method,
					"requestID": r.Context().Value("requestID"),
				})
				http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}

func generateRequestID() string {
	return fmt.Sprintf("%d-%s", time.Now().UnixNano(),
		uuid.NewString()[:8])
}
