package ztpclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
	msg "ztp"
)

func (zc *ZtpClient) Connect() (state ZtpClientState) {
	var connectMsg msg.Connect

	// call the device routine to fill in the connect message
	zc.Device.Connect(&connectMsg)

	b, err := json.Marshal(connectMsg)
	if err != nil {
		log.Println(err)
		zc.Device.ConnectResponse(err, nil, nil)
		return ZtpStateReDiscoverPause
	}

	r, err := http.NewRequest("PUT", zc.urlPrefix+"connect", bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		zc.Device.ConnectResponse(err, nil, nil)
		return ZtpStateReDiscoverPause
	}

	resp, body, err := zc.SendRequest(r)

	if err != nil {
		zc.Device.ConnectResponse(err, resp, nil)
		return ZtpStateReDiscoverPause
	}

	connectResp := &msg.ConnectResponse{}
	if body != nil {
		err = json.Unmarshal(*body, connectResp)
	}

	// call the device function with the filled in response data
	// the device has the opportunity to modify the response contents
	switch zc.Device.ConnectResponse(err, resp, connectResp) {
	case DeviceReturnOK:
		return ZtpStateUpgrade
	case DeviceReturnRestart:
		// URL redirect
		return ZtpStateConnect
	case DeviceReturnRetry:
		return ZtpStateReConnectPause
	case DeviceReturnFinish:
	case DeviceReturnAbort:
	}
	return ZtpStateUpgrade
}

// a state that pauses for connectRetry seconds before returning to the ZtpStateConnect state
func (zc *ZtpClient) ReConnectPause() (state ZtpClientState) {
	time.Sleep(time.Second * time.Duration(zc.connectRetry))
	return ZtpStateConnect
}
