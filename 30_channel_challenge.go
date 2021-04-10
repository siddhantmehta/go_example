package main

import (
	"fmt"
	"net/http"
)

func returnTypes(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("error : %s\n", err)
		return
	}

	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s\n", url, ctype)
}
func main() {
	urls := []string{
		"http://google.com",
		"http://facbook.com",
		"http://linkedin.com",
	}
	ch := make(chan string)
	for _, url := range urls {
		go returnTypes(url, ch)
	}
	for range urls {
		ctype := <-ch
		fmt.Printf("%s\n", ctype)
	}
}
