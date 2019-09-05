package ztpclient

import (
	"log"
	"time"
)

func (zc *ZtpClient) Running() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

	state = ZtpStateDone

	return state
}

func (zc *ZtpClient) ReRunningPause() (state ZtpClientState) {
	time.Sleep(time.Second * time.Duration(zc.runningRetry))
	return ZtpStateRunning
}
