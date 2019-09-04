package ztpclient

import (
	"bytes"
	"crypto/des"
	"crypto/cipher"
	"encoding/json"
	"encoding/base64"
	"log"
	"net/http"
	"net/http/httputil"
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
	zc.property = connectMsg.ApPropertyBlock

	b, err := json.Marshal(connectMsg)
	if err != nil {
		return ZtpStateDiscover
	}

	r, err := http.NewRequest("PUT", zc.urlPrefix+"connect", bytes.NewReader(b))
	if err != nil {
		if DEBUG {
			log.Println(err)
		}
		return ZtpStateDiscover
	}
	if DEBUG {
		x, _ := httputil.DumpRequest(r, true)
		log.Println(string(x))
	}

	userName := des3Encrypt("ezconfiguser")
	//userName := "vB65SDNui+3fN2/Jvh9Reg=="
    log.Println("userName", userName)
	r.SetBasicAuth(userName, userName)
	r.Header.Add("Content-type", "application/json")
	resp, err := zc.httpClient.Do(r)

	if err != nil {
		if DEBUG {
			log.Println(err)
		}
		return ZtpStateDiscover
	}
	if DEBUG {
		x, _ := httputil.DumpResponse(resp, true)
		log.Println(string(x))
	}
	switch resp.StatusCode {
	case http.StatusOK, http.StatusCreated:
		// fall below the switch
		break
	default:
		if resp.StatusCode >= http.StatusInternalServerError {
			return ZtpStateDiscover
		}
		return ZtpStateConnect
	}

	// unmarshal the response
	connectResp := msg.ConnectResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&connectResp); err != nil {
		// TODO pause before starting again
		if DEBUG {
			log.Println(err)
		}
		return ZtpStateConnect
	}
	if DEBUG {
		log.Printf("%+v\n", connectResp)
	}
	// decode the password provided
	if connectResp.Credentials.Login != "" {
        zc.login = des3Decrypt(connectResp.Credentials.Login)
        zc.password = des3Decrypt(connectResp.Credentials.Password)
        zc.login = des3Encrypt(zc.login[3 : len(zc.login)-3])
        zc.password = des3Encrypt(zc.password[3 : len(zc.password)-3])
	}

	return ZtpStateUpgrade
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
    if DEBUG {
        log.Printf("%s encrypt to %x base64 %s\n", plaintext, encrypted, out)
    }
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

    trimLen := decrypted[len(decrypted) - 1]
    out = string(decrypted[:len(decrypted) - int(trimLen)])
    if DEBUG {
        log.Printf("%x decrypt to %s trim %d out %s\n", encrypted, decrypted, trimLen, out)
    }
	return
}
