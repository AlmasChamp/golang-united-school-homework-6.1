package golang_united_school_homework

import "math"

// Circle must satisfy to Shape interface
type Circle struct {
	Radius float64
}

func (c *Circle) CalcPerimeter() float64 {
	return float64(2 * (c.Radius * 3.14))
}

func (c *Circle) CalcArea() float64 {
	root := c.Radius * c.Radius
	return math.Pi * root
}
