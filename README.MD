## USBIP AUTO

### Server
This tool simply just automates the binding of 
USB devices via `usbipd -D`, 

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