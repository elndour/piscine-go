package main

import (
	"os"
)

func OperationValid(test string) bool {
	op := []string{"+", "-", "*", "/", "%"}
	for _, v := range op {
		if v == test {
			return true
		}
	}
	return false
}

func Atoi(s string) (int, bool) {
	nbr := 0
	ind := 0
	if len(s) == 0 {
		ind = ind + 0
	} else if s[0] == '-' || s[0] == '+' {
		ind = ind + 1
	}
	for ; ind < len(s); ind++ {
		if s[ind] < '0' || s[ind] > '9' {
			return 0, false
		}
		if s[0] == '-' {
			nbr = nbr*10 - int(s[ind]-'0')
		} else {
			nbr = nbr*10 + int(s[ind]-'0')
		}
	}

	return nbr, true
}

func Itoi(n int) string {
	if n == 0 {
		return "0"
	}
	tab := []rune{}
	for n != 0 {

		tab = append(tab, (rune(n%10) + '0'))
		n = n / 10

	}
	for i, j := 0, len(tab)-1; i < j; i, j = i+1, j-1 {
		tab[i], tab[j] = tab[j], tab[i]
	}

	return string(tab)
}

func main() {
	args := os.Args[1:]
	if len(args) > 3 || len(args) < 3 {
		return
	} else if len(args) == 3 {
		if !(OperationValid(args[1])) {
			return
		} else {
			nbr1, errx := Atoi(args[0])
			nbr2, erry := Atoi(args[2])
			if (nbr1 >= -21474836448 && nbr1 <= 21474836448) && (nbr2 >= -21474836448 && nbr2 <= 21474836448) {

				var result int
				if args[1] == "%" && nbr2 == 0 {
					os.Stderr.WriteString("No modulo by 0\n")
					return
				}
				if args[1] == "/" && nbr2 == 0 {
					os.Stderr.WriteString("No division by 0\n")
					return
				}
				if args[1] == "+" && errx && erry {
					result = nbr1 + nbr2
				} else if args[1] == "-" && errx && erry {
					result = nbr1 - nbr2
				} else if args[1] == "*" && errx && erry {
					result = nbr1 * nbr2
				} else if args[1] == "/" && nbr2 != 0 && errx && erry {
					result = nbr1 / nbr2
				} else if args[1] == "%" && errx && erry && nbr2 != 0 {
					result = nbr1 % nbr2
				}

				if result == 0 && errx && erry {
					os.Stderr.WriteString(Itoi(result))
					os.Stderr.WriteString("\n")

				} else if result < 0 && errx && erry {
					result *= -1
					os.Stderr.WriteString("-")
					os.Stderr.WriteString(Itoi(result))
					os.Stderr.WriteString("\n")

				} else if result > 0 && errx && erry {
					os.Stderr.WriteString(Itoi(result))
					os.Stderr.WriteString("\n")

				} else {
					return
				}
			}
		}
	}
}
