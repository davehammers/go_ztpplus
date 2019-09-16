package main

import (
	msg "ztp"
	"ztp/client/device"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "Syslog" with the name of the feature.
//
//e.g. s/Syslog/Vlan/
type devSyslog struct{}

//Return an instance of the feature interface
func NewDevSyslog() (f device.Feature) {
	f = devSyslog{}
	return
}

// feature collects and updates the device specific DB with config data
func (p devSyslog) GetDBConfig() (err error) {
	return
}

// feature collects and updates the device specific DB with config statistics data
func (p devSyslog) GetDBStats() (err error) {
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devSyslog) GetConnect(m *msg.Connect) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devSyslog) GetConfig(m *msg.Configuration) (err error) {
	m.Capabilities.Syslog.FeatureAvailable = true
	//m.ConfigBlock.Syslog.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devSyslog) SetConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devSyslog) GetStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.Syslog.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devSyslog) SetStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the SetConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.SetConfig(&c)

	return
}
