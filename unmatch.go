package piscine

func Unmatch(a []int) int {
	for _, v := range a {
		compt := 0
		for _, s := range a {
			if v == s {
				compt++
			}
		}
		if compt == 1 || compt%2 == 1 {
			return v
		}
	}
	return -1
}
