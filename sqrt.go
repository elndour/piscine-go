package piscine

func Sqrt(nb int) int {
	sqrt := 0
	if nb == 0 {
		return 0
	}
	for i := 1; i*i <= nb; i++ {
		if i*i == nb {
			sqrt = i
			return sqrt
		}
	}
	return 0
}
