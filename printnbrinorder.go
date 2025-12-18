package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// handle negative numbers and zero
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		return
	}

	// array to count digits 0–9
	var digits [10]int

	// extract digits
	for n > 0 {
		d := n % 10
		digits[d]++
		n = n / 10
	}

	// print digits in ascending order
	for i := 0; i <= 9; i++ {
		for digits[i] > 0 {
			z01.PrintRune(rune(i + '0'))
			digits[i]--
		}
	}
}
