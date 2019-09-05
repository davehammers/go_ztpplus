package ztpclient

import (
	"log"
	"time"
)

func (zc *ZtpClient) Config() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

	state = ZtpStatePostConfig

	return state
}
func (zc *ZtpClient) ReConfigPause() (state ZtpClientState) {
	time.Sleep(time.Second * time.Duration(zc.configRetry))
	return ZtpStateConfig
}
