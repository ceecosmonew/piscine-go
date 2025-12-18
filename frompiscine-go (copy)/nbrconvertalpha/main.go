package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	upper := false
	if args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, a := range args {
		num, ok := toInt(a)

		if !ok || num < 1 || num > 26 {
			z01.PrintRune(' ')
			continue
		}

		if upper {
			z01.PrintRune(rune('A' + num - 1))
		} else {
			z01.PrintRune(rune('a' + num - 1))
		}
	}

	// print trailing newline as required by the tests
	z01.PrintRune('\n')
}

// Manual atoi (no strconv) — returns (value, valid)
func toInt(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := 1
	i := 0
	n := 0

	// Optional + or -
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	// Convert each digit
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		n = n*10 + int(c-'0')
	}

	return n * sign, true
}
