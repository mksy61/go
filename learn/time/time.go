package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Now()
	for i := 0; i < 1000000; i++ {

	}
	t1 := time.Now()

	t := t1.Sub(t0)

	fmt.Printf("%T\n%v", t, t)

}
