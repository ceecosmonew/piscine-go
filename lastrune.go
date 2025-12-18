package piscine

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}

	last := rune(0)

	for i := 0; i < len(s); {
		r, size := decodeRune(s[i:])
		last = r
		i += size
	}

	return last
}

func decodeRune(s string) (rune, int) {
	b := s[0]

	// 1-byte ASCII
	if b < 0x80 {
		return rune(b), 1
	}

	// 2-byte UTF-8
	if b < 0xE0 {
		return rune(b&0x1F)<<6 |
			rune(s[1]&0x3F), 2
	}

	// 3-byte UTF-8
	if b < 0xF0 {
		return rune(b&0x0F)<<12 |
			rune(s[1]&0x3F)<<6 |
			rune(s[2]&0x3F), 3
	}

	// 4-byte UTF-8
	return rune(b&0x07)<<18 |
		rune(s[1]&0x3F)<<12 |
		rune(s[2]&0x3F)<<6 |
		rune(s[3]&0x3F), 4
}
