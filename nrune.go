package piscine

func NRune(s string, n int) rune {
	run := []rune(s)
	long := len(run)
	if n <= 0 {
		return 0
	} else if n > long {
		return 0
	} else {
		return run[n-1]
	}
}
