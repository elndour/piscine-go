package piscine

func AlphaCount(s string) int {
	var compt int = 0

	for _, letr := range s {
		if (letr >= 'a' && letr <= 'z') || (letr >= 'A' && letr <= 'Z') {
			compt++
		}
	}
	return compt
}
