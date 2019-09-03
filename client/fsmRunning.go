package ztpclient

import (
	"log"
)

func (zc *ZtpClient) Running() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

	state = ZtpStateDone

	return state
}
