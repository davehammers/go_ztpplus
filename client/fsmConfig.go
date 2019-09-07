package ztpclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
	msg "ztp"
)

func (zc *ZtpClient) Config() (state ZtpClientState) {
	config := &msg.Configuration{}
	// call device function to fill in the config structure
	zc.Device.Config(config)

	b, err := json.Marshal(config)
	if err != nil {
		log.Println(err)
		return ZtpStateReConnectPause
	}

	r, err := http.NewRequest("PUT", zc.urlPrefix+"configuration", bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		return ZtpStateReConnectPause
	}
	if DEBUG {
		msg.DumpJson(config)
	}

	resp, body, err := zc.SendRequest(r)
	if err != nil {
		log.Println(err)
		return ZtpStateReConfigPause
	}

	// unmarshal the response
	configResp := &msg.ConfigurationResponse{}
	err = json.Unmarshal(*body, configResp)

	switch zc.Device.ConfigResponse(err, resp, configResp) {
	case DeviceReturnOK:
	case DeviceReturnRestart:
		return ZtpStateReConnectPause
	case DeviceReturnRetry:
		return ZtpStateReConfigPause
	case DeviceReturnFinish:
		return ZtpStateDone
	case DeviceReturnAbort:
		return ZtpStateInit
	}

	// send any events that config may have encountered
	zc.SendEvents()
	return ZtpStatePostConfig
}
func (zc *ZtpClient) ReConfigPause() (state ZtpClientState) {
	time.Sleep(time.Second * time.Duration(zc.configRetry))
	return ZtpStateConfig
}
