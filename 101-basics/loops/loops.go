package main

import "fmt"

func main() {
	slc_1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	slc_2 := []int{}

	for i, value := range slc_1 {
		fmt.Printf("index: %d, value: %d \n", i, value)
	}

	size := len(slc_1)
	for i := 0; i < size; i++ {
		slc_2 = append(slc_2, slc_1[i])
	}

	fmt.Println(slc_2) // [10 20 30 40 50 60 70 80 90]

}
