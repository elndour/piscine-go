package piscine

func ToLower(s string) string {
	tab := []rune(s)

	for index, letr := range tab {
		if letr >= 'A' && letr <= 'Z' {
			tab[index] = tab[index] + 32
		}
	}
	return string(tab)
}
