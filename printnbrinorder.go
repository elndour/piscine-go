package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	a := []int{}
	if n == 0 {
		z01.PrintRune(rune('0'))
	}
	for n > 0 {

		a = append(a, n%10)

		n = n / 10
	}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {

				n := a[i]
				a[i] = a[j]
				a[j] = n
			}
		}
	}
	for i := 0; i < len(a); i++ {
		z01.PrintRune(rune(a[i] + 48))
	}
}
