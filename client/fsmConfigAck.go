package ztpclient

import (
	"log"
)

func (zc *ZtpClient) ConfigAck() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin", zc.state)
	}

	state = ZtpStateRunning

	return state
}
