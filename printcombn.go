package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	comb := make([]int, n)
	for i := 0; i < n; i++ {
		comb[i] = i
	}

	print := func() {
		for _, v := range comb {
			z01.PrintRune(rune(v + '0'))
		}
	}

	for {
		print()

		// Check if this is the last combination
		done := true
		for i := 0; i < n; i++ {
			if comb[i] != 10-n+i {
				done = false
				break
			}
		}
		if done {
			break
		}

		z01.PrintRune(',')
		z01.PrintRune(' ')

		// Generate next combination
		i := n - 1
		for i >= 0 && comb[i] == 9-(n-1-i) {
			i--
		}

		comb[i]++

		for j := i + 1; j < n; j++ {
			comb[j] = comb[j-1] + 1
		}
	}

	z01.PrintRune('\n')
}