package piscine

func NRune(s string, n int) rune {
	a := []rune(s)

	if n <= 0 || n > len(a) {
		return 0
	}

	return rune(a[n-1])
}
