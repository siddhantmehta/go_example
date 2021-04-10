package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

func (s *Square) Area() float64 {
	return s.Length * s.Length
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
func sumAreas(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

//Shape is an interface
type Shape interface {
	Area() float64
}

func main() {
	s := &Square{10}
	fmt.Println(s.Area())
	c := &Circle{5}
	fmt.Println(c.Area())

	shapes := []Shape{s, c}
	sum_of_areas := sumAreas(shapes)
	fmt.Println(sum_of_areas)
}
