package main

import (
	"log"
	"net/http"
	msg "ztp"
	fsm "ztp/client/fsm"
)

//The state machine executes this function before sending a request
//to the configuration REST API to the Extreme Control service.
//
//It is expected the application update its capabilities,
//configuration,
//The state machine includes this data in the configuration REST
//API request.
func (dev Device) Config(config *msg.Configuration) {
	// fill in the config message
	config.ApPropertyBlock = *dev.property
	for _, f := range *dev.features {
		// call all of the feature routines to fill in the connect message
		f.GetConfig(config)
		// save capabilities
		*dev.capabilities = config.Capabilities
	}
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
//the application POST_CONFIG functions have been processed.
func (dev Device) ConfigResponse(err error, resp *http.Response, configResponse *msg.ConfigurationResponse) (ret fsm.DeviceReturnCode) {
	if err != nil {
		log.Println(err)
		return fsm.DeviceReturnRestart
	}
	switch resp.StatusCode {
	case http.StatusOK:
		// OK HTTP response, keep going
	case http.StatusPreconditionFailed:
		/*
			if DEBUG {
				return fsm.DeviceReturnOK
			}
		*/
		return fsm.DeviceReturnRetry
	case http.StatusNotAcceptable:
		if DEBUG {
			return fsm.DeviceReturnRestart
		}
		return fsm.DeviceReturnRetry
	default:
		// got an unexpected HTTP code
		return fsm.DeviceReturnAbort
	}

	msg.DumpJson(configResponse)
	for _, f := range *dev.features {
		// call all of the feature routines to process the connect response
		f.SetConfig(configResponse)
	}

	return fsm.DeviceReturnOK
}
