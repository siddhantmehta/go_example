package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	p := &Point{1, 2} // p is a pointer to Point struct
	p.Move(2, 3)      //since receiver type is pointer hence we are also passing a pointer, otherwise
	//if we pass reference to object and have a pointer as receiver type then go will dynamically type cast it
	fmt.Printf("%+v\n", p)
}
