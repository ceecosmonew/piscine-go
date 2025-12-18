package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sign := 1
	start := 0

	// Handle optional leading + or -
	if s[0] == '+' {
		start = 1
	} else if s[0] == '-' {
		sign = -1
		start = 1
	}

	// If there is more than 1 sign → invalid
	if len(s) > 1 && (s[1] == '+' || s[1] == '-') {
		return 0
	}

	result := 0

	for _, r := range s[start:] {
		if r < '0' || r > '9' {
			return 0 // invalid character
		}
		result = result*10 + int(r-'0')
	}

	return result * sign
}
