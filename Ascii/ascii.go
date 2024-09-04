package ascii

import (
	"strconv"
)

func Parser(decode bool, encode bool, input string) string {
	if encode {
		return encoding(input)
	} else if decode {
		return decoding(input)
	}
	return "invalid input provided"
}

func encoding(input string) string {
	temp := []byte(input)
	result := ""
	for i := 0; i < len(temp); i++ {
		result += strconv.Itoa(int(temp[i]))
		result += " "
	}
	return result
}

func decoding(input string) string {
	temp := []byte(input)
	result := ""
	for i := 0; i < len(temp); i++ {
		result += string(temp[i])
		result += " "
	}
	return result
}
