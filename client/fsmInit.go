package ztpclient

import (
	"log"
)

//This routine performs the 'INIT' state actions.
//
//Initializes the state machine and calls the application's INIT function
//callback.
//
//Adds the ztpclient to the upgrade asset list.
func (zc *ZtpClient) Init() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

	zc.device.Init()
	return ZtpStateDiscover
}
