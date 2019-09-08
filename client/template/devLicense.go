package main

import (
	msg "ztp"
)

type devLicense struct{}

func NewDevLicense() (f Feature) {
	f = devLicense{}
	return
}

func (p devLicense) getCapability(m *msg.Capabilities) (err error) {
	m.License.FeatureAvailable = true
	return
}

func (p devLicense) getConnect(m *msg.Connect) (err error) {
	m.DeviceInfo.License.FeaturePacks = make([]msg.FeaturePack, 0)
	m.DeviceInfo.License.PortCapacityLicense = ""
	m.DeviceInfo.License.EffectiveLicense = ""
	m.DeviceInfo.License.EnabledLicense = ""
	m.DeviceInfo.License.SystemLicense = ""
	m.DeviceInfo.License.EffectiveLevel = ""
	return
}
func (p devLicense) getConfig(m *msg.Configuration) (err error) {
	m.ConfigBlock.License.SystemLicense = ""
	m.ConfigBlock.License.EffectiveLevel = ""
	m.ConfigBlock.License.FeaturePacks = make([]msg.FeaturePack, 0)
	return
}

func (p devLicense) setConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

func (p devLicense) getStats(m *msg.Stats) (err error) {
	return
}

func (p devLicense) setStats(m *msg.StatsResponse) (err error) {
	return
}
