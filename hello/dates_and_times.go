package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	fmt.Println(n)

	t := time.Date(2021, time.April, 14, 15, 59, 0, 0, time.UTC)
	fmt.Println(t)
	fmt.Println(t.Format(time.ANSIC))
	parsedTime, _ := time.Parse(time.ANSIC, "Wed Apr 14 15:59:00 2021")
	fmt.Printf("The type of parseTime is %+v and type %T\n", parsedTime, parsedTime)
}
