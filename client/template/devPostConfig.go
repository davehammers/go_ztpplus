package main

import (
	ztp "ztp/client"
)

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
func (dev Device) PostConfig() (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOK
}
