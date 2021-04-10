package main

import (
	"fmt"
)

//trade is a trade in stocks
type Trade struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func main() {
	t1 := Trade{"MSFT", 10, 99.98, true}
	fmt.Println(t1)
	fmt.Printf("%+v\n", t1)
	fmt.Println(t1.Symbol)

	t2 := Trade{
		Symbol: "MSFT",
		Buy:    true,
		Volume: 20,
		Price:  99.98,
	}
	fmt.Printf("%+v\n", t2)
	t3 := Trade{}
	fmt.Printf("%+v\n", t3)
}
