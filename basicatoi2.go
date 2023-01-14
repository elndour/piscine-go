package piscine

func BasicAtoi2(s string) int {
	s2 := 0
	for i := 0; i < len(s); i++ {

		if int(s[i]-'0') < 0 || int(s[i]-'0') > 9 {
			return 0
		}
		s2 = s2*10 + int(s[i]-'0')
	}
	return s2
}
