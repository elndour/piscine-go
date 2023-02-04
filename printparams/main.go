package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args

	for i := 1; i < len(a); i++ {
		for _, v := range a[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')

	}
}
