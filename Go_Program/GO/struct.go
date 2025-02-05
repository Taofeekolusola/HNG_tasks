package main

import "fmt"

type Employee struct {
	name string
	age int
	salary float64
}

// Function to calculate the annual salary
func (e *Employee) calculateAnnualSalary() float64 {
    return e.salary * 12
}

func main() {
	// Create an employee object
    employee := Employee{name: "John Doe", age: 30, salary: 50000}

    // Calculate the annual salary
    annualSalary := employee.calculateAnnualSalary()

    // Print the annual salary
    fmt.Printf("%s's annual salary is $%.2f\n", employee.name, annualSalary)

	// Update the employee's salary
	employee.salary = 60000

    // Recalculate the annual salary after the update
    updatedAnnualSalary := employee.calculateAnnualSalary()

    // Print the updated annual salary
    fmt.Printf("%s's updated annual salary is $%.2f\n", employee.name, updatedAnnualSalary)
}