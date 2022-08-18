package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Add Middleware Logic Here
		start := time.Now()
		log.Println("Started", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		log.Println("Completed in", time.Since(start))
	})
}
