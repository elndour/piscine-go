package piscine

func Permute(tab []string, ind int, val string) {
	tab[ind] = val
}

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) > 0 {
				temp := a[i]
				Permute(a, i, a[j])
				Permute(a, j, temp)
			}
		}
	}
}
