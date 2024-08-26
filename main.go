package main

import (
	"GoCrypt/Encoding"
	"flag"
	"fmt"
)

func main() {
	fromFlag := flag.String("f", "", "current encoding method")
	toFlag := flag.String("t", "", "new encoding method")
	// Add more flags here

	flag.Parse()

	if *fromFlag != "" && *toFlag != "" {
		result := encoding.Parser(*fromFlag, *toFlag, flag.Args()[0])
		if result != "" {
			fmt.Printf("Converted from %s to %s :- %s \n", *fromFlag, *toFlag, result)
		} else {
			fmt.Println("Error in conversion!!!")
		}
	}
}
