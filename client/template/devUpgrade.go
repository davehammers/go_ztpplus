package main

import (
	msg "ztp"
	ztp "ztp/client"
)

//The state machine executes this function during the UPGRADE
//phase.  The function is issued prior to sending the upgradeimage
//REST API request to the Extreme Control service.
//
//If the application desires to upgrade software assets under its
//control, it MUST populate the UpgradeAssets list.
//Entries in this list MUST be an AssetEntry type.
//
//The application may use this function in combination with the
//UPGRADE function to perform software updates asynchronously without
//terminating the state machine.
//
//Use case:
//  The application processes an upgradeimage response, and is
//  required to upgrade the software of another application.  The
//  application would return RETRY for the UPGRADE
//  function, keeping the state machine in the UPGRADE phase.
//  The application would use this function to maintain that state
//  until the upgrade of the application is complete.  Once complete,
//  the application updates the asset's version information reported
//  to the Extreme Control service and returns OK.
//
//OK
//  Informs the state machine to proceed with the upgradeimage
//  REST API request.
//
//RETRY
//  Informs the state machine to wait a period of time, configurable
//  via self.data.args.retry_interval.
//
func (dev *Device) UpgradeCheck(in interface{}) ztp.DeviceReturnCode {
	return ztp.DeviceReturnOk
}

//The state machine executes this function after receiving the
//response to the upgradeimage REST API request from the Extreme
//Control service, which occurs in the UPGRADE phase.
//
//The state machines passes the JSON data response data
//The application is responsible for taking any actions on this data.
//
//OK:
//Informs the state machine to transition to the next state,
//CONFIG.
//
//RETRY:
//Causes the state machine to wait for a period of time before
//retrying.  The application may modify
//self.data.args.retry_interval to manipulate the wait time.
//
//FINISH:
//Informs the state machine to wrap things up by transitioning to
//DONE.
//
//A list of ConnectorEvents:
//The state machine sends these events to Extreme Control and
//transitions to CONFIG.
func (dev *Device) Upgrade(in interface{}) (ret ztp.DeviceReturnCode, events *[]msg.Event) {
	return ztp.DeviceReturnOk, nil
}
