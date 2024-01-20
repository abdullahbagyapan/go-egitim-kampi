package main

import "fmt"

func main() {

	map_1 := make(map[int]string, 3)

	map_1[1] = "1"
	map_1[2] = "2"

	fmt.Println(map_1) // map[1:1 2:2]
}
