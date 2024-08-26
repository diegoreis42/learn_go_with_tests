package str_met_int

import "math"

type Shape interface {
  Area() float64
}

type Rectangle struct {
  Width float64
  Height float64
}

func (rec Rectangle) Area() float64  {
  return rec.Width * rec.Height
}

func (rec Rectangle) Perimeter() float64 {
  return 2 * (rec.Width + rec.Height)
}

type Circle struct {
  Radius float64
}

func (circ Circle) Area() float64  {
  return math.Pi * circ.Radius * circ.Radius
}

type Triangle struct {
  Base float64
  Height float64
}

func (t Triangle) Area() float64  {
  return (t.Base * t.Height) / 2
}



