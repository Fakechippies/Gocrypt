package main

import (
	"flag"
	"fmt"

	"GoCrypt/Ascii"
	"GoCrypt/Base-Conversion"
	cracking "GoCrypt/Cracking"
	"GoCrypt/Encoding"
	"GoCrypt/Hashing"
	"GoCrypt/ROT-Conversion"
)

func main() {
	// Verbose flag
	vflag := flag.Bool("v", false, "Enable the verbose mode")

	// Basic Flags :
	decryptFlag := flag.Bool("d", false, "Decode the input")
	encryptFlag := flag.Bool("e", false, "Encode the input")

	// Encoding flags
	encFlag := flag.String("enc", "", "Enter the Encoding method")

	// Hashing flags
	hashFlag := flag.String("hash", "", "Enter the hashing method")

	// Base conversion flags
	baseFlag := flag.Bool("base", false, "Base Conversion")
	fromFlag := flag.String("f", "", "current encoding method")
	toFlag := flag.String("t", "", "new encoding method")

	// ROT flags
	rotFlag := flag.Int("rot", 0, "Enter ROT type")

	// ascii flag
	asciiFlag := flag.Bool("ascii", false, "ASCII encoding/decoding")

	// Cracking flags
	crackFlag := flag.String("crack", "", "Enter the Password Crack Hash Type")
	wordlistFlag := flag.String("wordlists", "", "Enter the wordlist path")

	flag.Parse()

	if *encFlag != "" {
		result := encoding.Parser(*encFlag, *decryptFlag, *encryptFlag, flag.Args()[0])
		if result != "" {
			if *decryptFlag {
				if *vflag {
					fmt.Printf("Decoded from %s : %s\n", *encFlag, result)
				} else {
					fmt.Printf("%s\n", result)
				}
			} else if *encryptFlag {
				if *vflag {
					fmt.Printf("Encoded to %s : %s\n", *encFlag, result)
				} else {
					fmt.Printf("%s\n", result)
				}
			}
		}
	} else if *rotFlag != 0 {
		result := rotconversion.Parser(*rotFlag, *encryptFlag, *decryptFlag, flag.Args()[0])
		if *encryptFlag {
			if *vflag {
				fmt.Printf("Encrypted in ROT%d : %s\n", *rotFlag, result)
			} else {
				fmt.Printf("%s\n", result)
			}
		} else if *decryptFlag {
			if *vflag {
				fmt.Printf("Decrypted ROT%d : %s\n", *rotFlag, result)
			} else {
				fmt.Printf("%s\n", result)
			}
		}
	} else if *hashFlag != "" {
		result := hashing.Parser(*hashFlag, flag.Args()[0])
		if *vflag {
			fmt.Printf("'%s' hash of string '%s' : %s\n", *hashFlag, flag.Args()[0], result)
		} else {
			fmt.Printf("%s\n", result)
		}
	} else if *baseFlag {
		result := baseconversion.Parser(*fromFlag, *toFlag, flag.Args()[0])
		if result != "" {
			if *vflag {
				fmt.Printf("'%s' conversion from %s to %s : %s \n", flag.Args()[0], *fromFlag, *toFlag, result)
			} else {
				fmt.Printf("%s\n", result)
			}
		} else {
			fmt.Println("Error in conversion!!!")
		}
	} else if *asciiFlag != false {
		result := ascii.Parser(*decryptFlag, *encryptFlag, flag.Args()[0])
		if *decryptFlag {
			fmt.Printf("Decoded: %s\n", result)
		} else if *encryptFlag {
			fmt.Printf("Encoded: %s\n", result)
		}
	} else if *crackFlag != "" {
		result := cracking.Parser(*crackFlag, *wordlistFlag, flag.Args()[0])
		if result != "" {
			if *vflag {
				fmt.Printf("'%s' %s hash cracked to : %s\n", flag.Args()[0], *crackFlag, result)
			} else {
				fmt.Printf("%s\n", result)
			}
		} else {
			fmt.Println("Error in cracking")
		}
	}
}
