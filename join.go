package piscine

func Join(strs []string, sep string) string {
	var a string

	for i, v := range strs {

		a = a + v
		if i == len(strs)-1 {
			break
		}
		a = a + sep
	}
	return a
}
