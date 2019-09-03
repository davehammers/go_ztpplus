package ztpclient

import (
	"log"
)

func (zc *ZtpClient) PostConfig() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

	state = ZtpStateConfigAck

	return state
}
