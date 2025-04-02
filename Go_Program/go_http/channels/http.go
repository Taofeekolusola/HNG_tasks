package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Simulate a long-running task
    time.Sleep(2 * time.Second)
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    // Register the handler
    http.HandleFunc("/", handler)

    // Start the server
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}