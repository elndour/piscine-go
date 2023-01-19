package piscine

func AppendRange(min, max int) []int {
	var tab []int = nil
	if max <= min {
		return tab
	}
	if max == 0 && min == 0 {
		return tab
	}
	for i := min; i < max; i++ {
		tab = append(tab, min)
		min++
	}
	return tab
}
