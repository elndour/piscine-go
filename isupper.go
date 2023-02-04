package piscine

func IsUpper(s string) bool {
	t := len(s)

	for i := 0; i < t; i++ {
		if s[i] < 'A' || s[i] > 'Z' || t == 0 {
			return false
		}
	}
	return true
}
