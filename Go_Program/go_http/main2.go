package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, World!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    username := vars["username"]
    fmt.Fprintf(w, "Hello, %s!", username)
}

func main() {
    // Create a new router
    r := mux.NewRouter()

    // Register handlers
    r.HandleFunc("/", helloHandler)
    r.HandleFunc("/user/{username}", userHandler)

    // Start the server on port 8080
    fmt.Println("Server is running on http://localhost:9091")
    http.ListenAndServe(":9091", r)
}