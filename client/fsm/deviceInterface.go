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
	DeviceReturnAbort
)

type Device interface {
	StartFSM()
	Init()
	Discover() *[]ZtpLookupEntry
	DiscoverResponse(*http.Response, *ZtpLookupEntry) DeviceReturnCode
	Connect(*msg.Connect)
	ConnectResponse(error, *http.Response, *msg.ConnectResponse) DeviceReturnCode
	Upgrade(*msg.ImageUpgrade)
	UpgradeResponse(error, *http.Response, *msg.ImageUpgradeResponse) DeviceReturnCode
	Config(*msg.Configuration)
	ConfigResponse(error, *http.Response, *msg.ConfigurationResponse) DeviceReturnCode
	ConfigAck() DeviceReturnCode
	PostConfig() DeviceReturnCode
	Running(*msg.Stats)
	RunningResponse(error, *http.Response, *msg.StatsResponse) DeviceReturnCode
	Done()
	ConnectionExpired()
	GetSourceIP()
	SendEvents(*msg.EventMsg) DeviceReturnCode
	SendEventsComplete()
}
