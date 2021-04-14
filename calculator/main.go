package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value  2 must be a numebr")
	}
	fmt.Print("Select an operation(+ - * /): ")
	input3, _ := reader.ReadString('\n')
	operation := strings.TrimSpace(input3)
	solution := compute(float1, float2, operation)
	solution = math.Round(solution*100) / 100
	fmt.Printf("The %s of %v and %v is %v\n", operation, float1, float2, solution)
}
func compute(value1 float64, value2 float64, operation string) float64 {
	solution := 0.0
	switch operation {
	case "+":
		solution = value1 + value2
	case "-":
		solution = value1 - value2
	case "*":
		solution = value1 * value2
	case "/":
		if value2 == 0.0 {
			panic("Divide by zero")
		}
		solution = value1 / value2
	default:
		panic("Invalid operation")
	}
	return solution
}
