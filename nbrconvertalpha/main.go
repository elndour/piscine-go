package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	var noZero, negative int
	for g, i := range s {
		count := 0
		if g == 0 && (i == '-' || i == '+') {
			if i == '-' {
				negative++
			}
			continue
		}
		if i < '0' || i > '9' {
			return 0
		}
		for j := '0'; j < i; j++ {
			count++
		}
		noZero = noZero*10 + count
	}

	if negative == 1 {
		return -noZero
	} else {
		return noZero
	}
}

func main() {
	a := os.Args

	var t int
	for index := range a {
		t = index + 1
	}

	if t > 1 {
		letr := 96

		if os.Args[1] == "--upper" {
			letr = 64
		}
		for index, element := range a {
			if index > 0 {

				nbr := Atoi(element) + letr
				if (nbr >= 97 && nbr <= 122) || (nbr >= 65 && nbr <= 90) {
					z01.PrintRune(rune(nbr))
				} else if os.Args[1] != "--upper" {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')

	}
	if os.Args[0] == "go run ." {
		z01.PrintRune('\n')
	}
}
