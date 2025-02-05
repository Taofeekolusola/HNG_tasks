package main

import "fmt"

type Rectangle struct {
    Width  float64
    Height float64
}

// Method to calculate the area of a Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{Width: 10, Height: 5}
    fmt.Println("Area:", rect.Area()) // Output: Area: 50
}