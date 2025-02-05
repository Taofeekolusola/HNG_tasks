package main

import "fmt"

func sum(number ...int) int {
	total := 0;
	for _, num := range number {
        total += num;
    }
	return total;
}

func main() {
    fmt.Println(sum(1, 2, 3, 4, 5)) // Output: 15
}