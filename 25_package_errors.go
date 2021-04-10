package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/errors"
)

type Config struct {
	//configuration fields go here
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		// return nil, err
		return nil, errors.NotFound("Error:cant open the file", err)
	}
	defer file.Close()
	cfg := &Config{}
	//Parsing file here
	return cfg, nil
}
func setupLogging() {
	//out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	out, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.SetOutput(out)
}
func main() {
	setupLogging()
	cfg, err := readConfig("/path/to/config.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %s\n", err)
		log.Printf("Error : %+v", err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}
