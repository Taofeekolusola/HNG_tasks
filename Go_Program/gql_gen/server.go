// package main

// import (
// 	"log"
// 	"my-gqlgen-app/graph"
// 	"net/http"
// 	"os"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// 	"github.com/go-chi/chi/v5"
// 	 database "my-gqlgen-app/internal/pkg/db/migrations/mysql"
// )

// const defaultPort = "8080"

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	router := chi.NewRouter()

// 	database.InitDB()
// 	defer database.CloseDB()
// 	database.Migrate()

// 	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

// 	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	router.Handle("/query", server)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, router))
// }

package main

import (
	"log"
	"my-gqlgen-app/graph"
	"my-gqlgen-app/internal/auth" // Import the auth middleware
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	database "my-gqlgen-app/internal/pkg/db/migrations/mysql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	database.InitDB()
	defer database.CloseDB()
	database.Migrate()

	server := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// 🛑 WRAP /query with auth middleware!
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", auth.Middleware(server)) // Apply auth middleware

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
