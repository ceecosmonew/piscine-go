package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := len(base)
	if nbr < 0 {
		z01.PrintRune('-')
		nbr = -nbr
	}

	if nbr < baseLen {
		z01.PrintRune(rune(base[nbr]))
	} else {
		PrintNbrBase(nbr/baseLen, base)
		z01.PrintRune(rune(base[nbr%baseLen]))
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	seen := make(map[rune]bool)
	for _, c := range base {
		if c == '+' || c == '-' || seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}
