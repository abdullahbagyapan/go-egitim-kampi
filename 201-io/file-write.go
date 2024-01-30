package main

import (
	"io/fs"
	"log"
	"os"
)

func main() {
	//writeFileByOsWrite()
	writeFileByFWrite()
}

func writeFileByOsWrite() {

	err := os.WriteFile("testfile1.txt", []byte("hello world"), fs.ModePerm)

	if err != nil {
		log.Printf("failed writing data to file %s", err)
	}
}

func writeFileByFWrite() {

	f, err := os.Create("testfile2.txt")

	if err != nil {
		log.Printf("failed creating file %s", err)
	}

	f.Write([]byte("1-hello world\n"))
	f.Write([]byte("2-hello world\n"))
}
