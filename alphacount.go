package piscine

func AlphaCount(s string) int {
	c := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'a' || s[i] == 'b' || s[i] == 'c' || s[i] == 'd' || s[i] == 'e' || s[i] == 'f' || s[i] == 'g' || s[i] == 'h' || s[i] == 'i' || s[i] == 'j' ||
			s[i] == 'k' || s[i] == 'l' || s[i] == 'm' || s[i] == 'o' || s[i] == 'p' || s[i] == 'q' || s[i] == 'r' || s[i] == 's' || s[i] == 't' || s[i] == 'u' ||
			s[i] == 'v' || s[i] == 'w' || s[i] == 'x' || s[i] == 'y' || s[i] == 'z' || s[i] == 'A' || s[i] == 'B' || s[i] == 'C' || s[i] == 'D' || s[i] == 'E' || s[i] == 'F' || s[i] == 'G' || s[i] == 'H' || s[i] == 'I' || s[i] == 'J' ||
			s[i] == 'K' || s[i] == 'L' || s[i] == 'M' || s[i] == 'O' || s[i] == 'P' || s[i] == 'Q' || s[i] == 'R' || s[i] == 'S' || s[i] == 'T' || s[i] == 'U' ||
			s[i] == 'V' || s[i] == 'W' || s[i] == 'X' || s[i] == 'Y' || s[i] == 'Z' {
			c = c + 1
		}
	}
	return c
}
