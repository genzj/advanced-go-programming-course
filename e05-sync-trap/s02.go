package main

var s string
var done bool

func setup() {
	s = "hello, world"
	done = true
}

func main() {
	go setup()
	for !done {
	}
	print(s)
}
