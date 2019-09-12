package ztpclient

import (
	"crypto/tls"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
	msg "ztp"
)

// Environment variable ztpclientDEBUG=1 sets the package DEBUG=true
const (
	EnvDebug              = "ztpclientDEBUG"
	DefaultRetrySeconds   = 30
	DefaultRunningSeconds = 60 * 5
)

// FSM states
type ZtpClientState int

const (
	ZtpStateInit ZtpClientState = 1 + iota
	ZtpStateDiscover
	ZtpStateReDiscoverPause // pause before returning to ZtpStateDiscover
	ZtpStateConnect
	ZtpStateReConnectPause // pause before returning to ZtpStateConnect
	ZtpStateUpgrade
	ZtpStateReUpgradePause // pause before returning to ZtpStateUpgrade
	ZtpStateConfig
	ZtpStateReConfigPause // pause before returning to ZtpStateConfig
	ZtpStatePostConfig
	ZtpStateRePostConfigPause // pause before returning to ZtpStatePostConfig
	ZtpStateConfigAck
	ZtpStateRunning
	ZtpStateReRunningPause // pause before returning to ZtpStateRunning
	ZtpStateDone
)

type ZtpClient struct {
	Device          Device // the device specific interfaces
	devID           string
	devType         string
	rebootEvent     string
	discoverRetry   int // seconds before retrying discovery
	connectRetry    int // seconds before retrying connect
	upgradeRetry    int // seconds before retrying upgrade
	configRetry     int // seconds before retrying config
	runningRetry    int // seconds before retrying running
	eventRetry      int // seconds before retrying running
	postConfigRetry int // seconds before retrying postConfig
	login           string
	password        string
	httpClient      *http.Client
	urlPrefix       string // Once discovered the URL prefix is the same for all transactions
	// [state] function table
	fsm map[ZtpClientState]func(*ZtpClient) ZtpClientState
}

// create a new ZTP client instance
func NewZtpClient() *ZtpClient {
	zc := &ZtpClient{
		devType:         "switch", // default type
		discoverRetry:   DefaultRetrySeconds,
		connectRetry:    DefaultRetrySeconds,
		upgradeRetry:    DefaultRetrySeconds,
		configRetry:     DefaultRetrySeconds,
		runningRetry:    DefaultRunningSeconds,
		eventRetry:      DefaultRetrySeconds,
		postConfigRetry: DefaultRetrySeconds,
		fsm: map[ZtpClientState]func(*ZtpClient) ZtpClientState{
			ZtpStateInit:              (*ZtpClient).Init,
			ZtpStateDiscover:          (*ZtpClient).Discover,
			ZtpStateReDiscoverPause:   (*ZtpClient).ReDiscoverPause,
			ZtpStateConnect:           (*ZtpClient).Connect,
			ZtpStateReConnectPause:    (*ZtpClient).ReConnectPause,
			ZtpStateUpgrade:           (*ZtpClient).Upgrade,
			ZtpStateReUpgradePause:    (*ZtpClient).ReUpgradePause,
			ZtpStateConfig:            (*ZtpClient).Config,
			ZtpStateReConfigPause:     (*ZtpClient).ReConfigPause,
			ZtpStatePostConfig:        (*ZtpClient).PostConfig,
			ZtpStateRePostConfigPause: (*ZtpClient).RePostConfigPause,
			ZtpStateConfigAck:         (*ZtpClient).ConfigAck,
			ZtpStateRunning:           (*ZtpClient).Running,
			ZtpStateReRunningPause:    (*ZtpClient).ReRunningPause,
			ZtpStateDone:              (*ZtpClient).Done,
		},
		httpClient: &http.Client{
			Timeout: time.Second * 60,
			Transport: &http.Transport{
                DisableKeepAlives: true,
                MaxIdleConnsPerHost: 1,
                MaxConnsPerHost: 1,
				TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			},
		},
	}

	return zc
}

type ZtpLookupEntry struct {
	Host            string
	Port            string
	AddClientPrefix bool
}

//DEBUG can be used during development to output log messages
var DEBUG = false

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	envSetup()
}

func envSetup() {
	if d, ok := os.LookupEnv(EnvDebug); ok {
		switch d {
		case "0":
			DEBUG = false
		case "1":
			DEBUG = true
		}
	}
}

func (zc *ZtpClient) SetDeviceID(devID string) {
	zc.devID = devID
}

func (zc *ZtpClient) SetDeviceType(devType string) {
	zc.devType = devType
}

func (zc *ZtpClient) SetRebootEvent(rebootEvent string) {
	zc.rebootEvent = rebootEvent
}

func (zc *ZtpClient) SetDiscoverRetry(retry int) {
	zc.discoverRetry = retry
}

func (zc *ZtpClient) SetConnectRetry(retry int) {
	zc.connectRetry = retry
}
func (zc *ZtpClient) SetUpgradeRetry(retry int) {
	zc.upgradeRetry = retry
}

func (zc *ZtpClient) SetConfigRetry(retry int) {
	zc.configRetry = retry
}

func (zc *ZtpClient) SetRunningRetry(retry int) {
	zc.runningRetry = retry
}

func (zc *ZtpClient) SetEventRetry(retry int) {
	zc.eventRetry = retry
}
func (zc *ZtpClient) SetLogin(login string) {
	zc.login = login
	//log.Println("login", zc.login)
}
func (zc *ZtpClient) SetPassword(password string) {
	zc.password = password
	//log.Println("password", zc.password)
}
func (zc *ZtpClient) SetRedirect(uri string) {
	if u, err := url.Parse(zc.urlPrefix); err == nil {
		u.Host = uri
		zc.urlPrefix = u.String()
	}
}

func (zc *ZtpClient) SendRequest(r *http.Request) (resp *http.Response, body *[]byte, err error) {
	r.SetBasicAuth(zc.login, zc.password)
	r.Header.Add("Content-type", "application/json")
	msg.DumpReqURL(r)
	//msg.DumpRequest(r)
    r.Close = true
	resp, err = zc.httpClient.Do(r)

	if err != nil {
        if DEBUG {
            log.Println(err)
        }
	} else {
		//msg.DumpResponse(resp)
		var i interface{}
        if DEBUG {
            log.Println("HTTP Response code", resp.StatusCode)
            json.NewDecoder(resp.Body).Decode(&i)
        }
        resp.Body.Close()
		b, err := json.Marshal(i)
		if err == nil {
            if DEBUG {
                log.Println(string(b))
            }
			body = &b
		}
	}
	return
}

func (zc *ZtpClient) StateMachine() {
	state := ZtpStateInit
	for {
		if state == ZtpStateDone {
			break
		}
		if DEBUG {
			log.Println("Current State", state)
		}
		f, ok := zc.fsm[state]
		if ok {
			state = f(zc)
			if DEBUG {
				log.Println("Next State", state)
			}
			continue
		}
		// if the state is unknow, break out of the for loop and return
		if DEBUG {
			log.Println("Unknown state", state)
		}
		break
	}
	if DEBUG {
		log.Println("End")
	}
}
