package piscine

func AppendRange(min, max int) []int {
	var tab []int = nil

	if min >= max {
		return tab
	}
	if min == 0 && max == 0 {
		return tab
	}

	if min < max {
		for i := min; i < max; i++ {
			tab = append(tab, min)
			min = min + 1
		}
	}
	return tab
}
