package main

import (
	"fmt"
)

func main() {
	p1 := person{"Ali", 48}
	p2 := person{"Veli", 55}

	fmt.Println(p1)
	//fmt.Println(strings.Repeat("-", 40))
	fmt.Println(p2)

}

type person struct {
	name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%16s\t%02d", p.name, p.age)
}

func (p person) Error() error {
	return fmt.Errorf("error detected")
}
