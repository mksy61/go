//log paketi dah verimli

package main

import (
	
	"log")


func main() {
	//log.SetFlags(0)
	//log.SetOutput(ioutil.Discard)
	log.Println("123abc")
	log.Panicln("456def")

}