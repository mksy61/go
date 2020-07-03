package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

func main() {
	listprocess()
}

func listprocess() {
	cmd := exec.Command("dir", "")
	resp, err := cmd.CombinedOutput()

	//fmt.Printf("%T", resp)
	if err != nil {
		log.Fatalln("Error:", err)
	}

	lines := strings.Split(string(resp), "\n")

	for _, v := range lines {
		fmt.Println(v)
	}

	time.Sleep(time.Second * 10)
}
