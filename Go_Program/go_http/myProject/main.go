package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Serve static files from the "static" directory
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Register a handler for the root path
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html")
    })

    // Start the server on port 8080
    fmt.Println("Server is running on http://localhost:9090")
    http.ListenAndServe(":9090", nil)
}