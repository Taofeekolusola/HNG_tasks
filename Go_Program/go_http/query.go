package main


import (
    "fmt"
    "net/http"
)

// func searchHandler(w http.ResponseWriter, r *http.Request) {
// 	// Get the query parameter from the request
//     query := r.URL.Query().Get("q")

//     // Perform the search operation using the provided query
//     // Replace this with actual search logic
//     searchResult := fmt.Sprintf("Search results for '%s': [Result 1, Result 2, Result 3]", query)

//     // Write the search results to the response
//     fmt.Fprintf(w, "%s", searchResult)
// }

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	q := query.Get("q")
	limit := r.URL.Query().Get("limit")

	fmt.Fprintf(w, "Search query: %s, Limit: %s\n", q, limit)
}

func main() {
	http.HandleFunc("/search", searchHandler)

    fmt.Println("Server is running on http://localhost:9090")
    http.ListenAndServe(":9090", nil)
}