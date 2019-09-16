package main

import (
	msg "ztp"
	"ztp/client/device"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "Poe" with the name of the feature.
//
//e.g. s/Poe/Vlan/
type devPoe struct{}

//Return an instance of the feature interface
func NewDevPoe() (f device.Feature) {
	f = devPoe{}
	return
}

// feature collects and updates the device specific DB with config data
func (p devPoe) GetDBConfig() (err error) {
	return
}

// feature collects and updates the device specific DB with config statistics data
func (p devPoe) GetDBStats() (err error) {
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devPoe) GetConnect(m *msg.Connect) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devPoe) GetConfig(m *msg.Configuration) (err error) {
	m.Capabilities.Poe.FeatureAvailable = true
	//m.ConfigBlock.Poe.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devPoe) SetConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devPoe) GetStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.Poe.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devPoe) SetStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the SetConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.SetConfig(&c)

	return
}
