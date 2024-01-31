package main

import (
	"io"
	"os"
)

func main() {

	pReader, pWriter := io.Pipe()

	done := make(chan struct{})

	go io.Copy(pWriter, os.Stdin)
	go io.Copy(os.Stdout, pReader)

	<-done

}
