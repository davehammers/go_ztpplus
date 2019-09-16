package main

import (
	msg "ztp"
	"ztp/client/device"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "ConfigDownload" with the name of the feature.
//
//e.g. s/ConfigDownload/Vlan/
type devConfigDownload struct{}

//Return an instance of the feature interface
func NewDevConfigDownload() (f device.Feature) {
	f = devConfigDownload{}
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devConfigDownload) GetCapability(m *msg.Capabilities) (err error) {
	m.ConfigDownload.FeatureAvailable = true
	return
}

//Update any feature informatino in the Connect message
func (p devConfigDownload) GetConnect(m *msg.Connect) (err error) {
	//m.DeviceInfo.ConfigDownload.<somefield> =
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devConfigDownload) GetConfig(m *msg.Configuration) (err error) {
	//m.ConfigBlock.ConfigDownload.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devConfigDownload) SetConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devConfigDownload) GetStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.ConfigDownload.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devConfigDownload) SetStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the SetConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.SetConfig(&c)

	return
}
