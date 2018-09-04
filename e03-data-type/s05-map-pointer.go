package main

import "fmt"

type data struct {
	name string
}

func main() {
	m := map[string]data{"x": {"one"}}
	fmt.Println(m["x"].name)
	m["x"].name = "two"
}
