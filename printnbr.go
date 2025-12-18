package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	if n < 0 {
		z01.PrintRune('-')
		// Avoid overflow on MinInt
		if n == -9223372036854775808 {
			// Print manually: "9223372036854775808"
			s := "9223372036854775808"
			for _, r := range s {
				z01.PrintRune(r)
			}
			return
		}
		n = -n
	}

	// Collect digits
	var digits []int
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	// Print digits in reverse
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i]
		z01.PrintRune(rune('0' + d))
	}
}
