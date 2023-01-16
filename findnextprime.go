package piscine

func FindNextPrime(nb int) int {
	if nb <= 0 {
		return 2
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				for i := 2; i < nb; i++ {
					if (nb+1)%i == 0 {
						return nb + 2
					}
				}
				return nb + 1
			}
		}
		return nb
	}
}
