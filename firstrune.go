package piscine

func FirstRune(s string) rune {
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		return rune(a[0])
	}
	return rune(a[0])
}
