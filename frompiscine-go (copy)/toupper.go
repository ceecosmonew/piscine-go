package piscine

func ToUpper(s string) string {
	result := ""

	for _, r := range s {
		// check if r is a lowercase letter
		if r >= 'a' && r <= 'z' {
			r = r - 32 // convert to uppercase
		}
		result += string(r)
	}

	return result
}
