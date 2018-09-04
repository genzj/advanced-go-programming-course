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

func grepInFile(fileName string, done *sync.WaitGroup) {
	defer done.Done()

	regex := getRegexSingleton()

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

var singleton *regexp.Regexp
var initializeOnce = &sync.Once{}

func getRegexSingleton() *regexp.Regexp {
	initializeOnce.Do(func() {
		compiled, err := regexp.Compile(os.Args[1])
		if err != nil {
			fmt.Printf("regex err: %v\n", err)
			os.Exit(FatalError)
		} else {
			fmt.Printf("singleton created %v\n", &compiled)
		}
		singleton = compiled
	})

	return singleton
}

func main() {
	done := &sync.WaitGroup{}
	// run grepInFile for each file name
	for _, fileName := range os.Args[2:] {
		done.Add(1)
		go grepInFile(fileName, done)
	}

	done.Wait()
}
