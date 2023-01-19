package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args
	tab := []rune(a[0])
	for i := range tab {
		if i >= 2 {
			z01.PrintRune(tab[i])
		}
	}
	z01.PrintRune('\n')
}
