package ztpclient

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"os"
	"time"
	msg "ztp"
)

// Environment variable ztpclientDEBUG=1 sets the package DEBUG=true
const (
	EnvDebug            = "ztpclientDEBUG"
	DefaultRetrySeconds = 30
)

type ZtpClientState int

const (
	ZtpStateInit ZtpClientState = 1 + iota
	ZtpStateDiscover
	ZtpStateConnect
	ZtpStateUpgrade
	ZtpStateConfig
	ZtpStatePostConfig
	ZtpStateConfigAck
	ZtpStateRunning
	ZtpStateDone
)

const (
	REST_Version = "v1"
)

type ZtpClient struct {
	devID         string
	devType       string
	rebootEvent   string
	discoverRetry int                 // seconds before retrying discovery
	property      msg.ApPropertyBlock // this block gets used in every message

	upgradeAssets []*msg.Assets
	state         ZtpClientState // keep track of the state machine
	httpClient    *http.Client
	device        Device // the device specific interfaces
	urlPrefix     string // Once discovered the URL prefix is the same for all transactions
	fsm           map[ZtpClientState]func() ZtpClientState
}

type ZtpLookupEntry struct {
	host            string
	port            string
	addClientPrefix bool
}

//DEBUG can be used during development to output log messages
var DEBUG = false

var (
	//DNS controller names
	controllerList = []ZtpLookupEntry{
		ZtpLookupEntry{"extremecontrol", ":8443", true},
		ZtpLookupEntry{"extremecontrol.extremenetworks.com", ":8443", true},
		//ZtpLookupEntry{"devices.extremenetworks.com", ""},
	}
)

func init() {
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

// create a new ZTP client instance
func NewZtpClient(dev Device) *ZtpClient {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	zc := &ZtpClient{
		httpClient:    &http.Client{Timeout: time.Second * 20},
		state:         ZtpStateInit,
		device:        dev,
		discoverRetry: DefaultRetrySeconds,
	}

	zc.httpClient.Transport = &http.Transport{
		MaxIdleConnsPerHost: 10,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
	}

	zc.fsm = make(map[ZtpClientState]func() ZtpClientState)
	zc.fsm[ZtpStateInit] = zc.Init
	zc.fsm[ZtpStateDiscover] = zc.Discover
	zc.fsm[ZtpStateConnect] = zc.Connect
	zc.fsm[ZtpStateUpgrade] = zc.Upgrade
	zc.fsm[ZtpStateConfig] = zc.Config
	zc.fsm[ZtpStatePostConfig] = zc.PostConfig
	zc.fsm[ZtpStateConfigAck] = zc.ConfigAck
	zc.fsm[ZtpStateRunning] = zc.Running
	zc.fsm[ZtpStateDone] = zc.Done

	return zc
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

func (zc *ZtpClient) SetController(ctlr string) {
	var err error
	var controller ZtpLookupEntry
	if ctlr == "" {
		return
	}
	controller.host, controller.port, err = net.SplitHostPort(ctlr)
	if err != nil {
		log.Println(err)
		return
	}
	if controller.port != "" {
		controller.port = ":" + controller.port
	}
	// insert the debug controller value at the front of the list
	controller.addClientPrefix = true
	controllerList = append([]ZtpLookupEntry{controller}, controllerList...)
	controller.addClientPrefix = false
	controllerList = append([]ZtpLookupEntry{controller}, controllerList...)
}

func (zc *ZtpClient) SetDiscoverRetry(retry int) {
	zc.discoverRetry = DefaultRetrySeconds
}

func (zc *ZtpClient) AddUpgradeAsset(deviceAsset *msg.Assets) {
	zc.upgradeAssets = append(zc.upgradeAssets, deviceAsset)
}

func (zc *ZtpClient) StateMachine() {
	if DEBUG {
		log.Println("Begin")
	}
	for {
		if zc.state == ZtpStateDone {
			break
		}
		if DEBUG {
			log.Println("Current State", zc.state)
		}
		f, ok := zc.fsm[zc.state]
		if ok {
			zc.state = f()
			if DEBUG {
				log.Println("Next State", zc.state)
			}
			continue
		}
		// if the state is unknow, break out of the for loop and return
		if DEBUG {
			log.Println("Unknown state", zc.state)
		}
		break
	}
	if DEBUG {
		log.Println("End")
	}
}
