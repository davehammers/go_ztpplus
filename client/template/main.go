package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	msg "ztp"
	fsm "ztp/client/fsm"
)

// Environment variable fsmclientDEBUG=1 sets the package DEBUG=true
const EnvDebug = "fsmDEBUG"

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
		fsm.DEBUG = *pDebug
		msg.DEBUG = *pDebug
	}
}

func main() {
	var wg sync.WaitGroup
	// Get device ID from the device.
	// Here we just code a string for the template
	// override with any command line debug
	devID := "12345-67890"
	if *deviceID != "" {
		devID = *deviceID
	}
	// using go routines, any number of device simulations can be created
	for i := 1; i < 2; i++ {
		d := fmt.Sprintf("%s-sim%d", devID, i)
		wg.Add(1)
		go func(devID string) {
			defer wg.Done()

			dev := NewDevice(devID)
			dev.StartFSM()
		}(d)
	}
	wg.Wait()
}
