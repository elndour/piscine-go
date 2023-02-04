package piscine

func Index(s string, toFind string) int {
	ts := len(s)
	tto := len(toFind)
	dif := ts - tto

	if tto > ts {
		return -1
	} else {
		for i := 0; i <= dif; i++ {
			if s[i:i+tto] == toFind {
				return i
			}
		}
	}

	return -1
}
