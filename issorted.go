package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	asc := true
	des := true

	for i := 0; i < len(a)-1; i++ {
		result := f(a[i], a[i+1])

		if result > 0 {
			asc = false
		}
		if result < 0 {
			des = false
		}
	}

	return asc || des
}
