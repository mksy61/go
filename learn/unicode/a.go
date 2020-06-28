package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	bs := []byte{240, 159, 152, 132, 240, 159, 152, 135, 226, 143, 128}

	fmt.Printf("%s\t\t% x", bs, bs)
	fmt.Println()
	fmt.Println(utf8.RuneCount(bs))

}
