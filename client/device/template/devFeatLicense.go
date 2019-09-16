package main

import (
	msg "ztp"
	"ztp/client/device"
)

type devLicense struct{}

func NewDevLicense() (f device.Feature) {
	f = devLicense{}
	return
}

func (p devLicense) GetCapability(m *msg.Capabilities) (err error) {
	m.License.FeatureAvailable = true
	return
}

func (p devLicense) GetConnect(m *msg.Connect) (err error) {
	m.DeviceInfo.License.FeaturePacks = make([]msg.FeaturePack, 0)
	m.DeviceInfo.License.PortCapacityLicense = ""
	m.DeviceInfo.License.EffectiveLicense = ""
	m.DeviceInfo.License.EnabledLicense = ""
	m.DeviceInfo.License.SystemLicense = ""
	m.DeviceInfo.License.EffectiveLevel = ""
	return
}
func (p devLicense) GetConfig(m *msg.Configuration) (err error) {
	m.ConfigBlock.License.SystemLicense = ""
	m.ConfigBlock.License.EffectiveLevel = ""
	m.ConfigBlock.License.FeaturePacks = make([]msg.FeaturePack, 0)
	return
}

func (p devLicense) SetConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

func (p devLicense) GetStats(m *msg.Stats) (err error) {
	return
}

func (p devLicense) SetStats(m *msg.StatsResponse) (err error) {
	return
}
