package main

import (
	msg "ztp"
	"ztp/client/device"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "MacAuth" with the name of the feature.
//
//e.g. s/MacAuth/Vlan/
type devMacAuth struct{}

//Return an instance of the feature interface
func NewDevMacAuth() (f device.Feature) {
	f = devMacAuth{}
	return
}

// feature collects and updates the device specific DB with config data
func (p devMacAuth) GetDBConfig() (err error) {
	return
}

// feature collects and updates the device specific DB with config statistics data
func (p devMacAuth) GetDBStats() (err error) {
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devMacAuth) GetConnect(m *msg.Connect) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devMacAuth) GetConfig(m *msg.Configuration) (err error) {
	m.Capabilities.MacAuth.FeatureAvailable = true
	//m.ConfigBlock.MacAuth.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devMacAuth) SetConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devMacAuth) GetStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.MacAuth.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devMacAuth) SetStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the SetConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.SetConfig(&c)

	return
}
