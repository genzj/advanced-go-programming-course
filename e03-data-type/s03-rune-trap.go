package main

import (
	"fmt"
	"strings"
)

func main() {
	const s1 = "\u00e0"
	const s2 = "\u0061\u0300"

	fmt.Printf("%s == %s -> %v\n", s1, s2, strings.Compare(s1, s2))
}
