package piscine

import "github.com/01-edu/z01"

func SortIntegerTable(table []int) {
	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {

				n := table[i]
				table[i] = table[j]
				table[j] = n
			}
		}
	}
	for i := 0; i < len(table); i++ {
		z01.PrintRune(rune(table[i] + 48))
	}
}
