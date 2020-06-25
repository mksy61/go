// Murat Aksoy
// Learning GO
// 2020
package main

import (
	"fmt"
	"os"
)

func main() {

	for i, v := range os.Args {
		fmt.Printf("%#v :%v\n", i, v)
	}

	var name string

	fmt.Printf("Birşeyler Yazın:")
	fmt.Scanf("%s", &name)
	fmt.Println("-->", name)
}
