package main

import (
	"flag"
	"fmt"

	"GoCrypt/Base-Conversion"
	"GoCrypt/Encoding"
	"GoCrypt/Hashing"
)

func main() {
	// Verbose flag
	vflag := flag.Bool("v", false, "Enable the verbose mode")

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
	}
}
