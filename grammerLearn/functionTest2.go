package main

import "fmt"

type Circle struct {
	radius float64
}

func getArea(circle Circle) float64 {
	return 3.14 * circle.radius * circle.radius
}

func main() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle(c1) = ", getArea(c1))
}