package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Save original values
	valA := ***a
	valB := *b
	valC := *******c
	valD := ****d

	// Rotate them
	*******c = valA
	****d = valC
	*b = valD
	***a = valB
}
