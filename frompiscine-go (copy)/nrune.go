package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	index := 0
	count := 0

	for index < len(s) {
		r, size := decodeRune(s[index:])
		count++

		if count == n {
			return r
		}

		index += size
	}

	return 0
}

// Manual UTF-8 decoder (same logic used for FirstRune / LastRune)
func decodeRune(s string) (rune, int) {
	b := s[0]

	if b < 0x80 {
		return rune(b), 1
	}
	if b < 0xE0 {
		return rune(b&0x1F)<<6 |
			rune(s[1]&0x3F), 2
	}
	if b < 0xF0 {
		return rune(b&0x0F)<<12 |
			rune(s[1]&0x3F)<<6 |
			rune(s[2]&0x3F), 3
	}
	return rune(b&0x07)<<18 |
		rune(s[1]&0x3F)<<12 |
		rune(s[2]&0x3F)<<6 |
		rune(s[3]&0x3F), 4
}
