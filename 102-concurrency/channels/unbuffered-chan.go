package main

import (
	"fmt"
	"time"
)

func printValue(ch <-chan string) {

	// blocking until read data
	value := <-ch

	fmt.Println(value)
}

func main() {

	ch := make(chan string)

	go printValue(ch)

	ch <- "hello world"

	time.Sleep(time.Second) // for schedule goroutine
}
