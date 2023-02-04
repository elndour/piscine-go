package piscine

func StrRev(s string) string {
	invrs := []rune(s)
	for i, j := 0, len(invrs)-1; i < j; i, j = i+1, j-1 {
		invrs[i], invrs[j] = invrs[j], invrs[i]
	}
	return string(invrs)
}
