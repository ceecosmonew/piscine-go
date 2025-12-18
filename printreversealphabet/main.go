package main

import "github.com/01-edu/z01"

func main() {
	r := 'z'
	for r >= 'a' {
		z01.PrintRune(r)
		r--
	}
	z01.PrintRune('\n')
}
