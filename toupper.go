package piscine

func ToUpper(s string) string {
	tab := []rune(s)
	for i, letr := range tab {
		if letr >= 'a' && letr <= 'z' {
			tab[i] = tab[i] - 32
		}
	}
	return string(tab)
}
