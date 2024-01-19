package main

import "fmt"

func main() {

	slice := make([]int, 0, 10) // 0: length,10: capasity

	slice = append(slice, 20)

	fmt.Println(slice, len(slice), cap(slice)) // [20] 1 10

}
