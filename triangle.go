package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

// CalcPerimeter returns calculation result of perimeter
func (c Triangle) CalcPerimeter() float64 {
	return c.Side * 3
}

// CalcArea returns calculation result of area
func (c Triangle) CalcArea() float64 {
	return math.Sqrt(3) / 4 * math.Pow(c.Side, 2)
}
