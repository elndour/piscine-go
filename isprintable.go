package piscine

func IsPrintable(s string) bool {
	a := " "
	t := len(s)
	if t == 0 {
		return false
	}
	for i := 0; i < t; i++ {
		if s[i] < a[0] {
			return false
		}
	}
	return true
}
