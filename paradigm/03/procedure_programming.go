package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Square struct {
	TopLeft Point
	Side    int
}

type Rectangle struct {
	TopLeft Point
	Height  int
	Width   int
}

type Circle struct {
	center Point
	radius int
}

func main() {
	area := getArea("CIRCLE")

	fmt.Println(area)
}

func getArea(shape string) int {
	if shape == "CIRCLE" {
		circle := &Circle{}
		return circle.radius * circle.center.Y * circle.center.X
	} else if shape == "SQUARE" {
		square := &Square{}
		return square.TopLeft.X * square.TopLeft.Y * square.Side
	} else {
		rectangle := &Rectangle{}
		return rectangle.Width * rectangle.Height * rectangle.TopLeft.X * rectangle.TopLeft.Y
	}
}
