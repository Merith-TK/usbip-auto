package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	defaultOut, err := exec.Command("usbip", "list", "-l").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(defaultOut), "\n ------------------------")

	for true {
		newOut, err := exec.Command("usbip", "list", "-l").Output()
		if err != nil {
			log.Fatal(err)
		}
		if string(defaultOut) != string(newOut) {
			defaultOut = newOut
			fmt.Println(string(newOut), "\n ------------------------")
		}
		time.Sleep(5 * time.Second)
	}
}
