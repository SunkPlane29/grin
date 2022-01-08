package util

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/SunkPlane29/grin/pkg/auth/token"
)

type ContextKey string

var IDK ContextKey = "user-id"

func AuthMiddleware(tokenIssuer token.JWT, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("missing authorization header"))
			return
		}

		token := strings.Split(auth, " ")
		if token[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("wrong authentication token type, required: Bearer"))
			return
		}

		if len(token) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid authorization token"))
			return
		}

		id, err := tokenIssuer.Validate(token[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token signature"))
			return
		}

		ctx := context.WithValue(r.Context(), IDK, id)

		f(w, r.WithContext(ctx))
	}
}

func CORSMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") //TODO: hardcoded origin
		f(w, r)
	}
}

type (
	responseData struct {
		status int
		size   int
	}

	loggingResponseWriter struct {
		http.ResponseWriter
		data *responseData
	}
)

func (lw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lw.ResponseWriter.Write(b)
	lw.data.size += size
	return size, err
}

func (lw *loggingResponseWriter) WriteHeader(statusCode int) {
	lw.ResponseWriter.WriteHeader(statusCode)
	lw.data.status = statusCode
}

func LoggerMiddleware(prefix string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		uri := r.URL.Path
		method := r.Method

		responseData := &responseData{
			status: 0,
			size:   0,
		}

		lw := &loggingResponseWriter{
			ResponseWriter: w,
			data:           responseData,
		}

		h.ServeHTTP(lw, r)

		duration := time.Since(start)

		logger := log.New(os.Stdout, prefix, log.Ldate|log.Ltime)
		logger.Printf("[%s] %s, status: %d, size: %d, elapsed: %dms\n",
			method,
			uri,
			responseData.status,
			responseData.size,
			duration.Round(time.Millisecond).Milliseconds(),
		)
	})
}

func RecoverMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				panic(err)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
