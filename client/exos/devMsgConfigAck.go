package main

import (
	fsm "ztp/client/fsm"
)

// evalute any events and tell FSM what to do
func (dev Device) ConfigAck() fsm.DeviceReturnCode {
	return fsm.DeviceReturnOK
}
