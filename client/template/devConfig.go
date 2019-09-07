package main

import (
	"log"
	"net/http"
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
func (dev Device) Config(config *msg.Configuration) {
	// fill in the config message
	config.ApPropertyBlock = *dev.data.property
	config.Capabilities = *dev.data.capabilities
	// config.ConfigBlock = the platform fills in the config block
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
func (dev Device) ConfigResponse(err error, resp *http.Response, configResponse *msg.ConfigurationResponse) (ret ztp.DeviceReturnCode) {
	if err != nil {
		log.Println(err)
		return ztp.DeviceReturnRestart
	}
	switch resp.StatusCode {
	case http.StatusOK:
		// OK HTTP response, keep going
	case http.StatusPreconditionFailed:
		/*
			if DEBUG {
				return ztp.DeviceReturnOK
			}
		*/
		return ztp.DeviceReturnRetry
	case http.StatusNotAcceptable:
		if DEBUG {
			return ztp.DeviceReturnRestart
		}
		return ztp.DeviceReturnRetry
	default:
		// got an unexpected HTTP code
		return ztp.DeviceReturnAbort
	}
	msg.DumpJson(configResponse)

	//configResponse.License
	//configResponse.Poe
	//configResponse.Dot1X
	//configResponse.Lacp
	//configResponse.Lldp
	//configResponse.Logins
	//configResponse.MacAuth
	//configResponse.Ports
	//configResponse.Snmp
	//configResponse.SpanningTree
	//configResponse.Syslog
	//configResponse.Vlans
	//configResponse.MgmtAccess
	//configResponse.Appliance
	//configResponse.MLag
	//configResponse.MLagv2
	//configResponse.Stack
	//configResponse.Mvrp
	//configResponse.Ospf
	//configResponse.Timestamp
	//configResponse.BpRequestID
	//configResponse.Status
	return ztp.DeviceReturnOK
}
