package piscine

func Maj(letr rune) bool {
	if (letr >= 'a' && letr <= 'z') || (letr >= 'A' && letr <= 'Z') || (letr >= '0' && letr <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	tab := []rune(s)

	T0 := true
	for i := range tab {
		if Maj(tab[i]) && T0 {
			if tab[i] >= 'a' && tab[i] <= 'z' {
				tab[i] = tab[i] - 32
			}
			T0 = false
		} else if tab[i] >= 'A' && tab[i] <= 'Z' {
			tab[i] = tab[i] + 32
		} else if !(Maj(tab[i])) {
			T0 = true
		}
	}

	return string(tab)
}
