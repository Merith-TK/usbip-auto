package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

var (
	foo = "bar"
	str = `======================
	- 192.168.0.113
		 1-1.4: SanDisk Corp. : unknown product (0781:5575)
			  : /sys/devices/platform/soc/3f980000.usb/usb1/1-1/1-1.4
			  : (Defined at Interface level) (00/00/00)
   
		 1-1.3: SanDisk Corp. : unknown product (0781:5597)
			  : /sys/devices/platform/soc/3f980000.usb/usb1/1-1/1-1.3
			  : (Defined at Interface level) (00/00/00)
`

	remote string
)

func main() {
	flag.StringVar(&remote, "r", "", "Define the remote USBIP server's IP")
	flag.Parse()
	usbParse(str, remote)
}

func usbMount(b string, r string) {
	cmd := exec.Command("usbip", "attach", "-r", r, "-b", b)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Start()
}

func usbParse(str string, remote string) {
	if remote == "" {
		flag.Usage()
		os.Exit(1)
	}
	var re = regexp.MustCompile(`(1-1\.\d*)(:)`)
	var found []string
	for _, match := range re.FindAllStringSubmatch(str, -1) {
		found = append(found, match[1])
	}
	for i := 0; i < len(found); i++ {
		busid := found[i]
		usbMount(busid, remote)
		fmt.Println("BUSID:", busid)
	}
}
