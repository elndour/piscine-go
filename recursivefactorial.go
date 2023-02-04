package piscine

func RecursiveFactorial(nb int) int {
	var nbr int

	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		nbr = nb * RecursiveFactorial(nb-1)
	}

	return nbr
}
