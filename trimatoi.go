package piscine

func TrimAtoi(s string) int {
	sign := 1
	result := 0
	foundDigit := false

	for _, r := range s {
		if r == '-' && !foundDigit {
			sign = -1
		} else if r >= '0' && r <= '9' {
			foundDigit = true
			result = result*10 + int(r-'0')
		}
	}

	return sign * result
}
