package encoding

import (
	b64 "encoding/base64"
	"fmt"
)

func Parser(encMethod string, decode bool, encode bool, input string) string {
	if encMethod == "base64" {
		return base64Conv(decode, encode, input)
	}
	return "Unsupported encoding method"
}

func base64Conv(decode bool, encode bool, input string) string {
	var result string

	if encode {
		result = b64.StdEncoding.EncodeToString([]byte(input + "\n"))
	}
	if decode {
		data, err := b64.StdEncoding.DecodeString(input)
		if err != nil {
			fmt.Println("Error decoding base64 : ", err)
			return ""
		}
		result = string(data)
	}
	return result
}
