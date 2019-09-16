package main

import (
	msg "ztp"
	"ztp/client/device"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "Dot1X" with the name of the feature.
//
//e.g. s/Dot1X/Vlan/
type devDot1X struct{}

//Return an instance of the feature interface
func NewDevDot1X() (f device.Feature) {
	f = devDot1X{}
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devDot1X) GetConnect(m *msg.Connect) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devDot1X) GetConfig(m *msg.Configuration) (err error) {
	m.Capabilities.Dot1X.FeatureAvailable = true
	//m.ConfigBlock.Dot1X.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devDot1X) SetConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devDot1X) GetStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.Dot1X.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devDot1X) SetStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the SetConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.SetConfig(&c)

	return
}
