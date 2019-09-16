package main

import (
	msg "ztp"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "Syslog" with the name of the feature.
//
//e.g. s/Syslog/Vlan/
type devSyslog struct{}

//Return an instance of the feature interface
func NewDevSyslog() (f Feature) {
	f = devSyslog{}
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devSyslog) getCapability(m *msg.Capabilities) (err error) {
	m.Syslog.FeatureAvailable = true
	return
}

//Update any feature informatino in the Connect message
func (p devSyslog) getConnect(m *msg.Connect) (err error) {
	//m.DeviceInfo.Syslog.<somefield> =
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devSyslog) getConfig(m *msg.Configuration) (err error) {
	//m.ConfigBlock.Syslog.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devSyslog) setConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devSyslog) getStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.Syslog.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devSyslog) setStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the setConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.setConfig(&c)

	return
}
