package main

import (
	msg "ztp"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "Snmp" with the name of the feature.
//
//e.g. s/Snmp/Vlan/
type devSnmp struct{}

//Return an instance of the feature interface
func NewDevSnmp() (f Feature) {
	f = devSnmp{}
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devSnmp) getCapability(m *msg.Capabilities) (err error) {
	m.Snmp.SnmpV1.FeatureAvailable = true
	m.Snmp.SnmpV2C.FeatureAvailable = true
	m.Snmp.SnmpV3.FeatureAvailable = true
	return
}

//Update any feature informatino in the Connect message
func (p devSnmp) getConnect(m *msg.Connect) (err error) {
	//m.DeviceInfo.Snmp.<somefield> =
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devSnmp) getConfig(m *msg.Configuration) (err error) {
	//m.ConfigBlock.Snmp.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devSnmp) setConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devSnmp) getStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.Snmp.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devSnmp) setStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the setConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.setConfig(&c)

	return
}
