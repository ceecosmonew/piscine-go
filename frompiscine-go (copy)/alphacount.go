package piscine

func AlphaCount(s string) int {
	count := 0

	for _, r := range s {
		// Check if r is in the Latin alphabet (A–Z or a–z)
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			count++
		}
	}

	return count
}
