package main

import (
	"log"
	"net/http"
	"strings"
	msg "ztp"
	fsm "ztp/client/fsm"
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
func (dev Device) Connect(connectMsg *msg.Connect) {
	for _, f := range *dev.features {
		// call all of the feature routines to fill in the connect message
		f.GetConnect(connectMsg)
	}

	// save the property block for the other message types
	*dev.property = connectMsg.ApPropertyBlock

	// use initial constant for auth
	userName := des3Encrypt("ezconfiguser")
	dev.fsm.SetLogin(userName)
	dev.fsm.SetPassword(userName)
	msg.DumpJson(connectMsg)

}

//The state machine executes this function during the CONNECT
//phase.  The function informs the application the state machine was
//successful in connecting to the controller.

//The state machine transitions to UPGRADE.
func (dev Device) ConnectResponse(err error, resp *http.Response, m *msg.ConnectResponse) (ret fsm.DeviceReturnCode) {
	if err != nil {
		if DEBUG {
			log.Println(err)
		}
		return fsm.DeviceReturnRestart
	}
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated:
	default:
		if int(resp.StatusCode/100) == 4 { //http 400's
			return fsm.DeviceReturnRetry
		}
		log.Println("response code", resp.StatusCode)
		return fsm.DeviceReturnRestart
	}
	msg.DumpJson(m)
	// if new auth is provided, decode the login/password provided
	if m.Credentials.Login != "" && m.Credentials.Password != "" {
		login := des3Decrypt(m.Credentials.Login)
		password := des3Decrypt(m.Credentials.Password)
		// trim 3 characters from either side of the password
		login = des3Encrypt(login[3 : len(login)-3])
		password = des3Encrypt(password[3 : len(password)-3])
		// update the FSM with the new credentials
		dev.fsm.SetLogin(login)
		dev.fsm.SetPassword(password)
	}

	m.Credentials.Login = "hidden"
	m.Credentials.Password = "hidden"
	msg.DumpJson(m)

	// is there a redirected to a different controller
	// ignore case when comparing strings
	if strings.EqualFold(m.Action, "redirect") {
		if strings.EqualFold(m.Redirect.Type, "https") {
			if m.Redirect.URI != "" {
				dev.fsm.SetRedirect(m.Redirect.URI)
				return fsm.DeviceReturnRetry
			}
		}
	}
	return fsm.DeviceReturnOK
}
