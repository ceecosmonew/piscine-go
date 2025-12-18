package piscine

func IsUpper(s string) bool {
	for _, me := range s {
		if me < 'A' || me > 'Z' {
			return false
		}
	}
	return true
}
