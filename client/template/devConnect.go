package main

import (
	msg "ztp"
	ztp "ztp/client"
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
func (dev *Device) Connect(connectMsg *msg.Connect) {
	connectMsg.ApPropertyBlock.RuSerialNumber = dev.devID
	connectMsg.ApPropertyBlock.BpWiredMacaddr = "00:04:96:9B:B7:E8"
	connectMsg.ApPropertyBlock.RuSwVersion = "30.2.2.17"
	connectMsg.ApPropertyBlock.RuModel = "X670G2-48x-4q"

	connectMsg.DeviceInfo.SysDescr = "ExtremeXOS (X670G2-48x-4q) version 30.2.2.17 30.2.2.17 by release-manager on Thu Jul 18 15:25:58 EDT 2019"
	connectMsg.DeviceInfo.SysUpTime = 128706300
	connectMsg.DeviceInfo.SysContact = "https://www.extremenetworks.com/support/"
	connectMsg.DeviceInfo.SysName = "X670G2-48x-4q"
	connectMsg.DeviceInfo.SysObjectID = "1.3.6.1.4.1.1916.2.195"
	connectMsg.DeviceInfo.OperatingSystem = "ExtremeXOS"
	connectMsg.DeviceInfo.SerialNumber = dev.devID
	connectMsg.DeviceInfo.MacAddr = connectMsg.ApPropertyBlock.BpWiredMacaddr
	connectMsg.DeviceInfo.MgmtIPAddr = "10.68.69.180"
	connectMsg.DeviceInfo.MgmtPort = "0"
	//connectMsg.DeviceInfo.License.FeaturePacks = make([]string, 0)
	//connectMsg.DeviceInfo.License.PortCapacityLicense = ""
	//connectMsg.DeviceInfo.License.EffectiveLicense = "Core"
	//connectMsg.DeviceInfo.License.EnabledLicense = "Core"
	//connectMsg.DeviceInfo.License.SystemLicense = ""
	//connectMsg.DeviceInfo.License.EffectiveLevel = ""
	//connectMsg.DeviceInfo.Ports
	//connectMsg.DeviceInfo.IfTable
	//connectMsg.DeviceInfo.PortsInfo

}

//The state machine executes this function during the CONNECT
//phase.  The function informs the application the state machine was
//successful in connecting to the controller.

//The state machine transitions to UPGRADE.
func (dev *Device) ConnectOk(in interface{}) {
}

//The state machine executes this function during the CONNECT
//phase.  The function informs the application that the state
//machine was unsuccessful in communicating with the discovered
//controller.
//
//If the application does not register for this function, the state
//machine transitions to DISCOVER if the HTTP status code is
//greater than or equal to 500.  For all other status codes that
//trigger a failed connection attempt, the state machine attempts
//to reconnect to the Controller.
//
//The state machine expects a return code, but it is not required.
//Unknown return codes result in default behavior.
//
//RESTART:
//Informs the state machine to rediscover the controller.  The
//state machine transitions to DISCOVER.
//
//RETRY:
//Informs the state machine to attempt to connect to the controller
//again after a configurable waiting period.  This value is
//controlled by self.data.args.retry_interval.  The state machine
//stays in CONNECT.
//
//FINISH:
//Informs the state machine to wrap things up by transitioning to
//DONE.
func (dev *Device) ConnectFail(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

//The state machine executes this function during the CONNECT
//phase.  The function informs the application that the state
//machine successfully connected to the Extreme Control service,
//but it wants the connector to communicate with a different
//controller.
//
//Unknown return codes are treated as OK.
//
//OK:
//Informs the state machine to accept the redirect and connect to
//the redirected controller.  The state machine transitions to
//DISCOVER to set the redirect address to the controller
//address value before transitioning to CONNECT.
//
//RESTART:
//Informs the state machine to ignore the redirect and attempt to
//rediscover the controller.  The state machine removes the
//reference to the redirect address and transitions to
//DISCOVER.
//
//FINISH:
//Informs the state machine to wrap things up by transitioning to
//DONE.
func (dev *Device) ConnectRedirect(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}
