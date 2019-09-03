package main

import (
	ztp "ztp/client"
)

// The Device defines the information required by the FSM to manage a ZTP device instance
// Additional information may be added below the comment for a device specific implementtion.
type Device struct {

	// add device specific data elements here
}

//Create a new device instance that implements the FSM Device interface.
//
//The devID is the unique device identifier string provided to the controller.
//The devID must be unique for all devices reporting to a single controller. Typically this is a serial number or a device MAC address.
func NewDevice() *Device {
	dev := Device{}
	return &dev
}

//This function allows the application to override the default
//routine for retrieving the source IP address of the socket used
//to communicate with Extreme Control.
//
//The override routine SHALL take two parameters, an IP address and
//a port number.
func (dev *Device) GetSourceIP(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}
