package piscine

func StringToIntSlice(str string) []int {
	// If empty string → return nil
	if str == "" {
		return nil
	}

	// Count runes first
	length := 0
	for range str {
		length++
	}

	result := make([]int, length)

	i := 0
	for _, r := range str {
		result[i] = int(r)
		i++
	}

	return result
}
