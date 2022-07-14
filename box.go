package golang_united_school_homework

import "errors"

func cut(i int, xs []Shape) (Shape, []Shape) {
	y := xs[i]
	ys := append(xs[:i], xs[i+1:]...)
	return y, ys
}

var errorNotFound = errors.New("not found")

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if b.shapesCapacity <= len(b.shapes) {
		return errors.New("capacity limit reached")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if b.shapesCapacity <= i {
		return nil, errorNotFound
	}
	if shape := b.shapes[i]; shape != nil {
		return shape, nil
	}

	return nil, errorNotFound
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if b.shapesCapacity <= i {
		return nil, errorNotFound
	}
	if shape, newShapes := cut(i, b.shapes); shape != nil {
		b.shapes = newShapes
		return shape, nil
	}
	return nil, errorNotFound
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if b.shapesCapacity <= i {
		return nil, errorNotFound
	}
	if oldShape := b.shapes[i]; oldShape != nil {
		b.shapes[i] = shape
		return oldShape, nil
	}
	return nil, errorNotFound

}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	res := 0.0
	for _, shape := range b.shapes {
		res += shape.CalcPerimeter()
	}
	return res
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	res := 0.0
	for _, shape := range b.shapes {
		res += shape.CalcArea()
	}
	return res
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	newShapes := []Shape{}
	for _, shape := range b.shapes {
		if _, ok := shape.(Circle); !ok {
			newShapes = append(newShapes, shape)
		}
	}

	if len(newShapes) == len(b.shapes) {
		return errorNotFound
	}
	b.shapes = newShapes
	return nil
}
