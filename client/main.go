package main

// Client software, the execute var will be
// delegated to $PATH in release, and remote
// will be delegated to the `-r` flag
import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"time"
)

var (
	execute = "C:/Tools/usbip/usbip.exe"
	//remote  = "192.168.0.113"
	remote string
)

func main() {
	flag.StringVar(&remote, "r", "", "Define the remote USBIP server's IP")
	flag.Parse()

	if remote == "" {
		flag.Usage()
		os.Exit(1)
	}
	for {
		usbParse(remote)
		time.Sleep(5 * time.Second)
	}
}

func usbMount(b string, r string) {
	cmd := exec.Command(execute, "attach", "-r", r, "-b", b)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Start()
}

func usbParse(remote string) {
	out, err := exec.Command(execute, "list", "-r", remote).Output()
	if err != nil {
		fmt.Println("CMD:", err)
		os.Exit(6)
	}
	str := string(out)
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
