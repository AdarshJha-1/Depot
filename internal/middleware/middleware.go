package middleware

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// colors (feels good)
var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func SetResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")

		next.ServeHTTP(w, r)
	})
}
func LogAPIInfo(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		fmt.Printf("%sTIME:%s %s | %sMETHOD:%s %s | %sPATH:%s %s | %sDURATION:%s %v\n",
			Yellow, Reset,
			start.Format("2006-01-02 15:04:05"),
			Blue, Reset,
			r.Method,
			Cyan, Reset,
			r.URL.Path,
			Green, Reset,
			duration,
		)
	})
}

var limiter = rate.NewLimiter(1, 5)

func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !limiter.Allow() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}
