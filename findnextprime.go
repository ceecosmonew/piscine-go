package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	// If nb is even and not 2, make it odd
	if nb%2 == 0 {
		nb++
	}

	for !IsPrime(nb) {
		nb += 2 // skip even numbers
	}

	return nb
}
