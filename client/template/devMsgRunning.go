package main

import (
	"log"
	"net/http"
	msg "ztp"
	fsm "ztp/client/fsm"
)

//The state machine calls this function to fill in the stats message
func (dev Device) Running(statsMsg *msg.Stats) {
	// populate the stats message
	statsMsg.ApPropertyBlock = *dev.property
	//statsMsg.Capabilities = *dev.capabilities
	//statsMsg.Assets = *dev.upgradeAssets
	//statsMsg.UpgradeAsset
	//statsMsg.ConfigBlock
	//statsMsg.Cooling
	//statsMsg.IfTable
	//statsMsg.PortsInfo
	//statsMsg.PowerSupply
	//statsMsg.Temperatue
	//statsMsg.Telemetry
	//statsMsg.IetfVxlanVxlan
	//statsMsg.IetfVxlanVxlanState
	//statsMsg.CPUUtilization
	//statsMsg.MemoryUtilization
	//statsMsg.Mlagv2
	//statsMsg.Timestamp
	//statsMsg.SysUpTime
	//statsMsg.CheckInTime
	//statsMsg.UpgradeTime

}

//The state machine issues this function while in RUNNING.
//
//The rate of the functions is configurable by the application.  This
//rate is stored in which has a unit
//type of seconds.
//
//The RUNNING state provides the application the ability to run
//asynchronously and provide the Extreme Control service feedback,
//e.g. statistics.
//
//The state machine expects a return code, but is not required.
//Unknown return codes are treated as OK.
//
//OK
//Informs the state machine stays in Running.
//
//RESTART
//Informs the state machine to transition to CONNECT.
//
//FINISH
//Informs the state machine to transition to DONE.
//
//A list of ConnectorEvents.
//The state machine will build an event message and send it to the
//Extreme Control service.
//
//A failure in the REST API attempt with the Extreme Control
//service causes the state machine to transition to CONNECT.
func (dev Device) RunningResponse(err error, resp *http.Response, statsResponseMsg *msg.StatsResponse) (ret fsm.DeviceReturnCode) {
	if err != nil {
		log.Println(err)
		return fsm.DeviceReturnRetry
	}
	if resp != nil {
		msg.DumpJson(statsResponseMsg)
	}
	switch resp.StatusCode {
	case http.StatusOK:
	default:
		if int(resp.StatusCode/100) == 4 {
			log.Println("400 ERRORS")
			// just keep trying for errors in the 400's
			return fsm.DeviceReturnOK
		}
		log.Println("SOME OTHER ERRORS")
		return fsm.DeviceReturnRestart
	}
	if statsResponseMsg.StatsInterval > 0 {
		dev.fsm.SetRunningRetry(statsResponseMsg.StatsInterval)
	}
	return fsm.DeviceReturnOK
}
