package ztpclient

import (
	"log"
)

func (zc *ZtpClient) Done() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

	state = ZtpStateDone

	return state
}
