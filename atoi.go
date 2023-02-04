package piscine

func Atoi(s string) int {
	nbr := 0
	ind := 0
	if len(s) == 0 {
		ind = ind + 0
	} else if s[0] == '-' || s[0] == '+' {
		ind = ind + 1
	}
	for ; ind < len(s); ind++ {

		if s[ind] < '0' || s[ind] > '9' {
			return 0
		}
		if s[0] == '-' {
			nbr = nbr*10 - int(s[ind]-'0')
		} else {
			nbr = nbr*10 + int(s[ind]-'0')
		}
	}

	return nbr
}
