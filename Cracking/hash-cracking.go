package cracking

import (
	"bufio"
	"fmt"
	"os"

	"GoCrypt/Hashing"
)

func Parser(hashMethod string, wordlistPath string, hashString string) string {
	if hashMethod == "md5" {
		return md5Crack(hashString, wordlistPath)
	} else if hashMethod == "sha1" {
		return sha1Crack(hashString, wordlistPath)
	}
	return "Password Cracking Method not found"
}

func md5Crack(hashString string, wordlistPath string) string {
	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Println("Error opening wordlist : ", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		hashedWord := hashing.Parser("md5", word)
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
