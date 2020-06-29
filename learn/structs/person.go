package main

import (
	"encoding/json"
	"fmt"
	"unsafe"
)

func main() {

	p1 := person{}

	p1 = person{
		TCKimlikNo: 33,
		Name:       "Veliyuddinil-hasan",
	}
	var p2 person = person{11, "Ali"}
	p3 := student{999, person{22, "Murat"}}

	fmt.Println(p1, p2, p3)

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(p3)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("JSON:", string(bs))

	bs, err = json.MarshalIndent(p3, " * ", " ")
	fmt.Println("JSON Ident:", string(bs))
	fmt.Println(unsafe.Sizeof(p1))
	fmt.Println(unsafe.Sizeof(p2))
	fmt.Println(unsafe.Sizeof(p3))
}

type person struct {
	TCKimlikNo int
	Name       string `json:"User Name"`
}

type student struct {
	StudentNo int
	person
}

func (p person) Len() int {
	return len(p.Name)
}
