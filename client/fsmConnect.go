package ztpclient

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
	msg "ztp"
)

var (
	des3Key []byte = []byte("99a3d014b62293332b85c3ab")
)

func (zc *ZtpClient) Connect() (state ZtpClientState) {
	var connectMsg msg.Connect
	if DEBUG {
		log.Println("Begin")
	}

	// call the device routine to fill in the connect message
	zc.device.Connect(&connectMsg)

	b, err := json.Marshal(connectMsg)
	if err != nil {
		log.Println(err)
		zc.device.ConnectFail(nil)
		return ZtpStateReDiscoverPause
	}

	r, err := http.NewRequest("PUT", zc.urlPrefix+"connect", bytes.NewReader(b))
	if err != nil {
		log.Println(err)
		zc.device.ConnectFail(nil)
		return ZtpStateReDiscoverPause
	}
	if DEBUG {
		x, _ := httputil.DumpRequest(r, true)
		log.Println(string(x))
	}

	// use initial constant for auth
	userName := des3Encrypt("ezconfiguser")
	r.SetBasicAuth(userName, userName)
	r.Header.Add("Content-type", "application/json")
	resp, err := zc.httpClient.Do(r)

	if err != nil {
		log.Println(err)
		zc.device.ConnectFail(resp)
		return ZtpStateReDiscoverPause
	}
	if DEBUG {
		log.Println("StatusCode", resp.StatusCode)
		x, _ := httputil.DumpResponse(resp, true)
		log.Println(string(x))
	}
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated:
		// fall below the switch statement to process the message
		break
	default:
		if resp.StatusCode >= http.StatusInternalServerError {
			// notify device
			zc.device.ConnectFail(resp)
			return ZtpStateReDiscoverPause
		}
		zc.device.ConnectFail(resp)
		return ZtpStateReConnectPause
	}

	// unmarshal the response
	connectResp := msg.ConnectResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&connectResp); err != nil {
		log.Println(err)
		zc.device.ConnectFail(resp)
		return ZtpStateReConnectPause
	}
	if DEBUG {
		log.Printf("%+v\n", connectResp)
	}
	// if new auth is provided, decode the login/password provided
	if connectResp.Credentials.Login != "" && connectResp.Credentials.Password != "" {
		zc.login = des3Decrypt(connectResp.Credentials.Login)
		zc.password = des3Decrypt(connectResp.Credentials.Password)
		// trim 3 characters from either side of the password
		zc.login = des3Encrypt(zc.login[3 : len(zc.login)-3])
		zc.password = des3Encrypt(zc.password[3 : len(zc.password)-3])
	}

	// is there a redirected to a different controller
	// ignore case when comparing strings
	if strings.EqualFold(connectResp.Action, "redirect") {
		if strings.EqualFold(connectResp.Redirect.Type, "https") {
			if connectResp.Redirect.URI != "" {
				// rebuild the prefix with the new URI as host
				if u, err := url.Parse(zc.urlPrefix); err == nil {
					u.Host = connectResp.Redirect.URI
					zc.urlPrefix = u.String()
					zc.device.ConnectRedirect(resp, &connectResp)
					// try to connect with the new controller
					return ZtpStateConnect
				}
			}
		}
	}
	zc.device.ConnectOK(resp, &connectResp)
	return ZtpStateUpgrade
}

// a state that pauses for connectRetry seconds before returning to the ZtpStateConnect state
func (zc *ZtpClient) ReConnectPause() (state ZtpClientState) {
	time.Sleep(time.Second * time.Duration(zc.connectRetry))
	return ZtpStateConnect
}

func des3Encrypt(in string) (out string) {
	des3, err := des.NewTripleDESCipher(des3Key)
	if err != nil {
		log.Println(err)
		return ""
	}
	padLen := byte(des3.BlockSize() - (len(in) % des3.BlockSize()))
	pad := bytes.Repeat([]byte{padLen}, int(padLen))
	plaintext := append([]byte(in), pad...)

	des3Out := make([]byte, des3.BlockSize())
	mode := cipher.NewCBCEncrypter(des3, des3Out)
	encrypted := make([]byte, len(plaintext))
	mode.CryptBlocks(encrypted, plaintext)
	out = base64.StdEncoding.EncodeToString(encrypted)
	/*
		if DEBUG {
			log.Printf("%s encrypt to %x base64 %s\n", plaintext, encrypted, out)
		}
	*/
	return
}
func des3Decrypt(in string) (out string) {
	encrypted, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		log.Println(err)
		return ""
	}
	des3, err := des.NewTripleDESCipher(des3Key)
	if err != nil {
		log.Println(err)
		return ""
	}
	des3Out := make([]byte, des3.BlockSize())
	decrypter := cipher.NewCBCDecrypter(des3, des3Out)
	decrypted := make([]byte, len(encrypted))
	decrypter.CryptBlocks(decrypted, encrypted)

	trimLen := decrypted[len(decrypted)-1]
	out = string(decrypted[:len(decrypted)-int(trimLen)])
	/*
		if DEBUG {
			log.Printf("%x decrypt to %s trim %d out %s\n", encrypted, decrypted, trimLen, out)
		}
	*/
	return
}
