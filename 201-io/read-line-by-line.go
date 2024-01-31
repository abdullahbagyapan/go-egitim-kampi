package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("testfile2.txt")

	if err != nil {
		fmt.Printf("error opening file, %v", err)
	}

	//readLineByLineWithScanner(file)
	readLineByLineWithReader(file)

}

func readLineByLineWithScanner(file *os.File) {

	scanner := bufio.NewScanner(file)
	i := 1

	for scanner.Scan() {
		fmt.Println(i)
		fmt.Println(scanner.Text())
		i++
	}

}

func readLineByLineWithReader(file *os.File) {

	reader := bufio.NewReader(file)
	i := 1
	for {
		buf, err := reader.ReadBytes('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("error reading file, %v", err)
			break
		}

		fmt.Println(i)
		fmt.Printf("%s", buf)
		i++
	}

}
