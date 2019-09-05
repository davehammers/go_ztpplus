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
	if DEBUG {
		log.Println("Begin")
	}
	for _, hostPort := range controllerList {
		url := fmt.Sprintf("https://%s%s/device/v1/", hostPort.host, hostPort.port)
		if hostPort.addClientPrefix {
			url = fmt.Sprintf("https://%s%s/Clients/device/v1/", hostPort.host, hostPort.port)
		}
		if DEBUG {
			log.Println(url + "discovery")
		}
		resp, err := zc.httpClient.Get(url + "discovery")
		if err != nil {
			if DEBUG {
				log.Println(err)
			}
			continue
		}

		switch resp.StatusCode {
		case http.StatusOK, http.StatusNotFound:
			// discovery is successful
			zc.device.DiscoverOK(&hostPort)
			// build up the URL prefix for all other message types
			zc.urlPrefix = fmt.Sprintf("%s%s/%s/", url, zc.devType, zc.devID)
			if DEBUG {
				log.Println(zc.urlPrefix)
			}
			return ZtpStateConnect
		default:
			continue
		}
	}
	// discovery failed
	switch zc.device.DiscoverFail() {
	case DeviceReturnFinish:
		return ZtpStateDone
	}

	return ZtpStateReDiscoverPause
}

func (zc *ZtpClient) ReDiscoverPause() (state ZtpClientState) {
	// pause here before trying discover again
	time.Sleep(time.Second * time.Duration(zc.discoverRetry))
	return ZtpStateDiscover
}
