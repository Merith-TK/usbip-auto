## USBIP AUTO

### Server
This tool simply just automates the binding of 
USB devices via `usbipd -D`, 

this should be run as the same permission level
needed to run `usbpid` normally

While this tool will work, you will have to set
your whitelist on what ports to not bind publicly
in `server/main.go` 
```go
	for i := 0; i < len(found); i++ {
		busid := found[i]
		if !(busid == "1-1.1" || busid == "1-1.2") {
			//fmt.Println("BUSID:", busid)
			usbBind(busid)
		}
    }
```

just extend this whith `|| busid == "your whitelisted busid"`
in the parenthesis

### Client
first off, you need to install the client software
and run at the same permission level as needed
to run `usbip attach -r 192.168.0.113 -b 1-1.2`
normally

> Windows users will need to put this program in
> the same folder where usbip.exe is stored, or
> recompile the program with hardpaths to the 
> exe files