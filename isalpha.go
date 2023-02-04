package piscine

func IsAlpha(s string) bool {
	t := len(s)
	a := " "

	if s == a {
		return true
	}
	for i := 0; i < t; i++ {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') && (s[i] < '0' || s[i] > '9') {
			return false
		}
	}
	return true
}
