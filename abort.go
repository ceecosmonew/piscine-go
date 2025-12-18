package piscine

func Abort(a, b, c, d, e int) int {
	// Simple sorting of 5 numbers manually
	for i := 0; i < 5; i++ {
		if a > b {
			a, b = b, a
		}
		if b > c {
			b, c = c, b
		}
		if c > d {
			c, d = d, c
		}
		if d > e {
			d, e = e, d
		}
	}

	// After sorting: a <= b <= c <= d <= e
	// Median is c (the 3rd number)
	return c
}
