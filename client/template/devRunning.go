package main

import (
	ztp "ztp/client"
)

//The state machine issues this function while in RUNNING.
//
//The rate of the functions is configurable by the application.  This
//rate is stored in self.data.args.running_interval, which has a unit
//type of seconds.
//
//The RUNNING state provides the application the ability to run
//asynchronously and provide the Extreme Control service feedback,
//e.g. statistics.
//
//The state machine expects a return code, but is not required.
//Unknown return codes are treated as OK.
//
//OK
//Informs the state machine stays in Running.
//
//RESTART
//Informs the state machine to transition to CONNECT.
//
//FINISH
//Informs the state machine to transition to DONE.
//
//A list of ConnectorEvents.
//The state machine will build an event message and send it to the
//Extreme Control service.
//
//A failure in the REST API attempt with the Extreme Control
//service causes the state machine to transition to CONNECT.
func (dev *Device) Running(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOK
}
