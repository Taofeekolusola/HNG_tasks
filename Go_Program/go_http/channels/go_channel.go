package main

import (
    "fmt"
    "time"
)

func printNumbers(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i // Send value to the channel
        time.Sleep(500 * time.Millisecond)
    }
    close(ch) // Close the channel when done
}

func main() {
    // Create a channel
    ch := make(chan int)
    fmt.Println("Creating channel")
    // Start a Goroutine
    go printNumbers(ch)

    // Receive values from the channel
    for num := range ch {
        fmt.Println(num)
    }
    fmt.Println("Finished receiving values")
}