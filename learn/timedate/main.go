package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format("02/01/2006"))
	fmt.Println(t.Format("15:04:05"))
	fmt.Println(t.Format("15.04.05"))
	fmt.Println(t.Unix())
	fmt.Println(t.AddDate(-20, 0, 0))
	fmt.Println(t.Add(10 * time.Hour))

}
