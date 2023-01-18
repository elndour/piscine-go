package piscine

func TrimAtoi(s string) int {
	s2 := 0
	for _, letter := range s {
		if letter >= '0' && letter <= '9' {
			s2 = (10 * s2) + int(letter-48)
		}
	}
	if Isneg(s) {
		s2 = s2 * -1
	}

	return s2
}

func Isneg(s string) bool {
	j := 0
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' && i > 0 {
			for j > 0 {
				if s[i-j] == '-' {
					return true
				}
				j--
			}
			return false
		}
		j++
	}
	return false
}
