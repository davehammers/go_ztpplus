package main

import (
	msg "ztp"
	ztp "ztp/client"
)

//The state machine issues this function while in SendingEvents.
//The device is expected to fill in the EventMsg with events.
//if the event list is empty, the state machine does not send the message
//
//OK
//Informs the state machine to send the event message
//
func (dev Device) SendEvents(eventMsg *msg.EventMsg) (ret ztp.DeviceReturnCode) {
	eventMsg.ApPropertyBlock = *dev.data.property
	eventMsg.Event = dev.data.events
	//eventMsg.ConfigACK
	return ztp.DeviceReturnOK
}

//The state machine issues this function when events have been sent to the controller
//The device resets the event list to an empty list
func (dev Device) SendEventsComplete() {
	dev.data.events = make([]msg.Event, 0)
}
