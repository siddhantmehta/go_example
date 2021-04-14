package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}
	fmt.Println("------------")
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}
	// Loop for goto statement
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum : ", sum)
		if sum > 200 {
			goto theEnd
		}
	}
theEnd:
	fmt.Println("End of a program")
}
