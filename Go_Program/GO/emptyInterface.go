package main

import "fmt"

func printVal(v interface{}) {
	switch v.(type) {
    case int:
        fmt.Printf("Value is an integer: %d\n", v.(int))
    case float64:
        fmt.Printf("Value is a float: %.2f\n", v.(float64))
    case string:
        fmt.Printf("Value is a string: %s\n", v.(string))
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
	printVal(5)
    printVal(3.14)
    printVal("Hello, World!")
    printVal(map[string]int{"key": 1}) // Unknown type, as it's not a primitive type or a struct with a known field
    printVal([]int{1, 2, 3})            // Unknown type, as it's not a primitive type or a struct with a known field
}