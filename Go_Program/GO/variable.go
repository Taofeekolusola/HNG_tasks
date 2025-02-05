package main

import "fmt"
var name string = "Taofeek";
var age int = 32;

func greet(name string, age int) {
    fmt.Printf("Hello, %s! You are %d years old.\n", name, age)
}

func main() {
    greet(name, age)
}