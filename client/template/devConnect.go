package main

import (
	msg "ztp"
	ztp "ztp/client"
)

//The state machine executes this function during the CONNECT
//phase.  The function is issued prior to attempting to connect to
//the Extreme Control service.
//
//The application is expected to populate the platform name, property
//block, and device information.  Failure to populate these fields
//will result in the failure of the connection attempt.
//
//The state machine issues a
//CONNECT_OK or CONNECT_FAIL function as a result of the connection
//attempt.
func (dev *Device) Connect(connectMsg *msg.Connect) {
}

//The state machine executes this function during the CONNECT
//phase.  The function informs the application the state machine was
//successful in connecting to the controller.

//The state machine transitions to UPGRADE.
func (dev *Device) ConnectOk(in interface{}) {
}

//The state machine executes this function during the CONNECT
//phase.  The function informs the application that the state
//machine was unsuccessful in communicating with the discovered
//controller.
//
//If the application does not register for this function, the state
//machine transitions to DISCOVER if the HTTP status code is
//greater than or equal to 500.  For all other status codes that
//trigger a failed connection attempt, the state machine attempts
//to reconnect to the Controller.
//
//The state machine expects a return code, but it is not required.
//Unknown return codes result in default behavior.
//
//RESTART:
//Informs the state machine to rediscover the controller.  The
//state machine transitions to DISCOVER.
//
//RETRY:
//Informs the state machine to attempt to connect to the controller
//again after a configurable waiting period.  This value is
//controlled by self.data.args.retry_interval.  The state machine
//stays in CONNECT.
//
//FINISH:
//Informs the state machine to wrap things up by transitioning to
//DONE.
func (dev *Device) ConnectFail(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

//The state machine executes this function during the CONNECT
//phase.  The function informs the application that the state
//machine successfully connected to the Extreme Control service,
//but it wants the connector to communicate with a different
//controller.
//
//Unknown return codes are treated as OK.
//
//OK:
//Informs the state machine to accept the redirect and connect to
//the redirected controller.  The state machine transitions to
//DISCOVER to set the redirect address to the controller
//address value before transitioning to CONNECT.
//
//RESTART:
//Informs the state machine to ignore the redirect and attempt to
//rediscover the controller.  The state machine removes the
//reference to the redirect address and transitions to
//DISCOVER.
//
//FINISH:
//Informs the state machine to wrap things up by transitioning to
//DONE.
func (dev *Device) ConnectRedirect(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}
