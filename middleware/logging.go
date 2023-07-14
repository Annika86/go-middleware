package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			next.ServeHTTP(w, r)

			duration := time.Since(start)
			log.Printf("Request: %s %s %s", r.Method, r.URL, &duration)
		})
	}
}
