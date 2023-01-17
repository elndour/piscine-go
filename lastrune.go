package piscine

func LastRune(s string) rune {
	a := []rune(s)

	return rune(a[len(a)-1])
}
