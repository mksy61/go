package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: bigff filesize in Bytes")
		return
	}

	fileSize, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		return
	}

	fileNames := []byte{}

	for _, file := range files {
		if file.Size() > fileSize {
			fmt.Printf("%16v\t%v\n", file.Name(), file.Size())
			fileNames = append(fileNames, file.Name()...)
			fileNames = append(fileNames, byte('\n'))
		}
	}
	fmt.Println(strings.Repeat("-", 50))
	fmt.Println(string(fileNames))
	err = ioutil.WriteFile("bigfiles.txt", fileNames, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
}
