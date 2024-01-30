package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//readFileByOsReadFile()
	readFileByBufio()
}

func readFileByOsReadFile() {

	buffer, err := os.ReadFile("testfile2.txt")

	if err != nil {
		log.Printf("error reading file %v", err)
	}

	fmt.Printf("%s", buffer)

}

func readFileByBufio() {

	file, err := os.Open("testfile2.txt")

	if err != nil {
		log.Printf("error opening file %v", err)
	}

	reader := bufio.NewReader(file) // read whole file
	io.Copy(os.Stdout, reader)

}
