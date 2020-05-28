package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"
)

func main() {
	cmd := exec.Command("usbipd", "-D")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	for true {
		out, err := exec.Command("usbip", "list", "-l").Output()
		if err != nil {
			log.Fatal(err)
		}
		usbParse(string(out))
		time.Sleep(10 * time.Second)
	}
}

func usbBind(busid string) {
	cmd := exec.Command("usbip", "bind", "-b", busid)
	cmd.Start()
}

func usbParse(str string) {
	var re = regexp.MustCompile(`(?m)- busid (.+) \((.+)\)`)
	var found []string
	for _, match := range re.FindAllStringSubmatch(str, -1) {
		found = append(found, match[1])
	}
	for i := 0; i < len(found); i++ {
		busid := found[i]
		if busid == "1-1.1" || busid == "1-1.2" {

		} else {
			fmt.Println("BUSID:", busid)
			usbBind(busid)
		}
	}
}
