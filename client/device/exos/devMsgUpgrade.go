package main

import (
	"log"
	"net/http"
	msg "ztp"
	fsm "ztp/client/fsm"
)

//The state machine executes this function during the UPGRADE
//phase.  The function is issued prior to sending the upgradeimage
//REST API request to the Extreme Control service.
//
//If the application desires to upgrade software asSets under its
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
//  the application updates the asSet's version information reported
//  to the Extreme Control service and returns OK.
//
//OK
//  Informs the state machine to proceed with the upgradeimage
//  REST API request.
//
//RETRY
//  Informs the state machine to wait a period of time, configurable
//
func (dev Device) Upgrade(upgradeMsg *msg.ImageUpgrade) {
	upgradeMsg.ApPropertyBlock = *dev.property
	upgradeMsg.Assets = make([]msg.Asset, 0)

	asset := msg.Asset{
		AssetName:    "sim",
		AssetVersion: "1.1.1.1",
		AssetType:    "FIRMWARE",
	}
	upgradeMsg.Assets = append(upgradeMsg.Assets, asset)

	asset = msg.Asset{
		AssetName:    "cloud_connector",
		AssetVersion: "3.3.2.1",
		AssetType:    "XMOD",
	}
	upgradeMsg.Assets = append(upgradeMsg.Assets, asset)
	msg.DumpJson(upgradeMsg)
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
//to manipulate the wait time.
//
//FINISH:
//Informs the state machine to wrap things up by transitioning to
//DONE.
//
//A list of ConnectorEvents:
//The state machine sends these events to Extreme Control and
//transitions to CONFIG.
func (dev Device) UpgradeResponse(err error, resp *http.Response, upgradeResponse *msg.ImageUpgradeResponse) (ret fsm.DeviceReturnCode) {
	if err != nil {
		log.Println(err)
		return fsm.DeviceReturnRetry
	}
	if resp == nil {
		// didn't Get a proper HTTP response
		return fsm.DeviceReturnRetry
	}
	switch resp.StatusCode {
	case http.StatusOK:
	default:
		return fsm.DeviceReturnRetry
	}
	// platform updates asSets in the upgradeResponse
	msg.DumpJson(upgradeResponse)
	for _, image := range upgradeResponse.ImageUpgradeBlock {
		if image.Upgrade {
			log.Println("Upgrading")
			msg.DumpJson(image)
			// device TODO
		} else {
			log.Println("NOT Upgrading")
			msg.DumpJson(image)
		}
	}
	//  fsm.DeviceReturnOK - Upgrade is successful, continue
	//  fsm.DeviceReturnRetry - something went wrong, ask to upgrade again
	//  fsm.DeviceReturnFinish - Only upgrade, nothing more
	//return fsm.DeviceReturnFinish
	return fsm.DeviceReturnOK
}
