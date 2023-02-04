package piscine

func SplitWhiteSpaces(str string) []string {
	tail := 0
	for index, element := range str {
		if (element == ' ' || element == '\n' || element == '\t') && (index != 0 && (str[index-1] != ' ' && str[index-1] != '\n' && str[index-1] != '\t')) {
			tail++
		}
	}

	tab := make([]string, tail+1)
	ind := 0
	a := ""
	for _, element := range str {
		if element == ' ' || element == '\n' || element == '\t' {
			if a != "" {
				tab[ind] = a
				a = ""
				ind++
			}
		} else {
			a = a + string(element)
		}
	}
	if a != "" {
		tab[ind] = a
	}
	return tab
}
