package main

import (
	msg "ztp"
	ztp "ztp/client"
)

//The state machine executes this function before sending a request
//to the configuration REST API to the Extreme Control service.
//
//It is expected the application update its capabilities,
//configuration, and portsInfo by modifying
//self.data.ccsm.capabilities, self.data.ccsm.config_data,
//and self.data.ccsm.ports_info_data respectively, if necessary.
//The state machine includes this data in the configuration REST
//API request.
func (dev *Device) ConfigUpdate(in interface{}) {
}

//The state machine executes this function after receiving the
//response to the configuration REST API request from the Extreme
//Control service.
//
//The state machines passes the config message
//The application is responsible for/taking any actions on this data.
//
//A return code is expected in the form of a list.  The list may be
//empty or it SHALL contain EventInstance objects.  The state
//machine scans this list for 'critical' or 'major' event errors.
//If any of these error events exist, the state machine transitions
//to CONFIG_ACK.
//
//If there are no errors, the state machine transitions to
//POST_CONFIG.  The state machine postpones the ConfigAck until
//the application POST_CONFIG functions have been processed.  The
//ConfigAck events are stored in self.data.ccsm.config_ack.
func (dev *Device) Config(in interface{}) (events *[]msg.Event) {
	return nil
}

//The state machine executes these functions during the
//POST_CONFIG phase.  The functions are issued in the order
//they were added to the function list.  The POST_CONFIG phase allows
//the application to perform additional tasks that are needed to
//onboard the application with the Extreme Control service.
//
//OK:
//Informs the state machine this task is complete and to move onto
//the next function.  If this is the last function, the state
//machine transitions to CONFIG_ACK.
//
//RETRY:
//Informs the state machine to retry the current function.  The
//retry attempts occurs after a configurable period of time.
//The application may modify self.data.args.retry_interval to
//control this time.
//
//RESTART:
//Informs the state machine to restart the POST_CONFIG functions.
//The state machine restarts these functions after a configurable
//period of time.  The application controls this time by modifying
//self.data.args.retry_interval.
//
//FINISH:
//Informs the state machine to abandon the rest of the functions
//and transition to CONFIG_ACK.
//
//ABORT:
//The function encountered a failure that requires the state
//machine to revert back to CONNECT.  The state machine
//waits for the self.data.args.error_interval to expire before
//reconnecting to the Extreme Control service.
func (dev *Device) PostConfig(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

//The state machine executes this function during the
//CONFIG_ACK phase.  The function is made if the event
//list returned by the CONFIG function DOES NOT contain any critical
//or major error events, the POST_CONFIG functions have been
//processed, and the Extreme Control service responds to the
//ConfigAck message with an HTTP 200 OK.
//
//An application MAY save the configuration at this time.
//
//The state machine transitions to
//RUNNING.  A connection failure when sending the ConfigAck
//message to the Extreme Control service causes the state machine to
//transition to CONNECT.
func (dev *Device) ConfigOK(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

//The state machine executes this function during the
//CONFIG_ACK phase.  The function is made if the event list
//provided by the CONFIG function CONTAINS critical or major error
//events, and the Extreme Control service responds to the ConfigAck
//message with an HTTP 200 OK.
//
//The configuration of the device MUST be an atomic operation.  A
//failure requires the application to restore the last known good
//configuration.
//
//A return code is not expected, and the state machine remains in
//CONFIG.  A connection failure when sending the ConfigAck
//message to the Extreme Control service causes the state machine to
//transition to CONNECT.
func (dev *Device) ConfigFail(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}
