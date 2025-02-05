package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Product struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Price float64 `json:"price"`
}

var products = []Product{
    {ID: 1, Name: "Laptop", Price: 999.99},
    {ID: 2, Name: "Phone", Price: 499.99},
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
    // Set response content type to JSON
    w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Encode products and sends JSON
    json.NewEncoder(w).Encode(products)
}

func addProductHandler(w http.ResponseWriter, r *http.Request) {
    // Check if the request method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Decode the JSON body
    var newProduct Product
    err := json.NewDecoder(r.Body).Decode(&newProduct)
    if err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    // Add the new product to the slice
    products = append(products, newProduct)

    // Respond with the updated list of products
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

func main() {
    http.HandleFunc("/products", getProductsHandler)
    http.HandleFunc("/add-product", addProductHandler)

    fmt.Println("Server is running on http://localhost:9090")
    http.ListenAndServe(":9090", nil)
}