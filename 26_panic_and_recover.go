package main

import (
	"fmt"
)

func safeValue(vals []int, index int) int {
	defer func() { //anonymous function used to handle panic with recover functions help
		if err := recover(); err != nil {
			fmt.Printf("Error %s\n", err)
		}
	}()
	return vals[index]
}
func main() {
	// val := []int{1, 2, 3}
	// v := val[10]
	// fmt.Println(v)
	// file, err := os.Open("abcd.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// fmt.Println("file opened")
	v := safeValue([]int{1, 2, 3}, 10)
	fmt.Println(v)
}
