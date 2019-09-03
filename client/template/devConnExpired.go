package main

import (
	ztp "ztp/client"
)

//The state machine executes this function when the transmission of
//the ConfigAck message fails.  The application SHALL restore the
//configuration to its prior state.  If needed, the application MAY
//reboot the device.  If the application reboots the device, it
//SHALL store the reason for the reboot and inform the state machine
//by adding ConnectorEvent(s) to the self.data.args.reboot_events
//upon connector startup after the reboot.
//
//The state machine transmits the reboot events after the 'CONNECT'
//state has completed, but prior to issuing the 'CONNECT_OK'
//function.
//
//A return code is not expected, and the state machine transitions to
//CONNECT.
func (dev *Device) ConnectionExpired(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}
