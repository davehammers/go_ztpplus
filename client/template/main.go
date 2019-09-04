package main

import (
	"flag"
	"log"
	"os"
	ztp "ztp/client"
)

// Environment variable ztpclientDEBUG=1 sets the package DEBUG=true
const EnvDebug = "ztpDEBUG"

//DEBUG can be used during development to output log messages
var (
	DEBUG    = false
	deviceID *string
	ctrlrIP  *string
	pDebug   *bool
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
	ctrlrIP = flag.String("c", "", "Controller FQDN:port or IP:port")
	flag.Parse()
	DEBUG = *pDebug
	ztp.DEBUG = *pDebug
}

func main() {
	// create a new device instance
	dev := NewDevice()
	// Get device ID from the device.
	// Here we just code a string for the template
	dev.devID = "12345-67890"

	if *deviceID != "" {
		dev.devID = *deviceID
	}

	// new FSM for our device instance
	fsm := ztp.NewZtpClient(dev)
	fsm.SetDeviceID(dev.devID)
	fsm.SetDeviceType("switch")
	//fsm.SetRebootEvent(rebootEvent)
	fsm.SetController(*ctrlrIP)
	//fsm.SetDiscoverRetry(retry)
	//fsm.SetPropertyBlock(property *msg.ApPropertyBlock)
	//fsm.AddUpgradeAsset(deviceAsset *msg.Assets)
	if DEBUG {
		log.Printf("%+v\n", fsm)
	}

	// start FSM
	fsm.StateMachine()
}
