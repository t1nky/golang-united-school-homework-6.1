package golang_united_school_homework

// Rectangle must satisfy to Shape interface
type Rectangle struct {
	Height, Weight float64
}

// CalcPerimeter returns calculation result of perimeter
func (c Rectangle) CalcPerimeter() float64 {
	return 2 * (c.Height + c.Weight)
}

// CalcArea returns calculation result of area
func (c Rectangle) CalcArea() float64 {
	return c.Weight * c.Height
}
