package piscine

func IsLower(s string) bool {
	for _, me := range s {
		if me < 'a' || me > 'z' {
			return false
		}
	}
	return true
}
