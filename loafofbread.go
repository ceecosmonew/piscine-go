package piscine

func LoafOfBread(str string) string {
	// Count non-space characters
	count := 0
	for _, r := range str {
		if r != ' ' {
			count++
		}
	}

	// If string has ZERO characters → return "\n"
	if count == 0 {
		return "\n"
	}

	// If string has less than 5 non-space chars → Invalid Output
	if count < 5 {
		return "Invalid Output\n"
	}

	result := ""
	word := ""
	real := 0
	skip := false

	for _, r := range str {
		if skip {
			skip = false
			continue
		}

		if r != ' ' {
			word += string(r)
			real++
		}

		if real == 5 {
			result += word + " "
			word = ""
			real = 0
			skip = true
		}
	}

	if word != "" {
		result += word
	}

	// Remove trailing space
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
