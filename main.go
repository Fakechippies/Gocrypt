package main

import (
	"flag"
	"fmt"

	"GoCrypt/Encoding"
	"GoCrypt/Hashing"
)

func main() {
	encFlag := flag.String("enc", "", "Enter the Encoding method")
	decryptFlag := flag.Bool("d", false, "Decode the input")
	encryptFlag := flag.Bool("e", false, "Encode the input")

	hashFlag := flag.String("hash", "", "Enter the hashing method")
	// Add more flags here

	flag.Parse()

	if *encFlag != "" {
		result := encoding.Parser(*encFlag, *decryptFlag, *encryptFlag, flag.Args()[0])
		if result != "" {
			if *decryptFlag {
				fmt.Printf("Decoded from %s : %s\n", *encFlag, result)
			} else if *encryptFlag {
				fmt.Printf("Encoded to %s : %s\n", *encFlag, result)
			}
		}
	} else if *hashFlag != "" {
		result := hashing.Parser(*hashFlag, flag.Args()[0])
		fmt.Printf("MD5 of string '%s' : %s\n", flag.Args()[0], result)
	}
}
