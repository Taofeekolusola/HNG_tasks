// package main

// import (
// 	"context"
// 	"log"
// 	"net/http"

// 	"github.com/99designs/gqlgen/graphql/handler"
// 	"github.com/99designs/gqlgen/graphql/playground"
// 	"github.com/TaofeekOlusola/ent-gql-todo/auth" // Import your auth package
// 	"github.com/TaofeekOlusola/ent-gql-todo/ent"
// 	"github.com/TaofeekOlusola/ent-gql-todo/graph"
// 	"github.com/TaofeekOlusola/ent-gql-todo/graph/generated"

// 	_ "github.com/lib/pq"
// )

// func main() {
// 	// PostgreSQL connection string
// 	dsn := "postgres://postgres:postgres@localhost:5444/entgql?sslmode=disable"

// 	// Connect to PostgreSQL
// 	client, err := ent.Open("postgres", dsn)
// 	if err != nil {
// 		log.Fatalf("failed to open database: %v", err)
// 	}
// 	defer client.Close()

// 	// Run database migrations
// 	ctx := context.Background()
// 	if err := client.Schema.Create(ctx); err != nil {
// 		log.Fatalf("failed to run migrations: %v", err)
// 	}

// 	// Setup GraphQL server
// 	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Client: client}}))

// 	// Use authentication middleware for GraphQL endpoint
// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", auth.Middleware(srv)) // Apply middleware here ✅

// 	log.Println("🚀 Server running on http://localhost:8080/")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
    _"github.com/lib/pq"
	"github.com/TaofeekOlusola/ent-gql-todo/ent"
	"github.com/TaofeekOlusola/ent-gql-todo/graph/generated"
	"github.com/TaofeekOlusola/ent-gql-todo/middleware" // Import your middleware package
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
    "github.com/TaofeekOlusola/ent-gql-todo/graph"
)

    func main() {
        // PostgreSQL connection string
    dsn := "postgres://postgres:postgres@localhost:5444/entgql?sslmode=disable"

    // Connect to PostgreSQL
    client, err := ent.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("failed to open database: %v", err)
    }
    defer client.Close()
	// Run database migrations
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	// Setup Gin
	r := gin.Default()

	// ✅ Apply AuthMiddleware globally
	r.Use(middleware.AuthMiddleware())

	// Setup GraphQL server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Client: client}}))

	// Use gin.WrapH to convert http.HandlerFunc to gin.HandlerFunc
	r.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	r.POST("/query", gin.WrapH(srv)) // AuthMiddleware now applies to this route

	// Start the server
	log.Println("Server is running on http://localhost:8080/")
	log.Fatal(r.Run(":8080"))
}