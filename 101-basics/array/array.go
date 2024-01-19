package main

import "fmt"

var arr_3 [3]int
var arr_5 = [5]int{1, 2, 3, 4, 5}

func main() {

	arr_6 := make([]int, 6)

	arr_6[1] = 10

	fmt.Println(arr_3, arr_5, arr_6) // [0 0 0] [1 2 3 4 5] [0 10 0 0 0 0]

}
