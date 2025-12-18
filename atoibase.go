package piscine

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	result := 0
	baseLen := len(base)

	for i := 0; i < len(s); i++ {
		// find value of s[i] in base
		value := 0
		for base[value] != s[i] {
			value++
		}
		result = result*baseLen + value
	}

	return result
}
