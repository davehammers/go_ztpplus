package main

import (
	ztp "ztp/client"
)

//The state machine executes this function during the DISCOVER
//phase.  The function is issued prior to auto discovering the
//Extreme Control service. The auto discovery mechanism attempts to
//locate the Extreme Control service on the local domain first, and,
//if NOT FOUND, attempts to locate an Extreme Control service in the
//cloud.
//
//The application MAY provide a fallback controller by setting
//self.data.args.controller_addr.  If the auto discovery mechanism
//fails to find the Extreme Control service, the state machine will
//use the static controller defined by the application.
//
//The state machine transitions
//CONNECT after discovering the Extreme Control service. If
//there is an error in the discovery process, the state machine stays
//in DISCOVER and continuously attempts to locate the Extreme
//Control service.  The application function is executed on each
//iteration.
func (dev *Device) Discover(in interface{}) {
}

func (dev *Device) DiscoverOK(*ztp.ZtpLookupEntry) {
}

//The state machine executes this function during the DISCOVER
//phase.  The function is issued after the state machine has failed
//to discover the Extreme Control service.
//
//The state machine expects a return code, but it is not required.
//Unknown return codes are treated as OK.
//
//Return
//
//OK:
//Informs the state machine to continue its discovery operation.
//
//FINISH:
//Informs the state machine to wrap things up by transitioning to DONE.
func (dev *Device) DiscoverFail() (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}
