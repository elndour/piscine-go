package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args

	for i := len(a) - 1; i >= 1; i-- {
		for _, let := range a[i] {
			z01.PrintRune(let)
		}

		z01.PrintRune('\n')
	}
}
