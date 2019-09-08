package main

import (
	msg "ztp"
)

type platLicense struct{}

func NewPlatLicense() (f Feature) {
	f = &platLicense{}
	return
}

func (p *platLicense) getCapability(m *msg.Capabilities) (err error) {
	m.License.FeatureAvailable = true
	return
}

func (p *platLicense) getConnect(m *msg.Connect) (err error) {
	m.DeviceInfo.License.FeaturePacks = make([]msg.FeaturePack, 0)
	m.DeviceInfo.License.PortCapacityLicense = ""
	m.DeviceInfo.License.EffectiveLicense = ""
	m.DeviceInfo.License.EnabledLicense = ""
	m.DeviceInfo.License.SystemLicense = ""
	m.DeviceInfo.License.EffectiveLevel = ""
	return
}
func (p *platLicense) getConfig(m *msg.ConfigBlock) (err error) {
	m.License.SystemLicense = ""
	m.License.EffectiveLevel = ""
	m.License.FeaturePacks = make([]msg.FeaturePack, 0)
	return
}

func (p *platLicense) setConfig(m *msg.ConfigBlock) (err error) {
	return
}
