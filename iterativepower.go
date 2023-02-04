package piscine

func IterativePower(nb int, power int) int {
	nbr := 1
	if nb > 20 || power < 0 {
		return 0
	} else if nb > 0 && power == 0 {
		return 1
	} else {
		for i := 1; i <= power; i++ {
			nbr = nbr * nb
		}
	}
	return nbr
}
