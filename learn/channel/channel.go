package main

import "fmt"

func main() {
	ch := make(chan bool)
	go f(ch)

	fmt.Println("main 1")
	ok := <-ch
	if ok {
		fmt.Println("main 2")
	}

	close(ch)
	fmt.Println("main 3")
}

func f(ch chan bool) {
	fmt.Println("f executed")
	ch <- true
}
