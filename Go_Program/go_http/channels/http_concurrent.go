package main

import (
    "fmt"
    "net/http"
    "time"
   _ "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Processing in the background...") // Respond immediately

    // Start background processing in a Goroutine
    go func() {
        time.Sleep(1 * time.Second) // Simulate a long task
        fmt.Println("Background Task 1 Completed")
    }()
    go func() {
        time.Sleep(2 * time.Second) // Simulate a long task
        fmt.Println("Background Task 2 Completed")
    }()
    go func() {
        time.Sleep(3 * time.Second) // Simulate another long task
        fmt.Println("Background Task 3 Completed")
    }()
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server is running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}