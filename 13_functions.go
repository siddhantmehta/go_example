package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

//with the help of three dots we can pass any number of arguments to a function
func addAllValues(values ...int) int {
	total := 0
	for _, value := range values {
		total = total + value
	}
	return total
}

func main() {
	sum := add(1, 2)
	fmt.Println(sum)
	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)
	multiSum := addAllValues(1, 2, 3)
	fmt.Println(multiSum)
	multiSum = addAllValues(1, 2, 3, 4)
	fmt.Println(multiSum)
}
