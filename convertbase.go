package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert nbr in baseFrom to an integer in base 10
	value := toIntBase(nbr, baseFrom)

	// Convert the integer value to baseTo
	return fromIntBase(value, baseTo)
}

// Convert number string in baseFrom → int (base 10)
func toIntBase(nbr, base string) int {
	baseLen := len(base)
	result := 0

	for _, r := range nbr {
		result = result*baseLen + indexInBase(r, base)
	}
	return result
}

// Convert int (base 10) → baseTo string
func fromIntBase(n int, base string) string {
	if n == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	result := ""

	for n > 0 {
		result = string(base[n%baseLen]) + result
		n /= baseLen
	}
	return result
}

// Find the index of a rune in base string
func indexInBase(r rune, base string) int {
	for i, b := range base {
		if r == b {
			return i
		}
	}
	return 0 // base is always valid, so this won't happen
}
