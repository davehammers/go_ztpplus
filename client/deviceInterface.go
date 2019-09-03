package ztpclient

import (
	msg "ztp"
)

type DeviceReturnCode int

const (
	DeviceReturnOk DeviceReturnCode = 0 + iota
	DeviceReturnRestart
	DeviceReturnRetry
	DeviceReturnFinish
	DeviceReturnDone
)

type Device interface {
	Init()
	Discover(interface{})
	DiscoverOK(*ZtpLookupEntry)
	DiscoverFail() DeviceReturnCode
	Connect(*msg.Connect)
	ConnectOk(interface{})
	ConnectFail(interface{}) DeviceReturnCode
	ConnectRedirect(interface{}) DeviceReturnCode
	UpgradeCheck(interface{}) DeviceReturnCode
	Upgrade(interface{}) (DeviceReturnCode, *[]msg.Event)
	ConfigUpdate(interface{})
	Config(interface{}) (events *[]msg.Event)
	PostConfig(interface{}) DeviceReturnCode
	ConfigOK(interface{}) DeviceReturnCode
	ConfigFail(interface{}) DeviceReturnCode
	Running(interface{}) DeviceReturnCode
	Done(interface{}) DeviceReturnCode
	ConnectionExpired(interface{}) DeviceReturnCode
	GetSourceIP(interface{}) DeviceReturnCode
}
