package piscine

func Rot14(s string) string {
	result := []rune{}

	for _, r := range s {
		// Lowercase letters
		if r >= 'a' && r <= 'z' {
			r = ((r-'a'+14)%26 + 'a')
		}

		// Uppercase letters
		if r >= 'A' && r <= 'Z' {
			r = ((r-'A'+14)%26 + 'A')
		}

		result = append(result, r)
	}

	return string(result)
}
