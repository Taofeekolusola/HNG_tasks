package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
	Email string `json:"email"`
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Decode the JSON body
    var user User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Respond with the received data
    fmt.Fprintf(w, "Registered user: %s\n", user.Username)
	fmt.Fprintf(w, "Email: %s\n", user.Email)
}

func main() {
    http.HandleFunc("/register", registerHandler)

    fmt.Println("Server is running on http://localhost:9090")
    http.ListenAndServe(":9090", nil)
}