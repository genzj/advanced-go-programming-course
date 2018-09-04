package main

var x string

var c = make(chan int)

func f2() {
	x = "hello, world"
	// A. use unbuffered way 1
	//c <- 0

	// B. use unbuffered way 2
	//<-c

	// C. use channel closing
	//close(c)
}

func main() {
	go f2()
	// A. use unbuffered way 1
	//<-c

	// B. use unbuffered way 2
	//c <- 0

	// C. use channel closing
	//<-c

	print(x)
}
