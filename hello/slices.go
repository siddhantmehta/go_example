package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"red", "green", "blue"}
	fmt.Println(colors)
	colors = append(colors, "purple")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)]) //deleting the first item of slice
	fmt.Println(colors)
	numbers := make([]int, 5) //first argument is type and second is the initial size
	numbers[0] = 1
	numbers[1] = 10
	numbers[2] = 30
	numbers[3] = 54
	numbers[4] = 29
	fmt.Println(numbers)
	numbers = append(numbers, 3)
	fmt.Println(numbers)
	sort.Ints(numbers)
	fmt.Println(numbers)
}
