package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// Middleware to log requests
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
        log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
    })
}

// Middleware to check for a valid API key
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("X-API-Key")
        if apiKey != "secret-key" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "You have access to the secret!")
}

func main() {
    // Chain the middlewares together
    handler := loggingMiddleware(authMiddleware(http.HandlerFunc(secretHandler)))

    // Register the chained handler
    http.Handle("/secret", handler)

    // Start the server
    log.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}