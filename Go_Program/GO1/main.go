package main

import (
    "github.com/TaofeekOlusola/pkg/handlers"
    "github.com/TaofeekOlusola/pkg/models"
)

func main() {
    user := models.User{Name: "John Doe", Email: "john@example.com"}
    handlers.PrintUser(user)
}