package rotconversion

func Parser(rot_type int, encrypt bool, decrypt bool, input string) string {
	if encrypt {
		return rot(input, rot_type)
	} else if decrypt {
		if rot_type%13 == 0 {
			return rot(input, rot_type)
		} else if rot_type <= 25 && rot_type >= 1 {
			return rot(input, 26-rot_type)
		} else {
			return rot(input, (-1)*(rot_type-26))
		}
	}
	return "Error in rot type provided or input"
}

func rot(input string, shift int) string {
	if shift > 26 {
		shift = shift - 26
	}

	letters := [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	var lengthText = len([]rune(input))

	var result string

	for i := 0; i < lengthText; i++ {
		for a := 0; a < 26; a++ {
			if string(input[i]) == letters[a] {
				if a+shift > 25 {
					index := (a + shift) - 26
					result = result + letters[index]
				} else {
					result = result + letters[a+shift]
				}
			}
		}
	}

	return result

}
