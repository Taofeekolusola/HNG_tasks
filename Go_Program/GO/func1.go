package main
import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
    return a - b
}

func multiply(a int, b int) int {
    return a * b
}

func divide(a int, b int) (int, bool) {
	if b == 0 {
        return 0, false
    }
    return a / b, true
}

func main() {
	fmt.Println("Addition:", add(5, 3))
    fmt.Println("Subtraction:", subtract(5, 3))
    fmt.Println("Multiplication:", multiply(5, 3))
    result, success := divide(10, 0)
    if success {
        fmt.Println("Division:", result)
    } else {
        fmt.Println("Error: Division by zero")
    }
}