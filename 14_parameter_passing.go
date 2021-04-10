package main

import (
	"fmt"
)

func doubleAt(n []int, i int) {
	n[i] *= 2
}
func double(n int) {
	n *= 2
}
func doublePtr(n *int) {
	*n *= 2
}
func main() {
	n := []int{1, 2, 3, 4}
	doubleAt(n, 2)
	fmt.Println(n)
	val := 10
	double(val)
	fmt.Println(val)
	val = 20
	doublePtr(&val)
	fmt.Println(val)
}
