package handlers

import (
    "github.com/TaofeekOlusola/pkg/models"
    "fmt"
)

func PrintUser(user models.User) {
    fmt.Println("User:", user.Name, "Email:", user.Email)
}