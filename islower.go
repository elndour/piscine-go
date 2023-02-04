package piscine

func IsLower(s string) bool {
	t := len(s)

	for i := 0; i < t; i++ {
		if s[i] < 'a' || s[i] > 'z' || t == 0 {
			return false
		}
	}
	return true
}
