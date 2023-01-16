package piscine

func IterativeFactorial(nb int) int {
	s2 := 1
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		s2 = s2 * i
	}
	return s2
}
