package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong!")
}

func main() {
	http.HandleFunc("/ping", pingHandler)
    fmt.Println("Server listening on port 9090...")
    err := http.ListenAndServe(":9090", nil)
    if err!= nil {
        panic(err)
    }
}