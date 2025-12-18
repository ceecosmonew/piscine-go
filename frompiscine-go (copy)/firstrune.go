package piscine

func FirstRune(s string) rune {
	r, _ := decodeRune(s)
	return r
}

func decodeRune(s string) (rune, int) {
	if len(s) == 0 {
		return 0, 0
	}

	b := s[0]

	// 1-byte rune (ASCII)
	if b < 0x80 {
		return rune(b), 1
	}

	// 2-byte rune
	if b < 0xE0 {
		return rune(b&0x1F)<<6 |
			rune(s[1]&0x3F), 2
	}

	// 3-byte rune
	if b < 0xF0 {
		return rune(b&0x0F)<<12 |
			rune(s[1]&0x3F)<<6 |
			rune(s[2]&0x3F), 3
	}

	// 4-byte rune
	return rune(b&0x07)<<18 |
		rune(s[1]&0x3F)<<12 |
		rune(s[2]&0x3F)<<6 |
		rune(s[3]&0x3F), 4
}
