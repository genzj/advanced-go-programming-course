package main

import (
	"bufio"
	"fmt"
	"github.com/pkg/profile"
	"math/rand"
	"os"
	"reflect"
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

func main() {
	defer profile.Start(profile.MemProfile).Stop()

	total := len(os.Args[2:])
	chans := make([]<-chan string, 0, total)
	closedWorkers := make(map[int]bool, total)

	for _, filename := range os.Args[2:] {
		chans = append(chans, makeGrepWorker(filename))
	}

	// one more channel for timeout
	cases := make([]reflect.SelectCase, total+1)

	for i, ch := range chans {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}
	cases[total] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(time.After(2 * time.Second))}

	for len(closedWorkers) < total {
		chosen, value, recvOK := reflect.Select(cases)
		if chosen == total {
			fmt.Printf("timeout\n")
			break
		} else if recvOK {
			msg := value.String()
			fmt.Printf("from worker %d\n%s\n", chosen+1, msg)
		} else if _, in := closedWorkers[chosen]; !in {
			fmt.Printf("worker %d exited\n", chosen+1)
			closedWorkers[chosen] = recvOK
		}
	}
}
