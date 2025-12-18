package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// join all args into a single slice of runes with spaces
	var runes []rune
	for i, a := range args {
		for _, r := range a {
			runes = append(runes, r)
		}
		if i < len(args)-1 {
			runes = append(runes, ' ')
		}
	}

	// collect vowels from runes
	vowels := []rune{}
	for _, r := range runes {
		if isVowel(r) {
			vowels = append(vowels, r)
		}
	}

	// reverse vowels
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// replace vowels in original order but with reversed ones
	vi := 0
	for i, r := range runes {
		if isVowel(r) {
			runes[i] = vowels[vi]
			vi++
		}
	}

	// print result
	for _, r := range runes {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}
