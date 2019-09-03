package ztpclient

import (
	"log"
)

func (zc *ZtpClient) Upgrade() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

	state = ZtpStateConfig

	return state
}
