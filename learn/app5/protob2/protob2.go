package main

import (
	"fmt"
	"learn/app5/protob2/pb2"
)

func main() {

	p1 := pb2.Animal{
		Id:         new(int32),
		AnimalType: new(string),
		Nickname:   new(string),
		Zone:       new(int32),
		Age:        new(int32),
	}

	*p1.Id = 1
	*p1.Age = 5
	*p1.AnimalType = "At"
	*p1.Nickname = "karatay"
	*p1.Zone = 3

	fmt.Println(p1.GetAge())

}
