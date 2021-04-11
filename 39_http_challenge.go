package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name        string `json:"name"`
	PublicRepos int    `json:"public_repos"`
}

func userInfo(login string) (*User, error) {
	url := "https://api.github.com/users/" + login
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error: can't call httpbin.org/get")
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	usr := &User{}
	if err := dec.Decode(usr); err != nil {
		return nil, err
	}
	return usr, nil
}
func main() {
	login := "siddhantmehta"
	usr, err := userInfo(login)
	if err != nil {
		log.Fatalf("error: can't find the data for %s - %s", login, err)
	}
	fmt.Printf("%+v\n", usr)
}
