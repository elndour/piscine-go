package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	} else {
		tab := []rune{}

		if n < 0 {
			z01.PrintRune('-')
		}

		nbr := n
		for nbr != 0 {
			if nbr < 0 {
				tab = append(tab, rune('0'-(nbr%10)))
			} else {
				tab = append(tab, rune('0'+(nbr%10)))
			}
			nbr = nbr / 10
		}
		tail := len(tab)
		for i := tail - 1; i >= 0; i-- {
			z01.PrintRune(tab[i])
		}

	}
}
