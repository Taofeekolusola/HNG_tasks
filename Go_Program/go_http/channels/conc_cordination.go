package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, job)
        time.Sleep(time.Second) // Simulate work
        fmt.Printf("Worker %d finished job %d\n", id, job)
        results <- job * 2
    }
}

func main() {
    // Create channels
    jobs := make(chan int, 10)
    results := make(chan int, 10)

    // Start 3 worker Goroutines
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send 5 jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect results
    for a := 1; a <= 5; a++ {
        fmt.Println("Result:", <-results)
    }
}