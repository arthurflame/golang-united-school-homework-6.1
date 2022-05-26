package main

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c Circle) CalcArea() float64 {
	// TODO: Area formula
	// Pi * Radius ** 2;
	return 0
}

func (c Circle) CalcPerimeter() float64 {
	// TODO: Perimeter formula
	// 2 * Pi * Radius

	return 0
}