package main

import (
	"flag"
	"fmt"

	"GoCrypt/Encoding"
)

func main() {
	encFlag := flag.String("enc", "", "Enter the Encoding method")
	decryptFlag := flag.Bool("d", false, "Decode the input")
	encryptFlag := flag.Bool("e", false, "Encode the input")
	// Add more flags here

	flag.Parse()

	result := encoding.Parser(*encFlag, *decryptFlag, *encryptFlag, flag.Args()[0])
	if result != "" {
		fmt.Printf("%s : %s\n", *encFlag, result)
	}
}
