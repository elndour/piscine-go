package piscine

func Rot14(s string) string {
	var a string
	for _, v := range s {
		if v >= 97 && v <= 122 {
			if v >= 97 && v <= 108 {
				a = a + string(v+14)
			} else if v >= 109 && v <= 122 {
				a = a + string(v-12)
			}
		} else if v >= 65 && v <= 90 {
			if v >= 65 && v <= 76 {
				a = a + string(v+14)
			} else if v >= 77 && v <= 90 {
				a = a + string(v-12)
			}
		} else {
			a = a + string(v)
		}
	}

	return a
}
