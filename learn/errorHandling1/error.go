package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Enter your age:")
	var sAge string
	_, err := fmt.Scanf("%s", &sAge)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	nAge, err := strconv.ParseInt(sAge, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Printf("Your age is:%v -->%T", nAge, nAge)
	}
}
