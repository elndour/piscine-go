package piscine

func ActiveBits(n int) int {
	v := 0

	for n != 0 {
		if n%2 == 1 {
			v++
		}
		n = n / 2
	}

	return v
}
