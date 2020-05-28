package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"time"
)

func main() {
	cmd := exec.Command("usbipd", "-D")
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
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	slurp, _ := ioutil.ReadAll(stderr)
	if string(slurp) != "usbip: error: device on busid "+busid+" is already bound to usbip-host\n" {
		fmt.Println("Bound", busid, "Successfully")
	}
}

func usbParse(str string) {
	var re = regexp.MustCompile(`(?m)- busid (.+) \((.+)\)`)
	var found []string
	for _, match := range re.FindAllStringSubmatch(str, -1) {
		found = append(found, match[1])
	}
	for i := 0; i < len(found); i++ {
		busid := found[i]
		if !(busid == "1-1.1" || busid == "1-1.2") {
			//fmt.Println("BUSID:", busid)
			usbBind(busid)
		}
	}
}
