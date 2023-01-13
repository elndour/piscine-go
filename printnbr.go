package piscine

import "github.com/01-edu/z01"

func intToRune(num int) rune {
	if num < 0 {
		num = num * -1
	}
	if num == 0 {
		return '0'
	}
	if num == 1 {
		return '1'
	}
	if num == 2 {
		return '2'
	}
	if num == 3 {
		return '3'
	}
	if num == 4 {
		return '4'
	}
	if num == 5 {
		return '5'
	}
	if num == 6 {
		return '6'
	}
	if num == 7 {
		return '7'
	}
	if num == 8 {
		return '8'
	}
	if num == 9 {
		return '9'
	}
	return rune(num)
}

func PrintNbr(n int) {
	var neg int = 1
	if n < 0 {
		neg = -1
		z01.PrintRune('-')
	}
	if n >= 10 || n <= -10 {
		PrintNbr(n / 10 * neg)
	}
	z01.PrintRune(intToRune(n % 10))
}
