package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	var b string
	for i, v := range a {

		b = b + string(v)
		if i != len(a)-1 {
			b = b + "\n"
		}

	}
	for i := 0; i < len(b); i++ {
		z01.PrintRune(rune(b[i]))
	}
	z01.PrintRune('\n')
}
