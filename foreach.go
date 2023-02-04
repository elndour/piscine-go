package piscine

func ForEach(f func(int), a []int) {
	for _, number := range a {
		f(number)
	}
}
