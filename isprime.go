package piscine

func IsPrime(nb int) bool {
	if nb <= 0 {
		return false
	} else {
		for i := 2; i < nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
}
