package piscine

func ToUpper(s string) string {
	n := []rune(s)
	for i := 0; i < len(s); i++ {
		if n[i] >= 'a' && n[i] <= 'z' {
			n[i] = n[i] - 32
		}
	}

	return string(n)
}
