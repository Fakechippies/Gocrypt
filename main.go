package main

import (
	"flag"
	"fmt"

	"GoCrypt/Base-Conversion"
	"GoCrypt/Encoding"
	"GoCrypt/Hashing"
)

func main() {
	// Encoding flags
	encFlag := flag.String("enc", "", "Enter the Encoding method")
	decryptFlag := flag.Bool("d", false, "Decode the input")
	encryptFlag := flag.Bool("e", false, "Encode the input")

	// Hashing flags
	hashFlag := flag.String("hash", "", "Enter the hashing method")

	// Base conversion flags
	baseFlag := flag.String("base", "", "Base Conversion Method")
	fromFlag := flag.String("f", "", "current encoding method")
	toFlag := flag.String("t", "", "new encoding method")

	// Add more flags here if needed

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
	} else if *baseFlag != "" {
		result := baseconversion.Parser(*fromFlag, *toFlag, flag.Args()[0])
		if result != "" {
			fmt.Printf("Base Conversion from %s to %s : %s \n", *fromFlag, *toFlag, result)
		} else {
			fmt.Println("Error in conversion!!!")
		}
	}
}
