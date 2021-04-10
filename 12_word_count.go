package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	words := strings.Fields(text)
	word_count := map[string]int{}
	for _, word := range words {
		word_count[strings.ToLower(word)]++
	}
	fmt.Println(word_count)
}
