package ztpclient

import ()

//This routine performs the 'INIT' state actions.
//
//Initializes the state machine and calls the application's INIT function
//callback.
//
//Adds the ztpclient to the upgrade asset list.
func (zc *ZtpClient) Init() (state ZtpClientState) {

	zc.Device.Init()
	return ZtpStateDiscover
}
