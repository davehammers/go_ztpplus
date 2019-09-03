package main

import (
	msg "ztp"
	ztp "ztp/client"
)

// The Device defines the information required by the FSM to manage a ZTP device instance
// Additional information may be added below the comment for a device specific implementtion.
type Device struct {

	// add device specific data elements here
}

//Create a new device instance that implements the FSM Device interface.
//
//The devID is the unique device identifier string provided to the controller.
//The devID must be unique for all devices reporting to a single controller. Typically this is a serial number or a device MAC address.
func NewDevice() *Device {
	dev := Device{}
	return &dev
}

//The state machine executes this function during the INIT
//phase.
//
//The application performs any initialization it may require to
//support the execution of the state machine.
//
//The state machine adds the ztp client to the upgrade asset
//list.
//
//The application SHOULD populate the TBD
//the reboot events if the connector rebooted the device.
//
//A return code of OK, and the state machine transitions
//to DISCOVER.
func (dev *Device) Init() {
}

//The state machine executes this function during the DISCOVER
//phase.  The function is issued prior to auto discovering the
//Extreme Control service. The auto discovery mechanism attempts to
//locate the Extreme Control service on the local domain first, and,
//if NOT FOUND, attempts to locate an Extreme Control service in the
//cloud.
//
//The application MAY provide a fallback controller by setting
//self.data.args.controller_addr.  If the auto discovery mechanism
//fails to find the Extreme Control service, the state machine will
//use the static controller defined by the application.
//
//The state machine transitions
//CONNECT after discovering the Extreme Control service. If
//there is an error in the discovery process, the state machine stays
//in DISCOVER and continuously attempts to locate the Extreme
//Control service.  The application function is executed on each
//iteration.
func (dev *Device) Discover(in interface{}) {
}

func (dev *Device) DiscoverOK(*ztp.ZtpLookupEntry) {
}

//The state machine executes this function during the DISCOVER
//phase.  The function is issued after the state machine has failed
//to discover the Extreme Control service.
//
//The state machine expects a return code, but it is not required.
//Unknown return codes are treated as OK.
//
//Return
//
//OK:
//Informs the state machine to continue its discovery operation.
//
//FINISH:
//Informs the state machine to wrap things up by transitioning to DONE.
func (dev *Device) DiscoverFail() (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

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

//The state machine executes this function before sending a request
//to the configuration REST API to the Extreme Control service.
//
//It is expected the application update its capabilities,
//configuration, and portsInfo by modifying
//self.data.ccsm.capabilities, self.data.ccsm.config_data,
//and self.data.ccsm.ports_info_data respectively, if necessary.
//The state machine includes this data in the configuration REST
//API request.
func (dev *Device) ConfigUpdate(in interface{}) {
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
//the application POST_CONFIG functions have been processed.  The
//ConfigAck events are stored in self.data.ccsm.config_ack.
func (dev *Device) Config(in interface{}) (events *[]msg.Event) {
	return nil
}

//The state machine executes these functions during the
//POST_CONFIG phase.  The functions are issued in the order
//they were added to the function list.  The POST_CONFIG phase allows
//the application to perform additional tasks that are needed to
//onboard the application with the Extreme Control service.
//
//OK:
//Informs the state machine this task is complete and to move onto
//the next function.  If this is the last function, the state
//machine transitions to CONFIG_ACK.
//
//RETRY:
//Informs the state machine to retry the current function.  The
//retry attempts occurs after a configurable period of time.
//The application may modify self.data.args.retry_interval to
//control this time.
//
//RESTART:
//Informs the state machine to restart the POST_CONFIG functions.
//The state machine restarts these functions after a configurable
//period of time.  The application controls this time by modifying
//self.data.args.retry_interval.
//
//FINISH:
//Informs the state machine to abandon the rest of the functions
//and transition to CONFIG_ACK.
//
//ABORT:
//The function encountered a failure that requires the state
//machine to revert back to CONNECT.  The state machine
//waits for the self.data.args.error_interval to expire before
//reconnecting to the Extreme Control service.
func (dev *Device) PostConfig(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

//The state machine executes this function during the
//CONFIG_ACK phase.  The function is made if the event
//list returned by the CONFIG function DOES NOT contain any critical
//or major error events, the POST_CONFIG functions have been
//processed, and the Extreme Control service responds to the
//ConfigAck message with an HTTP 200 OK.
//
//An application MAY save the configuration at this time.
//
//The state machine transitions to
//RUNNING.  A connection failure when sending the ConfigAck
//message to the Extreme Control service causes the state machine to
//transition to CONNECT.
func (dev *Device) ConfigOK(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

//The state machine executes this function during the
//CONFIG_ACK phase.  The function is made if the event list
//provided by the CONFIG function CONTAINS critical or major error
//events, and the Extreme Control service responds to the ConfigAck
//message with an HTTP 200 OK.
//
//The configuration of the device MUST be an atomic operation.  A
//failure requires the application to restore the last known good
//configuration.
//
//A return code is not expected, and the state machine remains in
//CONFIG.  A connection failure when sending the ConfigAck
//message to the Extreme Control service causes the state machine to
//transition to CONNECT.
func (dev *Device) ConfigFail(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

//The state machine issues this function while in RUNNING.
//
//The rate of the functions is configurable by the application.  This
//rate is stored in self.data.args.running_interval, which has a unit
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
func (dev *Device) Running(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

//The state machine executes this function prior to completion.
//
//The state machine does not perform any actions by default.  The
//application may set self.data.args.delete to True to cause the
//state machine to send a delete REST API request in order to
//remove the application's object from the Extreme Control service.
//
//A return code is not expected, and the execution of the state
//machine completes after the application returns from the function.
func (dev *Device) Done(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

//The state machine executes this function when the transmission of
//the ConfigAck message fails.  The application SHALL restore the
//configuration to its prior state.  If needed, the application MAY
//reboot the device.  If the application reboots the device, it
//SHALL store the reason for the reboot and inform the state machine
//by adding ConnectorEvent(s) to the self.data.args.reboot_events
//upon connector startup after the reboot.
//
//The state machine transmits the reboot events after the 'CONNECT'
//state has completed, but prior to issuing the 'CONNECT_OK'
//function.
//
//A return code is not expected, and the state machine transitions to
//CONNECT.
func (dev *Device) ConnectionExpired(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}

//This function allows the application to override the default
//routine for retrieving the source IP address of the socket used
//to communicate with Extreme Control.
//
//The override routine SHALL take two parameters, an IP address and
//a port number.
func (dev *Device) GetSourceIP(in interface{}) (ret ztp.DeviceReturnCode) {
	return ztp.DeviceReturnOk
}
