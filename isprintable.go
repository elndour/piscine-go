package piscine

func IsPrintable(s string) bool {
	for _, letter := range s {
		if (letter) >= 7 && (letter) <= 13 {
			return false
		}
	}
	return true
}
