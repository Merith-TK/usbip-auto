package main

// Server Software, currently the black listed
// usb ports have to be hardcoded, this needs
// to be ran as root or with a user that can
// run `usbipd -D` and `usbip bind -b busid`
import (
	"fmt"
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
		time.Sleep(2 * time.Second)
	}
}

func usbBind(busid string) {
	cmd := exec.Command("usbip", "bind", "-b", busid)
	cmd.Start()
	fmt.Println("BOUND: ", busid)
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
			usbBind(busid)
		}
	}
}
