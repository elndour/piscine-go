package piscine

func ConcatParams(args []string) string {
	n := ""

	for i, letter := range args {
		n = n + letter
		if i != (len(args) - 1) {
			n = n + "\n"
		}
	}

	return n
}
