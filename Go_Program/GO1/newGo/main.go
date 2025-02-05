package main

import (
    "github.com/OlusolaTaofeek/pkg/models"
    "github.com/OlusolaTaofeek/pkg/helper"
    "fmt"
)


func main() {
	// Create an instance of Employee
    emp := models.Employee{
        ID:        1,
        Name:    "John Doe",
        Email:   "john.doe@example.com",
        Age:     30,
        Position: "Software Engineer",
    }

    // Print the employee details
    helper.PrintEmp(emp)

    // Update the employee's age
    emp.Age = 31
    fmt.Println("Updated employee age:", emp.Age)

    // Print the updated employee details
    helper.PrintEmp(emp)
}