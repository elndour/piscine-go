package main

import "github.com/01-edu/z01"

func main() {
	for letr := 'z'; letr >= 'a'; letr-- {
		z01.PrintRune(letr)
	}

	z01.PrintRune('\n')
}
