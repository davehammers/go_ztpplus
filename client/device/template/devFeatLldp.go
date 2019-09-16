package main

import (
	msg "ztp"
	"ztp/client/device"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "Lldp" with the name of the feature.
//
//e.g. s/Lldp/Vlan/
type devLldp struct{}

//Return an instance of the feature interface
func NewDevLldp() (f device.Feature) {
	f = devLldp{}
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devLldp) GetConnect(m *msg.Connect) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devLldp) GetConfig(m *msg.Configuration) (err error) {
	m.Capabilities.Lldp.FeatureAvailable = true
	//m.ConfigBlock.Lldp.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devLldp) SetConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devLldp) GetStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.Lldp.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devLldp) SetStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the SetConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.SetConfig(&c)

	return
}
