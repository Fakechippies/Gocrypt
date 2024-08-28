package main

import (
	"flag"
	"fmt"

	"GoCrypt/Base-Conversion"
	"GoCrypt/Encoding"
	"GoCrypt/Hashing"
	"GoCrypt/ROT-Conversion"
)

func main() {
	// Encoding flags
	encFlag := flag.String("enc", "", "Enter the Encoding method")
	decryptFlag := flag.Bool("d", false, "Decode the input")
	encryptFlag := flag.Bool("e", false, "Encode the input")

	// Hashing flags
	hashFlag := flag.String("hash", "", "Enter the hashing method")

	// Base conversion flags
	baseFlag := flag.Bool("base", false, "Base Conversion")
	fromFlag := flag.String("f", "", "current encoding method")
	toFlag := flag.String("t", "", "new encoding method")

	// Add more flags here if needed
	rotFlag := flag.Int("rot", 0, "Enter ROT type")
	encrFlag := flag.Bool("encr", false, "encrypt the input")
	decrFlag := flag.Bool("decr", false, "decrypt the input")

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
	} else if *rotFlag != 0 {
		result := ROT.Parser(*rotFlag, *encrFlag, *decrFlag, flag.Args()[0])
		if *encrFlag {
			fmt.Printf("Encrypted in ROT%d : %s", *rotFlag, result)
		} else if *decrFlag {
			fmt.Printf("Decrypted ROT%d : %s", *rotFlag, result)
		}
	} else if *hashFlag != "" {
		result := hashing.Parser(*hashFlag, flag.Args()[0])
		fmt.Printf("MD5 of string '%s' : %s\n", flag.Args()[0], result)
	} else if *baseFlag {
		result := baseconversion.Parser(*fromFlag, *toFlag, flag.Args()[0])
		if result != "" {
			fmt.Printf("'%s' conversion from %s to %s : %s \n", flag.Args()[0], *fromFlag, *toFlag, result)
		} else {
			fmt.Println("Error in conversion!!!")
		}
	}
}
