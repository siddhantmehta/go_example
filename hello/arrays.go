package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	fmt.Println(colors)
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors)
	var numbers = [5]int{1, 2, 3, 4}
	fmt.Println(numbers)
	fmt.Println("Number of colors : ", len(colors))
}
