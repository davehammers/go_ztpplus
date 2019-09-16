package main

import (
	msg "ztp"
	"ztp/client/device"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "Vlans" with the name of the feature.
//
//e.g. s/Vlans/Vlan/
type devVlans struct{}

//Return an instance of the feature interface
func NewDevVlans() (f device.Feature) {
	f = devVlans{}
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devVlans) GetConnect(m *msg.Connect) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devVlans) GetConfig(m *msg.Configuration) (err error) {
	m.Capabilities.Vlans.FeatureAvailable = true
	//m.ConfigBlock.Vlans.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devVlans) SetConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devVlans) GetStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.Vlans.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devVlans) SetStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the SetConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.SetConfig(&c)

	return
}
