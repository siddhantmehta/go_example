package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//<-ch //without passing into the channel we are trying to access from it, so it will stuck in deadlock
	go func() {
		ch <- 120 //sending into the channel
	}()
	value := <-ch //receiving from the channel
	fmt.Println(value)
	//Sending data into the channel
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("Sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()
	for i := range ch {
		value := i
		fmt.Printf("Receving %d\n", value)
	}
}
