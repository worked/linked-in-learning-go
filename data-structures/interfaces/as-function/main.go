package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width, Height float64
}

func main() {
	fmt.Println("interfaces as functions")

	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}

	PrintShapeInfo(circle)
	PrintShapeInfo(rectangle)

	// empty interface
	Describe("pass anything")
	Describe(42)
	Describe(Rectangle{Width: 10, Height: 5})
	Describe(Circle{Radius: 3})
	Describe(true)
	Describe(1.23)
	Describe([]string{"apple", "banana", "cherry"})
	Describe(map[string]int{"apple": 1, "banana": 2, "cherry": 3})
	Describe(nil)
	Describe(Shape(circle))
	Describe(Shape(rectangle))
}

// by satifying the shape methods, rectangle is now a Shape
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// by satifying the shape methods, circle is now a Shape
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// interface as a function parameter
// any type that satisfies the Shape interface can be passed to this function
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

// empty interface - accepts any type
func Describe(i interface{}) {
	fmt.Printf("Type = %T, Value = %v\n", i, i)
}
