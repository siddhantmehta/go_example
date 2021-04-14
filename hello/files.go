package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello from Go!"
	file, err := os.Create("something.txt")
	checkError(err)
	defer file.Close()
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("wrote a file with %v characters\n", length)
	defer readFile("something.txt")
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	fmt.Println("Text read from file:", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
