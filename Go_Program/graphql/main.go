package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

// Define a struct for our "Book" type
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Sample data
var books = []Book{
	{ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan"},
	{ID: "2", Title: "Go in Action", Author: "William Kennedy"},
	{ID: "3", Title: "Effective Go", Author: "Russ Cox"},
}

// Define the GraphQL schema
var bookType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Book",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.String},
		"title":  &graphql.Field{Type: graphql.String},
		"author": &graphql.Field{Type: graphql.String},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"books": &graphql.Field{
			Type: graphql.NewList(bookType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return books, nil
			},
		},
	},
})

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
})

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %v\n", result.Errors)
	}
	return result
}

func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	var result struct {
		Query string `json:"query"`
	}
	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Execute the query
	response := executeQuery(result.Query, schema)

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Set up a GraphQL endpoint
	http.HandleFunc("/graphql", graphqlHandler)

	// Start the server
	fmt.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
