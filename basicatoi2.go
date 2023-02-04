package piscine

func BasicAtoi2(s string) int {
	nbr := 0
	for i := 0; i < len(s); i++ {

		nbr = nbr*10 + int(s[i]-'0')

		if int(s[i]-'0') < 0 || int(s[i]-'0') > 9 {
			return 0
		}
	}

	return nbr
}
