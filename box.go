package golang_united_school_homework

import "errors"

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
	if len(b.shapes) == b.shapesCapacity {
		return errors.New("error while addint the shape, the box is full")
	}
	b.shapes = append(b.shapes, shape)
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("error while getting the shape by index, index went out of range")
	}
	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("error while extracting the shape by index, index went out of range")
	}
	shape := b.shapes[i]
	tmpShapes := []Shape{}
	tmpShapes = append(tmpShapes, b.shapes[0:i]...)
	if i != len(b.shapes)-1 {
		tmpShapes = append(tmpShapes, b.shapes[i+1:]...)
	}
	b.shapes = tmpShapes
	return shape, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errors.New("error while replacing the shape by index, index went out of range")
	}
	rpShape := b.shapes[i]
	b.shapes[i] = shape
	return rpShape, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	var sum float64
	for _, shape := range b.shapes {
		sum += shape.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	cnt := 0
	for i, shape := range b.shapes {
		_, ok := shape.(Circle)
		if ok {
			cnt++
			tmpShapes := []Shape{}
			tmpShapes = append(tmpShapes, b.shapes[0:i]...)
			if i != len(b.shapes)-1 {
				tmpShapes = append(tmpShapes, b.shapes[i+1:]...)
			}
			b.shapes = tmpShapes
		}
	}
	if cnt == 0 {
		return errors.New("error while removing all circles, there is no circle in the box")
	}
	return nil
}
