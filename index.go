package piscine

func Index(s string, toFind string) int {
	a := len(s)
	b := len(toFind)

	if b > a {
		return -1
	}

	for i := 0; i <= a-b; i++ {
		if s[i:i+b] == toFind {
			return i
		}
	}

	return -1
}
