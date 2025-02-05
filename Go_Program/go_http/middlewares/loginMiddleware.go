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
        next.ServeHTTP(w, r) // Pass control to the next handler
        log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
    })
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    // Wrap the helloHandler with the logging middleware
    http.Handle("/", loggingMiddleware(http.HandlerFunc(helloHandler)))

    // Start the server
    log.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}