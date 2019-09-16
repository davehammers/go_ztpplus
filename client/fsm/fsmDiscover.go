package ztpclient

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//This routine performs the 'DISCOVER' state actions.
//
//Discovers the Extreme Control service by performing a DNS lookup of
//'extremecontrol.<local_domain>'.  The returned list of IP addresses
//are scanned one at a time, and the first one to successfully
//communicate with the Extreme Control service is used.
//
//If a local Extreme Control service is not found, the state machine
//tries to discover an Extreme Control service in the cloud via
//'devices.extremenetworks.com'.
//
//The application may configure a fallback controller if the discovery
//mechanism fails.  The fallback controller is configured by setting the
//self.data.args.controller_addr variable.
func (zc *ZtpClient) Discover() (state ZtpClientState) {
	var resp *http.Response
	var err error
	//get the list of controllers from the platform
	controllerList := zc.Device.Discover()
	for _, hostPort := range *controllerList {
		url := fmt.Sprintf("https://%s%s/device/v1/", hostPort.Host, hostPort.Port)
		if hostPort.AddClientPrefix {
			url = fmt.Sprintf("https://%s%s/Clients/device/v1/", hostPort.Host, hostPort.Port)
		}
		if DEBUG {
			log.Println(url + "discovery")
		}
		resp, err = zc.httpClient.Get(url + "discovery")

		// Tell the device
		switch zc.Device.DiscoverResponse(err, resp, &hostPort) {
		case DeviceReturnOK:
			zc.urlPrefix = fmt.Sprintf("%s%s/%s/", url, zc.devType, zc.devID)
			return ZtpStateConnect
		case DeviceReturnRestart:
			return ZtpStateReDiscoverPause
		case DeviceReturnRetry:
			continue
		case DeviceReturnFinish:
			return ZtpStateDone
		case DeviceReturnAbort:
			return ZtpStateReDiscoverPause
		}
	}
	zc.Device.DiscoverResponse(err, nil, nil)
	return ZtpStateReDiscoverPause
}

func (zc *ZtpClient) ReDiscoverPause() (state ZtpClientState) {
	// pause here before trying discover again
	time.Sleep(time.Second * time.Duration(zc.discoverRetry))
	return ZtpStateDiscover
}
