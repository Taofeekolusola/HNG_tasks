package auth

import (
	"errors"
	"time"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("123adeshina12?") // Change this to a strong secret key

// GenerateToken generates a JWT token for a given user ID.
func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken validates a JWT token and returns the user ID.
func ValidateToken(tokenString string) (int, error) {

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("invalid signing method")
        }
        return secretKey, nil
    })

    if err != nil {
        fmt.Println("❌ JWT Parsing Error:", err)
        return 0, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

        userID, ok := claims["id"].(float64)
        if !ok {
            return 0, errors.New("invalid token payload")
        }

        return int(userID), nil
    }

    return 0, errors.New("invalid token")
}