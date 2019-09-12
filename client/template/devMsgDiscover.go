package main

import (
	"log"
	"net"
	"net/http"
	msg "ztp"
	fsm "ztp/client/fsm"
)

//The state machine executes this function during the DISCOVER
//phase.  The function is issued prior to auto discovering the controller.
//The Discover() function returns the list of controller domain names/ip address
//for the state machine to look for.
//
//
func (dev Device) Discover() (controllerList *[]fsm.ZtpLookupEntry) {
	//DNS controller names
	var err error
	controller := fsm.ZtpLookupEntry{}
	c := make([]fsm.ZtpLookupEntry, 0)
	controllerList = &c
	// if a command line override was provided, look for that entry first
	if *debugCtrlrIP != "" {
		controller.Host, controller.Port, err = net.SplitHostPort(*debugCtrlrIP)
		if err == nil {
			if controller.Port != "" {
				controller.Port = ":" + controller.Port
			}
			// prepend the /Client URI
			controller.AddClientPrefix = true
			// insert the debug controller value at the front of the list
			*controllerList = append(*controllerList, controller)
			controller.AddClientPrefix = false
			*controllerList = append(*controllerList, controller)
		} else {
			log.Println(err)
		}
	}
	*controllerList = append(*controllerList, fsm.ZtpLookupEntry{"extremecontrol", ":8443", true})
	*controllerList = append(*controllerList, fsm.ZtpLookupEntry{"extremecontrol.extremenetworks.com", ":8443", true})
	msg.DumpJson(controllerList)
	return
}

//The state machine calls the DiscoverResponse function with the HTTP response
//and ZtpLookupEntry of controller entry found after each pass through the controller list
//provided. If no controller is found, discoverMsg is nil.
//
//OK: The state machine will continue to look for
//the domain names/ip address provided until at least one controller is found.
//The state maching continues to the the CONNECT state
//
//FINISH:
//Informs the state machine to wrap things up by transitioning to DONE.
func (dev Device) DiscoverResponse(err error, resp *http.Response, discoverResponseMsg *fsm.ZtpLookupEntry) (ret fsm.DeviceReturnCode) {
	if err != nil {
		// usually an HTTP connectivity issue
		log.Println(err)
		return fsm.DeviceReturnRetry
	}
	if resp == nil {
		// no host found
		log.Println("No HOST found")
		return fsm.DeviceReturnRestart
	}
	msg.DumpJson(discoverResponseMsg)
	switch resp.StatusCode {
	case http.StatusOK: // not sure how this happens since we are just determinine connectivity
	case http.StatusNotFound:
	default:
		return fsm.DeviceReturnRetry // found something on HTTP but it doesn't like us
	}
	dev.controller = discoverResponseMsg
	return fsm.DeviceReturnOK
}
