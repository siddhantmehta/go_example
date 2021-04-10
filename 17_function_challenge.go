package main

import (
	"fmt"
	"net/http"
)

func content(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close() // make sure we close the body
	contentType := response.Header.Get("Content-Type")
	if contentType == "" {
		return "", fmt.Errorf("Can't find content-type header")
	}
	return contentType, nil
}
func main() {
	contentType, err := content("https://linkedin.com")
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	} else {
		fmt.Println(contentType)
	}
}
