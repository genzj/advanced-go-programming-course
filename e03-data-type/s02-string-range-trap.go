package main

import (
	"fmt"
)

func main() {
	const s = "Chinese: 中文"
	for index, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	for index := 0; index < len(s); index++ {
		fmt.Printf("s[%d] = %x (%c)\n", index, s[index], s[index])
	}

	//for index, byteValue := range []byte(s) {
	//	fmt.Printf("byteValue[%d] = %x (%c)\n", index, byteValue, byteValue)
	//}
}
