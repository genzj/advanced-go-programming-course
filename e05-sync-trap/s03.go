package main

type T struct {
	msg string
}

var g2 *T

func setup2() {
	t := new(T)
	t.msg = "hello, world"
	g2 = t
}

func main() {
	go setup2()
	for g2 == nil {
	}
	print(g2.msg)
}
