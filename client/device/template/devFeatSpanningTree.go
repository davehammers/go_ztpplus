package main

import (
	msg "ztp"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "SpanningTree" with the name of the feature.
//
//e.g. s/SpanningTree/Vlan/
type devSpanningTree struct{}

//Return an instance of the feature interface
func NewDevSpanningTree() (f Feature) {
	f = devSpanningTree{}
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devSpanningTree) getCapability(m *msg.Capabilities) (err error) {
	m.SpanningTree.FeatureAvailable = true
	return
}

//Update any feature informatino in the Connect message
func (p devSpanningTree) getConnect(m *msg.Connect) (err error) {
	//m.DeviceInfo.SpanningTree.<somefield> =
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devSpanningTree) getConfig(m *msg.Configuration) (err error) {
	//m.ConfigBlock.SpanningTree.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devSpanningTree) setConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devSpanningTree) getStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.SpanningTree.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devSpanningTree) setStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the setConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.setConfig(&c)

	return
}
