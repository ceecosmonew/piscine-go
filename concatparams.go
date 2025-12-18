package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}

	// Find total length needed
	totalLen := 0
	for _, s := range args {
		totalLen += len(s) + 1 // +1 for newline
	}

	// Create a byte slice for efficient concatenation
	result := make([]byte, 0, totalLen)

	for i, s := range args {
		for _, ch := range s {
			result = append(result, byte(ch))
		}
		if i < len(args)-1 {
			result = append(result, '\n')
		}
	}

	return string(result)
}
