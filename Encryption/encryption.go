package encryption

import (
	"crypto/aes"
	"crypto/des"
)

func Parser(encMethod string, encrypt bool, decrypt bool, input string, key string) string {
	if encMethod == "aes" {
		return aesConversion(encrypt, decrypt, input, key)
	}
	if encMethod == "3des" {
		return tripleDesConversion(encrypt, decrypt, input, key)
	}
	return ""
}

func aesConversion(encrypt bool, decrypt bool, input string, key string) string {
	if encrypt {
		aes, err := aes.NewCipher([]byte(key))
		if err != nil {
			panic(err)
		}

		ciphertext := make([]byte, len(input))
		aes.Encrypt(ciphertext, []byte(input))

		return string(ciphertext)
	} else if decrypt {
		aes, err := aes.NewCipher([]byte(key))
		if err != nil {
			panic(err)
		}

		plaintext := make([]byte, len(input))
		aes.Decrypt(plaintext, []byte(input))
		return string(plaintext)
	}
	return ""
}

func tripleDesConversion(encrypt bool, decrypt bool, input string, key string) string {
	cipher, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	if encrypt {
		ciphertext := make([]byte, len(input))

		cipher.Encrypt(ciphertext, []byte(input))
		return string(ciphertext)
	} else if decrypt {
		plaintext := make([]byte, len(input))

		cipher.Decrypt(plaintext, []byte(input))
		return string(plaintext)
	}
	return ""
}
