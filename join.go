package piscine

func Join(strs []string, sep string) string {
	n := ""
	for j, letter := range strs {

		n = n + letter
		if j != (len(strs) - 1) {
			n = n + sep
		}

	}
	return n
}
