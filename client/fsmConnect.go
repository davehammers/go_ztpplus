package ztpclient

import (
	"bytes"
	//"crypto/des"
	"encoding/json"
	"log"
	"net/http"
	msg "ztp"
)

func (zc *ZtpClient) Connect() (state ZtpClientState) {
	var connectMsg msg.Connect
	if DEBUG {
		log.Println("Begin")
	}

	// call the device routine to fill in the connect message
	zc.device.Connect(&connectMsg)
	zc.property = connectMsg.ApPropertyBlock

	//userName, _ := des.NewTripleDESCipher([]byte("ezconfiguser"))
	userName := "ezconfiguser"
	b, err := json.Marshal(connectMsg)
	if err != nil {
		return ZtpStateDiscover
	}

	r, err := http.NewRequest("PUT", zc.urlPrefix+"connect", bytes.NewReader(b))
	if err != nil {
		return ZtpStateDiscover
	}

	r.SetBasicAuth(userName, userName)
	resp, err := zc.httpClient.Do(r)

	if err != nil {
		if DEBUG {
			log.Println(err)
		}
		return ZtpStateDiscover
	}
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated:
		// unmarshal the response
		connectResp := msg.ConnectResponse{}
		if err := json.NewDecoder(resp.Body).Decode(&connectResp); err != nil {
			// TODO pause before starting again
			return ZtpStateConnect
		}
	}

	state = ZtpStateUpgrade

	return state
}
