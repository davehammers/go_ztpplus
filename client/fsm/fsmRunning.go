package ztpclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
	msg "ztp"
)

func (zc *ZtpClient) Running() (state ZtpClientState) {
	statsMsg := &msg.Stats{}
	// call device to fill in the stats message
	zc.Device.Running(statsMsg)

	b, err := json.Marshal(statsMsg)
	if err != nil {
		log.Println(err)
		return ZtpStateReRunningPause
	}

	r, err := http.NewRequest("PUT", zc.urlPrefix+"stats", bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		return ZtpStateReRunningPause
	}

	resp, body, err := zc.SendRequest(r)
	if err != nil {
		log.Println(err)
		return ZtpStateConnect
	}

	// unmarshal the response
	statsResponseMsg := &msg.StatsResponse{}
	err = json.Unmarshal(*body, &statsResponseMsg)
	switch zc.Device.RunningResponse(err, resp, statsResponseMsg) {
	case DeviceReturnOK:
		return ZtpStateReRunningPause
	case DeviceReturnRestart:
		return ZtpStateConnect
	//case DeviceReturnRetry:
	case DeviceReturnFinish:
		return ZtpStateDone
	//case DeviceReturnAbort:
	default:
		return ZtpStateInit
	}
	return ZtpStateReRunningPause
}

func (zc *ZtpClient) ReRunningPause() (state ZtpClientState) {
	log.Println("sleep", zc.runningRetry)
	time.Sleep(time.Second * time.Duration(zc.runningRetry))
	return ZtpStateRunning
}
