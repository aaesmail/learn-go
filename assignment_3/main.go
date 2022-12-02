package main

import (
	"fmt"
	"io"
	"os"
)

type consoleWriter struct {}

func main() {
	fileName := os.Args[1]
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	cw := consoleWriter{}

	io.Copy(cw, file)
}

func (consoleWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))

	return len(p), nil
}
