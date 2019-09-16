package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
	msg "ztp"
	fsm "ztp/client/fsm"
)

// Environment variable fsmclientDEBUG=1 Sets the package DEBUG=true
const EnvDebug = "fsmDEBUG"

//DEBUG can be used during development to output log messages
var (
	DEBUG        = false
	deviceID     *string
	debugCtrlrIP *string
	pDebug       *bool
	pNumSim      *int
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
	pNumSim = flag.Int("sim", 0, "Number of additiona simulations of this device")
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
	// start the main instance
	dev := NewDevice(devID, false)
	wg.Add(1)
	go dev.StartFSM()
	// using go routines, any number of device simulations can be created
	for i := 0; i < *pNumSim; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			d := fmt.Sprintf("%s-sim%04d", devID, i+1)
			dev := NewDevice(d, true)
			dev.StartFSM()
		}(i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Wait()
}
