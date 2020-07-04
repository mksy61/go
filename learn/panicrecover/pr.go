package main

import "fmt"

func main() {
	fmt.Println("1")
	panictest()
	fmt.Println("2")

}

func panictest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered:", err)
		}
	}()
	panic("panic")
}
