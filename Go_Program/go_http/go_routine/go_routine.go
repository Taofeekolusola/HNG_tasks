package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
    // Start a Goroutine
    go printNumbers()

    // Main Goroutine continues execution
    for i := 6; i <= 10; i++ {
        fmt.Println(i)
        time.Sleep(500 * time.Millisecond)
    }

    // Sleep to allow the Goroutine to finish
    time.Sleep(3 * time.Second)
}