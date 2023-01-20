package piscine

func SplitWhiteSpaces(str string) []string {
	len := 0
	for index, element := range str {
		if (element == ' ' || element == '\n' || element == '\t') && (index != 0 && (str[index-1] != ' ' && str[index-1] != '\n' && str[index-1] != '\t')) {
			len++
		}
	}

	array := make([]string, len+1)
	index := 0
	tab := ""
	for _, element := range str {
		if element == ' ' || element == '\n' || element == '\t' {
			if tab != "" {
				array[index] = tab
				tab = ""
				index++
			}
		} else {
			tab = tab + string(element)
		}
	}
	if tab != "" {
		array[index] = tab
	}
	return array
}
