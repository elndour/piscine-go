package piscine

func Atoi(s string) int {
	s2 := 0
	if len(s) > 0 {
		f := 1
		for i := len(s) - 1; i >= 0; i-- {

			if (int(s[i]-'0') < 0 || int(s[i]-'0') > 9) && (i != 0 || ((s[0]) != '-' && s[0] != '+')) {
				return 0
			}
			if s[i] != '-' && s[i] != '+' {

				s2 += (int(s[i]) - '0') * f
				f = f * 10
			}
		}
		if s[0] == '-' {
			return -s2
		}

	}
	return s2
}
