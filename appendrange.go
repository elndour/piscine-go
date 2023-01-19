package piscine

func AppendRange(min, max int) []int {
	var tab []int
	if max < min || max == min || max == 1 {
		return tab
	}
	for i := min; i < max; i++ {
		tab = append(tab, min)
		min++
	}
	return tab
}
