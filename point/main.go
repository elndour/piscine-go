package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func Nombre(n int) {
	if n < 10 {
		for i, j := 0, '0'; i <= n%10; i, j = i+1, j+1 {
			if i == n%10 {
				z01.PrintRune(j)
			}
		}
	} else {
		nbr := n / 10
		mod := n % 10

		Nombre(nbr)
		Nombre(mod)
	}
}

func main() {
	points := point{}
	setPoint(&points)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	Nombre(points.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	Nombre(points.y)
	z01.PrintRune('\n')
}
