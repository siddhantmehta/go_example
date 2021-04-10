package main

import (
	"fmt"
	"os"

	"github.com/go-openapi/errors"
)

type Config struct {
	//configuration fields go here
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		//return nil,err
		return nil, errors.NotFound("Error:cant open the file", err)
	}
	defer file.Close()
	cfg := &Config{}
	//Parsing file here
	return cfg, nil
}
func main() {
	cfg, err := readConfig("/path/to/config.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error : %s\n", err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}
