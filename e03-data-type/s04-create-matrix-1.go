package main

import "fmt"

func main() {
	x := 2
	y := 4

	table := make([][]int, x)
	for i := range table {
		table[i] = make([]int, y)
	}
	fmt.Printf("%+v\n", table)
}
