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

func grepInFile(fileName string, out chan<- string) {
	defer close(out)

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
			out <- fmt.Sprintf("%s:%d %s\n", fileName, lineNo, scanner.Text())
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

func makeGrepWorker(fileName string) <-chan string {
	ch := make(chan string)
	go grepInFile(fileName, ch)
	return ch
}

func output(s string, worker int) (exited bool) {
	if s != "" {
		fmt.Printf("from worker %d\n%s\n", worker, s)
		return false
	} else {
		return true
	}
}

func main() {
	// run grepInFile for each file name
	// start from 2 files
	ch1 := makeGrepWorker(os.Args[2])
	ch2 := makeGrepWorker(os.Args[3])
	finished1 := false
	finished2 := false

	for !finished1 || !finished2 {
		select {
		case t := <-ch1:
			finished1 = output(t, 1)
		case t := <-ch2:
			finished2 = output(t, 2)
		}
	}
}
