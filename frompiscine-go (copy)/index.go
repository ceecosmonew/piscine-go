package piscine

func Index(s string, toFind string) int {
	// If toFind is empty, return 0 (same as Go's strings.Index)
	if toFind == "" {
		return 0
	}

	sLen := len(s)
	fLen := len(toFind)

	// If toFind is longer than s, it can't be found
	if fLen > sLen {
		return -1
	}

	// Loop through each valid start index in s
	for i := 0; i <= sLen-fLen; i++ {
		match := true
		for j := 0; j < fLen; j++ {
			if s[i+j] != toFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}
