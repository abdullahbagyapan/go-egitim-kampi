package main

import (
	"fmt"
)

func main() {

	str_1 := "hello world"

	str_2 := str_1[6:]

	fmt.Printf("%c - %s", str_1[0], str_2) // h - world
}
