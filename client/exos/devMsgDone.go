package main

import (
//fsm "ztp/client/fsm"
)

//The state machine executes this function prior to completion.
//
//The state machine does not perform any actions by default.  The
//application may set self.data.args.delete to True to cause the
//state machine to send a delete REST API request in order to
//remove the application's object from the Extreme Control service.
//
//A return code is not expected, and the execution of the state
//machine completes after the application returns from the function.
func (dev Device) Done() {
}
