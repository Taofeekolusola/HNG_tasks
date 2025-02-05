package main

import "fmt"

func main() {
	var num1, num2 int
    fmt.Print("Enter the first number: ")
    fmt.Scan(&num1)
    fmt.Print("Enter the second number: ")
    fmt.Scan(&num2)

    sum := num1 + num2
    fmt.Printf("Sum: %d\n", sum)

    product := num1 * num2
    fmt.Printf("Product: %d\n", product)

    diff := num1 - num2
    fmt.Printf("Difference: %d\n", diff)

    quotient := num1 / num2
    fmt.Printf("Quotient: %d\n", quotient)

    remainder := num1 % num2
    fmt.Printf("Remainder: %d\n", remainder)
}