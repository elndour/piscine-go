package piscine

func Map(f func(int) bool, arr []int) []bool {
	tab := make([]bool, len(arr))

	for ind, v := range arr {
		tab[ind] = f(v)
	}
	return tab
}
