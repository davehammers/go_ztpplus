package ztpclient

import (
	"time"
)

//This routine processes the any device specific functions after the configuration is applied.
func (zc *ZtpClient) PostConfig() (state ZtpClientState) {
	// call the device function to process PostConfig activities
	switch zc.Device.PostConfig() {
	case DeviceReturnOK:
	case DeviceReturnRestart:
		return ZtpStateReConnectPause
	case DeviceReturnRetry:
		return ZtpStateRePostConfigPause
	case DeviceReturnFinish:
		return ZtpStateConfigAck
	case DeviceReturnAbort:
		zc.Device.ConnectionExpired()
		return ZtpStateReConnectPause
	}

	return ZtpStateConfigAck
}

// a state that pauses for postConfigRetry seconds before returning to the ZtpStateConnect state
func (zc *ZtpClient) RePostConfigPause() (state ZtpClientState) {
	time.Sleep(time.Second * time.Duration(zc.postConfigRetry))
	return ZtpStatePostConfig
}
