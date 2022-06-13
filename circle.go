package main

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcArea() float64 {
	// Formula:
	// Pi * Radius ** 2;
	return math.Pi * (c.Radius * 2)

}

func (c Circle) CalcPerimeter() float64 {
	// Formula:
	// 2 * Pi * Radius

	return 2 * math.Pi * c.Radius
}
