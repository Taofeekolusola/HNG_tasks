package main

import "fmt"

var age int = 18;

func main(){
	fmt.Println("Your age is %d.", age)
    if age >= 18 {
        fmt.Printf("You are an adult.\n")
    } else {
        fmt.Printf("You are a child.\n")
    }
}