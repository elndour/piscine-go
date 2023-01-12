package main

import "github.com/01-edu/z01"

func main() {
	for ch := 'z'; ch >= 'a'; ch-- {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')
	// alpha := "zyxwvutsrqponmlkjihgfedcba"
	// for _, vh := range alpha {
	// 	z01.PrintRune(vh)
	// }
	// z01.PrintRune('\n')
}
