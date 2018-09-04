package main

import (
	"fmt"
	"os"
)

const FatalError = 2

func grepInFile(regex, fileName string) {
	// open file
	____, ____ := os.____(fileName)

	// check possible open error
	if ____ != nil {
		fmt.Printf("fata error: %v\n", err)
		os.Exit(FatalError)
	}

	// ensure file will be closed at return
	____ ____.Close()

	lineNo := 1

	// create scanner
	scanner := ____

	// while scanner found something
	for ____() {

		// match regex against scanner Text
		matched, err := ____
		if err != nil {
			fmt.Printf("regex error: %v\n", err)
			os.Exit(FatalError)
		} else if matched {
			// output the line
			fmt.Printf("%s:%d %s\n", fileName, lineNo, ____)
		}
		lineNo++
	}

	// handle possible scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("scanner error: %v\n", err)
		os.Exit(FatalError)
	}
}

func main() {
	// extract regex from arguments, assume the 1st argument is regex, followed by one or more file names
	regex := ____

	// run grepInFile for each file name
	for ____ {
		grepInFile(regex, fileName)
	}
}
