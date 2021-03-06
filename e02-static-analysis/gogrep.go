package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const FatalError = 2

func grepInFile(regex, file_name string) {
	// open file
	file, err := os.Open(file_name)

	// check possible open error
	if err != nil {
		fmt.Printf("fata error: %v\n", err)
		os.Exit(FatalError)
	}

	// ensure file will be closed at return
	defer file.Close()

	lineNo := 1

	// create scanner
	scanner := bufio.NewScanner(file)

	// while scanner found something
	for scanner.Scan() {

		// match regex against scanner Text
		matched, err := regexp.MatchString(regex, scanner.Text())
		if (err != nil) {
			fmt.Printf("regex error: %v\n", err)
			os.Exit(FatalError)
		}else
		if matched {
			// output the line
			fmt.Printf("%s:%d %s\n", file_name, lineNo, scanner.Text())
		}
		lineNo++
	}

	// handle possible scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("scanner error: %v\n", err)
		os.Exit(FatalError)
	}
}

func main() {
	// extract regex from arguments, assume the 1st argument is regex, followed by one or more file names
	regex := os.Args[1]

	// run grepInFile for each file name
	for _, fileName := range os.Args[2:] {
		grepInFile(regex, fileName)
	}
}
