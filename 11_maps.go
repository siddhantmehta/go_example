package main

import (
	"fmt"
)

func main() {
	a := map[string]float64{
		"one": 1.0,
		"two": 2.0,
	}
	fmt.Println(len(a))
	fmt.Println(a["three"]) //get 0 value if not found
	fmt.Println(a["one"])
	value, ok := a["three"]
	if !ok {
		fmt.Println("three not found")
	} else {
		fmt.Println(value)
	}
	a["three"] = 3.0 //add
	fmt.Println(a)
	delete(a, "three") //delete
	fmt.Println(a)

	//single value for
	for key := range a {
		fmt.Println(key)
	}

	//double value for
	for key, value := range a {
		fmt.Printf("%s -> %f\n", key, value)
	}

}
