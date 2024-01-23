package main

import (
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		time.Sleep(time.Second)
		wg.Done()
		println("Process done")
	}()

	wg.Wait() //block until goroutine finish
	println("Exitting...")

}
