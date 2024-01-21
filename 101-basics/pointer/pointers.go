package main

import (
	"fmt"
)

type String *string

func main() {

	var str string
	var pstr String

	fmt.Println(str, pstr) // <nil>

	str = "Hello World"
	pstr = &str

	fmt.Println(str, pstr)  // Hello World <memory address>
	fmt.Println(str, *pstr) // Hello World Hello World

	str = "Pointers are easy"

	fmt.Println(str, pstr)  // Pointers are easy <memory addess>
	fmt.Println(str, *pstr) // Pointers are easy Pointers are easy

	*pstr = "Golang is easy"

	fmt.Println(str, *pstr) // Golang is easy Golang is easy

	/*
		printUpperCase(str)
		fmt.Println(str)
	*/
}

/*

func printUpperCase(s string) {
	s = strings.ToUpper(s)
}

func printUpperCase(s *string) {
	*s = strings.ToUpper(*s)
}

*/
