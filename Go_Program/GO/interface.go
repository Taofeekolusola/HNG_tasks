package main

import (
	"math"
    "fmt"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
    width  float64
    height float64
}

type Circle struct {
    radius float64
}

func (r Rectangle) Area() float64 {
    return r.width * r.height
}

func (c Circle) Area() float64 {
    return math.Pi * c.radius * c.radius
}

func calculateTotalArea(shapes []Shape) float64 {
	var totalArea float64 = 0.0
    for _, shape := range shapes {
        totalArea += shape.Area()
    }
    return totalArea
}

func main() {
	rectangles := []Shape{
        Rectangle{width: 10, height: 5},
        Rectangle{width: 8, height: 6},
    }

    circles := []Shape{
        Circle{radius: 3},
        Circle{radius: 4},
    }

    totalAreaRectangles := calculateTotalArea(rectangles)
    totalAreaCircles := calculateTotalArea(circles)

    fmt.Printf("Total area of rectangles: %.2f\n", totalAreaRectangles)
    fmt.Printf("Total area of circles: %.2f\n", totalAreaCircles)
}