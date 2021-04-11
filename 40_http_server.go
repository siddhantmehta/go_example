package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

type MathRequest struct {
	Op    string  `json:"op"`
	Left  float64 `json:"left"`
	Right float64 `json:"right"`
}
type MathResponse struct {
	Error  string  `json:"error"`
	Result float64 `json:"result"`
}

func mathHandler(w http.ResponseWriter, r *http.Request) {
	//decoding request
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	req := &MathRequest{}
	if err := dec.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//doing work
	resp := &MathResponse{}
	switch req.Op {
	case "+":
		resp.Result = req.Left + req.Right
	case "-":
		resp.Result = req.Left - req.Right
	case "*":
		resp.Result = req.Left * req.Right
	case "/":
		if req.Right == 0.0 {
			resp.Error = "Divide by 0"
		} else {
			resp.Result = req.Left / req.Right
		}
	default:
		resp.Error = fmt.Sprintf("unkown operation: %s", req.Op)
	}
	//Encoding and sending result
	w.Header().Set("Content-Type", "application/json")
	if resp.Error != "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		log.Printf("Can't encode %v - %s", resp, err)
	}
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/math", mathHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

//Here we will send POST request via postman to the server on localhost:8080/math
