package ztpclient

import (
	"log"
)

func (zc *ZtpClient) Config() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

	state = ZtpStatePostConfig

	return state
}
