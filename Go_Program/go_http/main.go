package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })

    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/about", aboutHandler)

    fmt.Println("Server running on port 9090...")
    err := http.ListenAndServe(":9090", nil)
    if err!= nil {
        panic(err)
    }
}