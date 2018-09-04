package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"sync"
	"time"
)

const FatalError = 2

var regex *regexp.Regexp

func grepInFile(fileName string, done *sync.WaitGroup) {
	defer done.Done()

	// open file
	file, err := os.Open(fileName)

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
		matched := regex.MatchString(scanner.Text())
		if matched {
			// output the line
			fmt.Printf("%s:%d %s\n", fileName, lineNo, scanner.Text())
			time.Sleep(time.Duration(rand.Float32()*3) * time.Second)
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
	done := &sync.WaitGroup{}

	// extract regex from arguments, assume the 1st argument is regex, followed by one or more file names
	compiled, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Printf("regex err: %v\n", err)
		os.Exit(FatalError)
	}

	regex = compiled

	// run grepInFile for each file name
	for _, fileName := range os.Args[2:] {
		done.Add(1)
		go grepInFile(fileName, done)
	}

	done.Wait()
}
