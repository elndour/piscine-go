package piscine

func StrRev(s string) string {
	s2 := []rune(s)
	for i, j := 0, len(s2)-1; i < j; i, j = i+1, j-1 {
		s2[i], s2[j] = s2[j], s2[i]
	}
	return string(s2)
}
