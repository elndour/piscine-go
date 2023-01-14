package piscine

func BasicAtoi(s string) int {
	s2 := 0
	for i := 0; i < len(s); i++ {
		s2 = s2*10 + int(s[i]-'0')
	}
	return s2
}
