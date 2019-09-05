package ztpclient

import (
	"net/http"
	msg "ztp"
)

type DeviceReturnCode int

const (
	DeviceReturnOK DeviceReturnCode = 0 + iota
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
	ConnectOK(*http.Response, *msg.ConnectResponse)
	ConnectFail(*http.Response) DeviceReturnCode
	ConnectRedirect(*http.Response, *msg.ConnectResponse) DeviceReturnCode
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
