package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Operating System:\t", runtime.GOOS)
	fmt.Println("CPU Number:\t\t", runtime.NumCPU())
	fmt.Println("Computer:\t\t", runtime.GOARCH)
}
