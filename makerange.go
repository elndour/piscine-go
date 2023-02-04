package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	tab := make([]int, max-min)

	if min < max {
		for i := 0; i < max-min; i++ {
			tab[i] = min + i
		}
	}

	return tab
}
