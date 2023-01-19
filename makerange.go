package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}
	tab := make([]int, max-min)

	for i := 0; i < max-min; i++ {
		tab[i] = min + i
	}
	return tab
}
