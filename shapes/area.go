package shapes

import "math"

//Circle Structure for circle
type Circle struct {
	Radius float64
}

//Rect Structure for rectangle
type Rect struct {
	Width  float64
	Height float64
}

//Shape : generic interface for shapes
type Shape interface {
	Area() float64
}

//Area for rectangle
func (r *Rect) Area() float64 {
	return r.Width * r.Height
}

//Area for circle
func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
