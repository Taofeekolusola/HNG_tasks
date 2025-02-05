package main

import "fmt"

func sayHello(name string, greet string) string {
	return fmt.Sprintf("%s, %s!", greet, name)
}

func main() {
    fmt.Println(sayHello("John Doe", "Good evening"))
}