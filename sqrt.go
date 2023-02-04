package piscine

func Sqrt(nb int) int {
	var rac int
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			rac = i
		}
	}

	return rac
}
