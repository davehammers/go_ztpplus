package main

import (
	ztp "ztp/client"
)

// evalute any events and tell FSM what to do
func (dev Device) ConfigAck() ztp.DeviceReturnCode {
	return ztp.DeviceReturnOK
}
