package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// Sort arguments using ASCII order
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-1-i; j++ {
			if compare(args[j], args[j+1]) > 0 {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Print sorted arguments
	for _, word := range args {
		for _, r := range word {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}

// compare returns:
// < 0 if a < b
// > 0 if a > b
// 0 if equal
func compare(a, b string) int {
	ra := []rune(a)
	rb := []rune(b)

	i := 0
	for i < len(ra) && i < len(rb) {
		if ra[i] != rb[i] {
			return int(ra[i]) - int(rb[i])
		}
		i++
	}
	return len(ra) - len(rb)
}
