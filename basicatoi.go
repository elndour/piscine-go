package piscine

func BasicAtoi(s string) int {
	nbr := 0

	for i := 0; i <= len(s)-1; i++ {
		nbr = nbr*10 + int((s[i] - '0'))
	}

	return nbr
}
