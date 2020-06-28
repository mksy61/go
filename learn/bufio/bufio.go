package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		is := bufio.NewScanner(os.Stdin)
		if !is.Scan() {
			fmt.Println("Error!")
			return
		}

		fmt.Printf("%s\n", strings.Repeat("-", 50))
		fmt.Println(is.Text())
		fmt.Printf("% x\n", is.Bytes())
		fmt.Printf("% v\n", is.Bytes())
		fmt.Printf("%s\n", strings.Repeat("-", 50))
	}
}
