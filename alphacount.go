package piscine

func AlphaCount(s string) int {
	c := 0
	for _, letter := range s {
		if letter >= 'a' && (letter) <= 'z' || (letter) >= 'A' && (letter) <= 'Z' {
			c++
		}
	}

	return c
}
