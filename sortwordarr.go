package piscine

func SortWordArr(a []string) {
	for _, v := range a {
		if (v >= "0" && v <= "9") || (v >= "a" && v <= "z") || (v >= "A" && v <= "Z") {
			for i := 0; i < len(a); i++ {
				for j := i + 1; j < len(a); j++ {
					if a[i] > a[j] {
						b := a[i]
						a[i] = a[j]
						a[j] = b
					}
				}
			}
		}
	}
}
