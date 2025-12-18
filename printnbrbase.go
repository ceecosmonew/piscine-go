package piscine

import "github.com/01-edu/z01"

// Check if base is valid
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

// Recursive printer (expects positive number)
func printNbr(n int, base string) {
	baseLen := len(base)

	if n >= baseLen {
		printNbr(n/baseLen, base)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	// Handle zero
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	baseLen := len(base)

	// Handle negative numbers
	if nbr < 0 {
		z01.PrintRune('-')

		// Handle MinInt safely
		if nbr == -nbr {
			printNbr(-(nbr / baseLen), base)
			z01.PrintRune(rune(base[-(nbr % baseLen)]))
			return
		}

		nbr = -nbr
	}

	printNbr(nbr, base)
}
