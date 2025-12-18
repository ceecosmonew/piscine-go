package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for i := 0; i < len(str); i++ {
		c := str[i]

		if c == ' ' {
			// treat empty or finished words as valid entries
			result[word]++
			word = ""
		} else {
			word += string(c)
		}
	}

	// last word after loop ends
	result[word]++

	return result
}
