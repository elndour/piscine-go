package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args

	for i := range a {
		if i >= 1 {
			for _, let := range a[i] {
				z01.PrintRune(let)
			}

			z01.PrintRune('\n')
		}
	}
}
