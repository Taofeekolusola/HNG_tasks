package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    done := make(chan string)

    go func() {
        time.Sleep(2 * time.Second) // Simulate long task
        done <- "Hello, World!" // Send response to channel
    }()

    fmt.Fprintln(w, <-done) // Wait for response from Goroutine
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
