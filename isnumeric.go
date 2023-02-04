package piscine

func IsNumeric(s string) bool {
	t := len(s)
	if t == 0 {
		return false
	}
	for i := 0; i < t; i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
