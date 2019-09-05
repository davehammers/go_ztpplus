package ztpclient

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
	msg "ztp"
)

func (zc *ZtpClient) Upgrade() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}
	upgradeMsg := msg.ImageUpgrade{}
	upgradeMsg.ApPropertyBlock = zc.property

	// ask device to fill in the assets that need upgrading
	switch zc.device.UpgradeCheck(&upgradeMsg) {
	case DeviceReturnRetry:
		return ZtpStateReUpgradePause
	}

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
	if DEBUG {
		x, _ := httputil.DumpRequest(r, true)
		log.Println(string(x))
	}

	// use initial constant for auth
	r.SetBasicAuth(zc.login, zc.password)
	r.Header.Add("Content-type", "application/json")
	resp, err := zc.httpClient.Do(r)
	if err != nil {
		log.Println(err)
		return ZtpStateReUpgradePause
	}
    if DEBUG {
		log.Println("StatusCode", resp.StatusCode)
		x, _ := httputil.DumpResponse(resp, true)
		log.Println(string(x))
    }
	switch resp.StatusCode {
	case http.StatusOK:
		// OK HTTP response, keep going
	default:
		return ZtpStateReConnectPause
	}

	// unmarshal the response
	upgradeResp := msg.ConnectResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&upgradeResp); err != nil {
		log.Println(err)
		return ZtpStateReConnectPause
	}
	ret, events := zc.device.Upgrade(&upgradeResp)
	switch ret {
	case DeviceReturnOK:
		if events != nil {
			zc.SendEvents(events)
		}
		return ZtpStateConfig
	//case DeviceReturnRestart:
	case DeviceReturnRetry:
		return ZtpStateReUpgradePause
	case DeviceReturnFinish:
		return ZtpStateDone
	case DeviceReturnDone:
		return ZtpStateDone
	default:
		return ZtpStateReUpgradePause
	}

	return ZtpStateConfig
}

func (zc *ZtpClient) ReUpgradePause() (state ZtpClientState) {
	time.Sleep(time.Second * time.Duration(zc.upgradeRetry))
	return ZtpStateUpgrade
}
