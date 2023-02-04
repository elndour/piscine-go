package piscine

func Max(a []int) int {
	max := 0
	if len(a) == 0 {
		return 0
	} else {
		for _, v := range a {
			if max < v {
				max = v
			}
		}
	}
	return max
}
