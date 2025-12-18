package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]

	// Remove path and keep only the name after the last '/'
	lastSlash := -1
	for i, r := range name {
		if r == '/' {
			lastSlash = i
		}
	}

	// If '/' exists, take substring after it
	if lastSlash != -1 {
		name = name[lastSlash+1:]
	}

	// Print the program name
	for _, r := range name {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}
