package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	tab := []rune{}

	if n == 0 {
		z01.PrintRune('0')
	}
	if n != 0 {
		for n != 0 { // Tant que
			tab = append(tab, (rune('0' + (n % 10))))
			n = n / 10
		}
		for i := 0; i < len(tab); i++ {
			for j := i + 1; j < len(tab); j++ {
				if tab[i] > tab[j] {
					val := tab[i]
					tab[i] = tab[j]
					tab[j] = val
				}
			}
		}
		for i := 0; i < len(tab); i++ {
			z01.PrintRune(tab[i])
		}
	}
}
