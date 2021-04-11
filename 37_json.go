package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
{
	"user": "Siddhant Mehta",
	"type": "deposit",
	"amount": 10500.45
}
`

type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	rdr := bytes.NewBufferString(data)
	dec := json.NewDecoder(rdr) //decoding request
	req := &Request{}
	if err := dec.Decode(req); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	fmt.Printf("got : %+v\n", req)

	//creating response
	prevBalance := 5000.0
	response := map[string]interface{}{ //interface{} is known as any type it allow us to put any value in map dynamically
		"ok":      true,
		"balance": prevBalance + req.Amount,
	}
	//encoding response
	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(response); err != nil {
		log.Fatalf("error: can't encode - %s", err)
	}
}
