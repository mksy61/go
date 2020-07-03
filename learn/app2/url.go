package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "http://golang.org"

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln("Error", err)
		return
	}
	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Error", err)
		return
	}

	fmt.Println(string(html))
}
