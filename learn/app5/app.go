package main

import (
	"encoding/json"
	"learn/app5/portal"
	"log"
	"os"
)

func main() {
	configfile, err := os.Open("config.json")
	if err != nil {
		log.Panicln("Error:", err)
	}
	defer configfile.Close()

	var config = new(webserverconfig)
	err = json.NewDecoder(configfile).Decode(config)
	if err != nil {
		log.Panicln("Error:", err)
	}

	log.Println("Web server starting...")
	err = portal.RunWebPortal(config.WebServer)
	if err != nil {
		log.Panicln("Error:", err)
	}
}

type webserverconfig struct {
	WebServer string
}
