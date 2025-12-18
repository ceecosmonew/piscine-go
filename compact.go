package piscine

func Compact(ptr *[]string) int {
	original := *ptr
	newSlice := make([]string, 0)

	for _, v := range original {
		if v != "" {
			newSlice = append(newSlice, v)
		}
	}

	*ptr = newSlice

	return len(newSlice)
}
