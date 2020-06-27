package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/seehuhn/mt19937"
)

func main() {
	rn := rand.New(mt19937.New())
	rn.Seed(time.Now().UnixNano())

	randomNumberSlice := []int{}
	//rnd2 := make([]int, 3, 5)
	//rnd2[0] = 44
	fmt.Printf("\t\t\t\t\t\t%p\n", randomNumberSlice)

	for j := 0; j < 5; j++ {
		for i := 0; i < 10; i++ {
			randomNumberSlice = append(randomNumberSlice, rn.Intn(1000))
		}
		sort.Ints(randomNumberSlice)
		fmt.Printf("%03d\t", randomNumberSlice)
		fmt.Printf("%p\t%p\n", randomNumberSlice, &randomNumberSlice[0])

		randomNumberSlice = nil
	}

}
