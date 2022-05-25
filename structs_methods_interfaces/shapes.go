package shapes

import "math"

// ----------------------------------------------------------
// Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// ----------------------------------------------------------
// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}


// ----------------------------------------------------------
// Triangle
type Triangle struct {
    Base float64
    Height float64
}

func (t Triangle) Area() float64 {
    return t.Base * t.Height / 2
}


// ----------------------------------------------------------
// Interface
type Shape interface {
    Area() float64
}



