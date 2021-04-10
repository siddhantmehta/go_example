package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}
type Square struct {
	Center Point
	Length int
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length < 1 {
		return nil, fmt.Errorf("Length cant be <1 (was %d)", length)
	}
	point_object := Point{
		X: x,
		Y: y,
	}
	s := &Square{
		Center: point_object,
		Length: length,
	}
	return s, nil
}
func (s *Square) Move(dx int, dy int) {
	s.Center.X += dx
	s.Center.Y += dy
}
func (s *Square) Area() int {
	return s.Length * s.Length
}
func main() {
	s, err := NewSquare(1, 1, 5)
	if err != nil {
		log.Fatalf("Error: can't create square %s", err)
	}
	fmt.Printf("%+v\n", s)
	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
