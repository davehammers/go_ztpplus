package ztpclient

import ()

func (zc *ZtpClient) ConfigAck() (state ZtpClientState) {
	// ask the device what to do once the config ACK is sent
	ret := zc.Device.ConfigAck()

	// send any pending events to the controller
	err := zc.SendEvents()
	if err != nil {
		zc.Device.ConnectionExpired()
		// try to reconnect
		return ZtpStateReConnectPause
	}

	// the device told us how the config went
	switch ret {
	case DeviceReturnOK:
		return ZtpStateRunning
	case DeviceReturnRestart:
	case DeviceReturnRetry:
		return ZtpStateReConfigPause
	case DeviceReturnFinish:
	case DeviceReturnAbort:
	}
	return ZtpStateRunning
}
