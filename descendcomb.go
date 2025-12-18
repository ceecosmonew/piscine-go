package piscine

import "github.com/01-edu/z01"

func printNumber(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}

func DescendComb() {
	for a := 99; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {

			// Print a
			printNumber(a)
			z01.PrintRune(' ')
			printNumber(b)

			// Do NOT print comma after the last combination (01 00)
			if !(a == 1 && b == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
