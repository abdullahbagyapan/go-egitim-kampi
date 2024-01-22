package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		j := i
		go func() {
			//fmt.Println(i)
			fmt.Println(j)
		}()
	}

	time.Sleep(time.Second * 2)
}
