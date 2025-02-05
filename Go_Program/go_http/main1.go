package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Welcome to the home page! (GET)")
	}else if r.Method == "POST" {
		fmt.Fprintf(w, "Welcome to the home page! (POST)")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)

    fmt.Println("Server running on port 9090...")
    err := http.ListenAndServe(":9090", nil)
    if err!= nil {
        panic(err)
    }
}