package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// Middleware to validate the request body
func validationMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        var user User
        err := json.NewDecoder(r.Body).Decode(&user)
        if err != nil {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        if user.Username == "" || user.Password == "" {
            http.Error(w, "Username and password are required", http.StatusBadRequest)
            return
        }
        next.ServeHTTP(w, r) // Pass control to the next handler
    })
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    fmt.Fprintf(w, "Registered user: %s\n", user.Username)
}

func main() {
    // Wrap the registerHandler with the validationMiddleware
    http.Handle("/register", validationMiddleware(http.HandlerFunc(registerHandler)))

    // Start the server
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}