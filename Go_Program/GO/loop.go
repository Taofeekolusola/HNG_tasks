package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	j := 0
	for j < 10 {
        fmt.Println("j:", j)
        j++
    }

	k := 0;
	for{
		fmt.Println("k:", k)
        k++
        if k >= 5 {
            break
        }
	}
}