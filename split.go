package piscine

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}

	var result []string
	sepLen := len(sep)
	start := 0

	for i := 0; i+sepLen <= len(s); {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			i += sepLen
			start = i
		} else {
			i++
		}
	}

	// Add the last piece after final separator
	result = append(result, s[start:])

	return result
}
