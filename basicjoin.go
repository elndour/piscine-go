package piscine

func BasicJoin(elems []string) string {
	n := ""
	for _, letter := range elems {
		n = n + letter
	}
	return string(n)
}
