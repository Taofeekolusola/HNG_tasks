package main

import (
    "fmt"
    "net/http"
    "strings"
)

// Middleware to check for a valid API key
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        apiKey := r.Header.Get("X-API-Key")
        if apiKey != "secret-key" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r) // Pass control to the next handler
    })
}

func secretHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "You have access to the secret!")
}

func main() {
    // Wrap the secretHandler with the authMiddleware
    http.Handle("/secret", authMiddleware(http.HandlerFunc(secretHandler)))

    // Start the server
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}