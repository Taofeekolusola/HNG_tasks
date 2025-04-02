// package auth

// import (
// 	"context"
// 	"net/http"
// 	"strconv"

// 	"my-gqlgen-app/internal/users" 
// 	"my-gqlgen-app/internal/pkg/jwt" 
// )

// var userCtxKey = &contextKey{"user"}

// type contextKey struct {
// 	name string
// }

// func Middleware() func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			header := r.Header.Get("Authorization")

// 			// Allow unauthenticated users in
// 			if header == "" {
// 				next.ServeHTTP(w, r)
// 				return
// 			}

// 			//validate jwt token
// 			tokenStr := header
// 			username, err := jwt.ParseToken(tokenStr)
// 			if err != nil {
// 				http.Error(w, "Invalid token", http.StatusForbidden)
// 				return
// 			}

// 			// create user and check if user exists in db
// 			user := users.User{Username: username}
// 			id, err := users.GetUserIdByUsername(username)
// 			if err != nil {
// 				next.ServeHTTP(w, r)
// 				return
// 			}
// 			user.ID = strconv.Itoa(id)
// 			// put it in context
// 			ctx := context.WithValue(r.Context(), userCtxKey, &user)

// 			// and call the next with our new context
// 			r = r.WithContext(ctx)
// 			next.ServeHTTP(w, r)
// 		})
// 	}
// }

// // ForContext finds the user from the context. REQUIRES Middleware to have run.
// func ForContext(ctx context.Context) *users.User {
// 	raw, _ := ctx.Value(userCtxKey).(*users.User)
// 	return raw
// }

package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	_ "strconv" 

	"github.com/dgrijalva/jwt-go"
	"my-gqlgen-app/internal/users"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        fmt.Println("Auth Header:", authHeader) // Debugging

        if authHeader == "" {
            fmt.Println("No Authorization header found")
            next.ServeHTTP(w, r)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            fmt.Println("Invalid token format")
            http.Error(w, "Invalid authorization token format", http.StatusUnauthorized)
            return
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Ensure token uses the correct signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte("secretCode"), nil // Replace with your actual secret
        })

        if err != nil {
            fmt.Println("JWT Parse Error:", err)
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        if !token.Valid {
            fmt.Println("Invalid token")
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            fmt.Println("Invalid token claims")
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }

        username, ok := claims["username"].(string)
        if !ok || username == "" {
            fmt.Println("Username not found in token claims")
            http.Error(w, "Invalid token payload", http.StatusUnauthorized)
            return
        }

        fmt.Println("Authenticated User:", username) // Debugging

        ctx := context.WithValue(r.Context(), userCtxKey, &users.User{Username: username})
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// ForContext retrieves the user from the request context
func ForContext(ctx context.Context) *users.User {
	user, ok := ctx.Value(userCtxKey).(*users.User) // Retrieve as *users.User
	if !ok {
		return nil
	}
	return user
}

// func Middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		authHeader := r.Header.Get("Authorization")
// 		fmt.Println("Auth Header:", authHeader) // Debugging

// 		if authHeader == "" {
// 			fmt.Println("No Authorization header found")
// 			next.ServeHTTP(w, r)
// 			return
// 		}

// 		// Ensure token starts with "Bearer "
// 		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
// 		if tokenString == authHeader {
// 			fmt.Println("Invalid token format")
// 			http.Error(w, "Invalid authorization token format", http.StatusUnauthorized)
// 			return
// 		}

// 		// Parse the JWT token
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// Change "YOUR_SECRET_KEY" to your actual JWT secret
// 			return []byte("secretCode"), nil
// 		})

// 		if err != nil {
// 			fmt.Println("JWT Parse Error:", err)
// 			http.Error(w, "Invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		if !token.Valid {
// 			fmt.Println("Invalid token")
// 			http.Error(w, "Invalid token", http.StatusUnauthorized)
// 			return
// 		}

// 		// Extract user claims
// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok {
// 			fmt.Println("Invalid token claims")
// 			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
// 			return
// 		}

// 		fmt.Printf("User Claims: %+v\n", claims) // Debugging

// 		// Store user in context
// 		ctx := context.WithValue(r.Context(), userCtxKey, claims)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// func ForContext(ctx context.Context) interface{} {
// 	raw, _ := ctx.Value(userCtxKey).(jwt.MapClaims)
// 	return raw
// }
