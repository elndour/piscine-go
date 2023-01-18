package piscine

func Capitalize(s string) string {
	n := []rune(s)

	if !checkMg(n[0]) {
		n[0] = n[0] - 32 + 32
	}
	for i := range n {
		if (n[i] == ' ' || n[i] == '+') && checkMg(n[i+1]) {
			n[i+1] = n[i+1] - 32
		} else if (n[i] == ' ' || n[i] == '+') && (n[i+1] >= '0' && n[i+1] <= '9') {
			n[i+1] = n[i+1] + 32 - 32
		}
	}

	return string(n)
}

func checkMg(b rune) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}

	return false
}
