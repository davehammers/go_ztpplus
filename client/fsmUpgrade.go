package ztpclient

import (
	"log"
	"time"
)

func (zc *ZtpClient) Upgrade() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

	state = ZtpStateConfig

	return state
}

func (zc *ZtpClient) ReUpgradePause() (state ZtpClientState) {
	time.Sleep(time.Second * time.Duration(zc.upgradeRetry))
	return ZtpStateUpgrade
}
