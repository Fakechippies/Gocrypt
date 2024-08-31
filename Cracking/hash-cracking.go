package cracking

import (
	"bufio"
	"fmt"
	"os"

	"GoCrypt/Hashing"
)

func Parser(hashMethod string, wordlistPath string, hashString string) string {
	return Crack(hashString, wordlistPath, hashMethod)
}

func Crack(hashString string, wordlistPath string, hashMethod string) string {
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Error opening wordlist : ", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		hashedWord := hashing.Parser(hashMethod, word)
		if hashedWord == hashString {
			return word
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading wordlist:", err)
		return ""
	}

	fmt.Println("Cannot crack password")
	return ""
}
