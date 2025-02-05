package main

import (
    "fmt"
    "github.com/TaofeekOlusola/go_learn/mathutils"
)

func main() {
    sum := mathutils.Add(10, 5)
    difference := mathutils.Subtract(10, 5)

    fmt.Println("Sum:", sum)           // Output: Sum: 15
    fmt.Println("Difference:", difference) // Output: Difference: 5
}