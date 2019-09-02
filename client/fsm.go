package ztpclient

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"time"
	"ztp"
)

// Environment variable ztpclientDEBUG=1 sets the package DEBUG=true
const EnvDebug = "ztpclientDEBUG"

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

type ZtpDeviceReturnCode int

const (
	ZtpDeviceReturnOk ZtpDeviceReturnCode = 0 + iota
	ZtpDeviceReturnRestart
	ZtpDeviceReturnRetry
	ZtpDeviceReturnFinish
	ZtpDeviceReturnDone
)

type ZtpDevice interface {
	Connect(*ztp.Connect) ZtpDeviceReturnCode
	Upgrade(*ztp.ImageUpgrade) ZtpDeviceReturnCode
	Config(*ztp.Configuration) ZtpDeviceReturnCode
	PostConfig(msg *ztp.Configuration, err error, retCode int) ZtpDeviceReturnCode
	ConfigAcA(msg *ztp.Configuration, err error, retCode int) ZtpDeviceReturnCode
	Running() ZtpDeviceReturnCode
	Done()
}

type ZtpClient struct {
	clientID   string
	state      ZtpClientState
	httpClient *http.Client
	host       string
	port       string
	device     *ZtpDevice
	fsm        map[ZtpClientState]func() ZtpClientState
}

type ZtpLookupEntry struct {
	host string
	port string
}

//DEBUG can be used during development to output log messages
var DEBUG = false

var (
	//DNS lookup
	hostList = []ZtpLookupEntry{
		ZtpLookupEntry{"extremecontrol", "8443"},
		ZtpLookupEntry{"extremecontrol.extremenetworks.com", "8443"},
		//ZtpLookupEntry{"devices.extremenetworks.com", "8443"},
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
func NewZtpClient(clientID string) *ZtpClient {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	zc := &ZtpClient{
		clientID:   clientID,
		httpClient: &http.Client{Timeout: time.Second * 20},
		state:      ZtpStateInit,
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

func (zc *ZtpClient) SetDevice(dev *ZtpDevice) {
	zc.device = dev
}

// Override the discovery by setting the IP/Port to a known value
func (zc *ZtpClient) SetHost(host, port string) {
	if DEBUG {
		log.Println("host", host, "port", port)
	}
	zc.host = host
	zc.port = port
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

func (zc *ZtpClient) Init() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

    state = ZtpStateDiscover

	return state
}
func (zc *ZtpClient) Discover() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

    state = ZtpStateConnect

	return state
}
func (zc *ZtpClient) Connect() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

    state = ZtpStateUpgrade

	return state
}
func (zc *ZtpClient) Upgrade() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

    state = ZtpStateConfig

	return state
}
func (zc *ZtpClient) Config() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

    state = ZtpStatePostConfig

	return state
}
func (zc *ZtpClient) PostConfig() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

    state = ZtpStateConfigAck

	return state
}
func (zc *ZtpClient) ConfigAck() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin", zc.state)
	}

    state = ZtpStateRunning

	return state
}
func (zc *ZtpClient) Running() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

    state = ZtpStateDone

	return state
}
func (zc *ZtpClient) Done() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

    state = ZtpStateDone

	return state
}
