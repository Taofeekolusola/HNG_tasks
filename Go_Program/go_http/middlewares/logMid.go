package main

import (
    "fmt"
    "log"
    "net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Capture request details
        method := r.Method
        url := r.URL.String()
        timestamp := time.Now().Format(time.RFC3339)

        // Wrap the ResponseWriter to capture the status code
        lrw := &loggingResponseWriter{w, http.StatusOK}

        // Pass control to the next handler
        next.ServeHTTP(lrw, r)

        // Log the details including the status code
        log.Printf("[%s] %s %s - %d", timestamp, method, url, lrw.statusCode)
    })
}