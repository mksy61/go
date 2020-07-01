package main

import (
	"fmt"
	"unicode"
)

func main() {
	t := unicode.TurkishCase

	const lci = 'ı'
	fmt.Printf("%#U\n", t.ToLower(lci))
	fmt.Printf("%#U\n", t.ToUpper(lci))

	const uci = 'Ğ'
	fmt.Printf("%#U\n", t.ToLower(uci))
	fmt.Printf("%#U\n", t.ToUpper(uci))
}
