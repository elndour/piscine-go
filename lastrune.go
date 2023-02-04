package piscine

func LastRune(s string) rune {
	run := []rune(s)
	long := len(s)

	return run[long-1]
}
