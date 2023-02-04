package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 || nb > 20 {
		return 0
	} else if nb == 0 {
		return result
	} else {
		for i := 1; i <= nb; i++ {
			result = result * i
		}
	}
	return result
}
