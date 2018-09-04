package main

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"regexp"
	"time"
)

func grepInFile(regex *regexp.Regexp, fileName string, out chan<- string) {
	defer close(out)

	// open file
	file, err := os.Open(fileName)

	// check possible open error
	if err != nil {
		out <- fmt.Sprintf("fatal error: %v\n", err)
		return
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
		out <- fmt.Sprintf("scanner error: %v\n", err)
	}
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		regex, err := regexp.Compile(c.Query("r"))
		if err != nil {
			c.JSON(401, gin.H{
				"error": fmt.Sprintf("invalid regex (%s)", err),
			})
		}
		fileName := c.Query("filename")

		fmt.Printf("run grep \"%v\" on file %v\n", regex, fileName)

		ch := make(chan string)
		go grepInFile(regex, fileName, ch)

		data := []string{}

		for text := <-ch; text != ""; text = <-ch {
			data = append(data, text)
		}

		c.JSON(200, gin.H{
			"result": data,
		})
	})
	r.Run("127.0.0.1:9000") // listen and serve on 0.0.0.0:8080
}

// try:
// http://127.0.0.1:9000/?r=HTTP%2F1.1%22%2050.&filename=data-text%2Fapache_logs.txt