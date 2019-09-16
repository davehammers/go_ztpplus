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
	*dev.property = msg.ApPropertyBlock{
		RuSerialNumber: dev.devID,
		BpWiredMacaddr: "00:04:96:9B:B7:E8",
		RuSwVersion:    "1.2.3.4",
		RuModel:        "sim-template",
	}
	connectMsg.ApPropertyBlock = *dev.property

	connectMsg.DeviceInfo.SysDescr = "template simulation for fsm"
	connectMsg.DeviceInfo.SysUpTime = 128706300
	connectMsg.DeviceInfo.SysContact = "https://www.extremenetworks.com/support/"
	connectMsg.DeviceInfo.SysName = connectMsg.ApPropertyBlock.RuModel
	connectMsg.DeviceInfo.SysObjectID = "1.3.6.1.4.1.1916.2.195"
	connectMsg.DeviceInfo.OperatingSystem = "ExtremeXOS"
	connectMsg.DeviceInfo.SerialNumber = dev.devID
	connectMsg.DeviceInfo.MacAddr = connectMsg.ApPropertyBlock.BpWiredMacaddr
	connectMsg.DeviceInfo.MgmtIPAddr = "10.10.10.1"
	connectMsg.DeviceInfo.MgmtPort = "0"
	for _, f := range *dev.features {
		// call all of the feature routines to fill in the connect message
		f.GetConnect(connectMsg)
	}

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
func (dev Device) ConnectResponse(err error, resp *http.Response, connectResponse *msg.ConnectResponse) (ret fsm.DeviceReturnCode) {
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
	msg.DumpJson(connectResponse)
	// if new auth is provided, decode the login/password provided
	if connectResponse.Credentials.Login != "" && connectResponse.Credentials.Password != "" {
		login := des3Decrypt(connectResponse.Credentials.Login)
		password := des3Decrypt(connectResponse.Credentials.Password)
		// trim 3 characters from either side of the password
		login = des3Encrypt(login[3 : len(login)-3])
		password = des3Encrypt(password[3 : len(password)-3])
		// update the FSM with the new credentials
		dev.fsm.SetLogin(login)
		dev.fsm.SetPassword(password)
	}

	connectResponse.Credentials.Login = "hidden"
	connectResponse.Credentials.Password = "hidden"
	msg.DumpJson(connectResponse)

	// is there a redirected to a different controller
	// ignore case when comparing strings
	if strings.EqualFold(connectResponse.Action, "redirect") {
		if strings.EqualFold(connectResponse.Redirect.Type, "https") {
			if connectResponse.Redirect.URI != "" {
				dev.fsm.SetRedirect(connectResponse.Redirect.URI)
				return fsm.DeviceReturnRetry
			}
		}
	}
	return fsm.DeviceReturnOK
}
