package main

import (
	"flag"
	"fmt"

	"GoCrypt/Base-Conversion"
	cracking "GoCrypt/Cracking"
	"GoCrypt/Encoding"
	"GoCrypt/Hashing"
	"GoCrypt/ROT-Conversion"
)

func main() {
	// Verbose flag
	vflag := flag.Bool("v", false, "Enable the verbose mode")

	// Help flag
	hflag := flag.Bool("h", false, "Help")

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

	// Cracking flags
	crackFlag := flag.String("crack", "", "Enter the Password Crack Hash Type")
	wordlistFlag := flag.String("w", "", "Enter the wordlist path")

	// Encryption flags

	flag.Parse()

	if *hflag {
		help := `
		go run main.go [options] [input]

		-h :- Help
		-v :- Verbose
		-d :- Decrypt
		-e :- Encrypt
		-base :- Base Inter-Conversion (base2, base10, base16)
		-enc :- Encodings (base64, base58, base32, UR)	
		-w :- Wordlist
		-crack :- Password hash cracking (md5, sha1, sha256, sha224, sha384, sha512)
		-hash :- Hashing (md5, sha1, sha256, sha224, sha384, sha512)
		-rot :- ROT (rot13, rot25, rot47 etc)
		
		Examples
		go run main.go -crack md5 -w /usr/share/wordlists/rockyou.txt 5f4dcc3b5aa765d61d8327deb882cf99
		go run main.go -hash sha512 GoCrypt
		go run main.go -enc base64 -e Gocrypt
		go run main.go -base -f base10 -t base2 777
		go run main.go -rot 13 -e hello 
		`
		fmt.Println(help)
	}

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
