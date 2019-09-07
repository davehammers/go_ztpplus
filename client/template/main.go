package main

import (
	"flag"
	"os"
	"sync"
	msg "ztp"
	ztp "ztp/client"
)

// Environment variable ztpclientDEBUG=1 sets the package DEBUG=true
const EnvDebug = "ztpDEBUG"

//DEBUG can be used during development to output log messages
var (
	DEBUG        = false
	deviceID     *string
	debugCtrlrIP *string
	pDebug       *bool
)

func init() {
	envSetup()
	flagSetup()
}
func envSetup() {
	if d, ok := os.LookupEnv(EnvDebug); ok {
		switch d {
		case "0":
			DEBUG = false
		case "1":
			DEBUG = true
		}
	}
}
func flagSetup() {
	pDebug = flag.Bool("debug", false, "enable debug")
	deviceID = flag.String("d", "", "Device ID to use for this device")
	debugCtrlrIP = flag.String("c", "", "Controller FQDN:port or IP:port")
	flag.Parse()
	DEBUG = *pDebug
	if DEBUG {
		ztp.DEBUG = *pDebug
		msg.DEBUG = *pDebug
	}
}

func main() {
	var wg sync.WaitGroup
	// Get device ID from the device.
	// Here we just code a string for the template
	devID := "12345-67890-sim1"
	// override with any command line debug
	if *deviceID != "" {
		devID = *deviceID
	}
	// using go routines, any number of device simulations can be created
	wg.Add(1)
	go func(devID string) {
		defer wg.Done()

		dev := NewDevice(devID)
		dev.StartFSM()
	}(devID)
	wg.Wait()
}
