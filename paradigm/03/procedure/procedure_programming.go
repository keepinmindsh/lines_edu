package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Shape interface {
	getArea() int
}

type Square struct {
	TopLeft Point
	Side    int
}

func (s *Square) getArea() int {
	return s.Side * s.TopLeft.X * s.TopLeft.Y
}

func NewSquare() Shape {
	return &Square{}
}

type Rectangle struct {
	TopLeft Point
	Height  int
	Width   int
}

func (r *Rectangle) getArea() int {
	return r.TopLeft.X * r.TopLeft.Y * r.Height * r.Width
}

func NewRectangle() Shape {
	return &Rectangle{}
}

type Circle struct {
	center Point
	radius int
}

func (c *Circle) getArea() int {
	return c.radius * c.center.X * c.center.Y
}

func NewCircle() Shape {
	return &Circle{}
}

func main() {
	area := FactoryShape("CIRCLE")

	fmt.Println(area.getArea())
}

func FactoryShape(shape string) Shape {
	switch shape {
	case "CIRCLE":
		return NewCircle()
	case "SQUARE":
		return NewSquare()
	case "RECTANGLE":
		return NewRectangle()
	default:
		panic("error")
	}
}
