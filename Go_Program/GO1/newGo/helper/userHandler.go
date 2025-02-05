package helper

import (
	"github.com/OlusolaTaofeek/pkg/models"
	"fmt"
)

func PrintEmp(user models.Employee) {
	fmt.Printf("Name: %s, Email: %s, ID: %d, Age: %d, Position: %s \n", user.Name, user.Email, user.ID, user.Age, user.Position)
}