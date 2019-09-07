package ztpclient

import ()

func (zc *ZtpClient) Done() (state ZtpClientState) {

	state = ZtpStateDone

	return state
}
