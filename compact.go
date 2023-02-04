package piscine

func Compact(ptr *[]string) int {
	var tab []string
	n := 0
	for _, v := range *ptr {
		if v != "" {

			tab = append(tab, v)
			n++
		}
	}

	*ptr = tab
	return n
}
