package main

import (
	"GoCrypt/Base-Conversion"
	"flag"
	"fmt"
)

func main() {
	baseFlag := flag.String("base", "", "Base Conversion Method")
	fromFlag := flag.String("f", "", "current encoding method")
	toFlag := flag.String("t", "", "new encoding method")
	// Add more flags here

	flag.Parse()

	if *baseFlag != "" {
		result := baseconversion.Parser(*fromFlag, *toFlag, flag.Args()[0])
		if result != "" {
			fmt.Printf("Base Conversion from %s to %s : %s \n", *fromFlag, *toFlag, result)
		} else {
			fmt.Println("Error in conversion!!!")
		}
	}
}
