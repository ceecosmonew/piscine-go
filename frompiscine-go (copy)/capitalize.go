package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, r := range runes {
		// Check if alphanumeric
		isLetter := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
		isDigit := (r >= '0' && r <= '9')

		if isLetter || isDigit {
			if newWord {
				// First character of a word → uppercase letters only
				if r >= 'a' && r <= 'z' {
					runes[i] = r - 32
				}
				newWord = false
			} else {
				// Other characters → lowercase letters only
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + 32
				}
			}
		} else {
			// Non-alphanumeric → next char starts a new word
			newWord = true
		}
	}

	return string(runes)
}
