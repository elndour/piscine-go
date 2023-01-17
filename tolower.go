package piscine

func ToLower(s string) string {
	n := []rune(s)
	for i := 0; i < len(s); i++ {
		if n[i] >= 'A' && n[i] <= 'Z' {
			n[i] = n[i] + 32
		}
	}

	return string(n)
}
