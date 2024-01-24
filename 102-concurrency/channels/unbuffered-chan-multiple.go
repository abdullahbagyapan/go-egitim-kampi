package main

import (
	"fmt"
	"time"
)

func readValueFromChannel(ch <-chan int) {
	value := <-ch
	fmt.Println(value)
}

func writeValueToChannel(ch chan<- int) {
	ch <- 10
}

func main() {

	ch := make(chan int)

	go readValueFromChannel(ch)
	go writeValueToChannel(ch)

	time.Sleep(time.Second) // give a time sleep to schedule goroutines
}
