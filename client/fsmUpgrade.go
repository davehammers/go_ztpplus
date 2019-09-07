package ztpclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
	msg "ztp"
)

func (zc *ZtpClient) Upgrade() (state ZtpClientState) {
	upgradeMsg := msg.ImageUpgrade{}

	// ask device to fill in the assets that need upgrading
	zc.Device.Upgrade(&upgradeMsg)

	b, err := json.Marshal(upgradeMsg)
	if err != nil {
		log.Println(err)
		return ZtpStateReConnectPause
	}

	r, err := http.NewRequest("PUT", zc.urlPrefix+"imageupgrade", bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		return ZtpStateReConnectPause
	}

	resp, body, err := zc.SendRequest(r)
	if err != nil {
		// networking issue
		log.Println(err)
		return ZtpStateReConnectPause
	}
	// unmarshal the response if present
	upgradeResp := &msg.ImageUpgradeResponse{}
	if body != nil {
		err = json.Unmarshal(*body, upgradeResp)
	}
	switch resp.StatusCode {
	case http.StatusOK:
	}

	switch zc.Device.UpgradeResponse(err, resp, upgradeResp) {
	case DeviceReturnOK:
		zc.SendEvents()
		return ZtpStateConfig
	case DeviceReturnRestart:
	case DeviceReturnRetry:
		return ZtpStateReUpgradePause
	case DeviceReturnFinish:
		return ZtpStateDone
	case DeviceReturnAbort:
	}

	return ZtpStateConfig
}

func (zc *ZtpClient) ReUpgradePause() (state ZtpClientState) {
	time.Sleep(time.Second * time.Duration(zc.upgradeRetry))
	return ZtpStateUpgrade
}
