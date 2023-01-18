package piscine

func Capitalize(s string) string {
	n := []rune(s)
	first := true

	for i := range n {
		if checkMg(n[i]) && first {
			if n[i] >= 'a' && n[i] <= 'z' {
				n[i] = n[i] - 32
			}
			first = false

		} else if n[i] >= 'A' && n[i] <= 'Z' {
			n[i] = n[i] + 32
		} else if !checkMg(n[i]) {
			first = true
		}
	}

	return string(n)
}

func checkMg(b rune) bool {
	if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9') {
		return true
	}

	return false
}
