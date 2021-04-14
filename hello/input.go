package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	input, _ := reader.ReadString('\n') //'\n' is telling ReadString to read until u find \n
	fmt.Println("You entered:", input)
}
