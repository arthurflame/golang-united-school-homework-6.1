package main

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}


func (t Triangle) CalcArea() float64 {
	// Formula:
	// (1.732 / 4) * side ** 2
	return 0
}

func (t Triangle) CalcPerimeter() float64 {
	// Formula:
	// side * 3

	return 0
}