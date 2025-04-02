package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/TaofeekOlusola/ent-gql-todo/auth"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            return
        }

        tokenString = strings.TrimPrefix(tokenString, "Bearer ")

        userID, err := auth.ValidateToken(tokenString)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        // Store user_id in request context with a proper type
        ctx := context.WithValue(c.Request.Context(), "id", userID)
        c.Request = c.Request.WithContext(ctx)

        c.Next()
    }
}


// func AuthMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         tokenString := c.GetHeader("Authorization")

//         if tokenString == "" {
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
//             return
//         }

// 		if tokenString == "" {
//             fmt.Println("❌ Missing Token")
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
//             return
//         }

//         // Remove "Bearer " prefix if present
//         tokenString = strings.TrimPrefix(tokenString, "Bearer ")

//         userID, err := auth.ValidateToken(tokenString)
//         if err != nil {
//             c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
//             return
//         }

//         fmt.Println("✅ Middleware extracted userID:", userID) // Debugging line

//         // Store userID in context
//         ctx := context.WithValue(c.Request.Context(), "id", float64(userID))
//         c.Request = c.Request.WithContext(ctx)

//         c.Next()
//     }
// }