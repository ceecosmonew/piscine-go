package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	fileName := args[1]
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	_, _ = io.Copy(os.Stdout, file)
}
