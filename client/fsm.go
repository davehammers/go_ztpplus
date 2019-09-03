package ztpclient

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"net"
	msg "ztp"
)

// Environment variable ztpclientDEBUG=1 sets the package DEBUG=true
const (
    EnvDebug = "ztpclientDEBUG"
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

type ZtpLookupEntry struct {
	host            string
	port            string
	addClientPrefix bool
}

type ZtpClient struct {
	devID          string
    devType        string
	rebootEvent    string
	controllerName string // FQDN or IP of fallback controller
	discoverRetry  int    // seconds before retrying discovery
	property      *msg.ApPropertyBlock
	upgradeAssets []*msg.Assets
	state      ZtpClientState // keep track of the state machine
	controller ZtpLookupEntry // after discovery, this is the controller IP
	httpClient *http.Client
	device     Device // the device specific interfaces
	fsm        map[ZtpClientState]func() ZtpClientState
	urlPrefix  string
}

//DEBUG can be used during development to output log messages
var DEBUG = false

var (
	//DNS lookup
	hostList = []ZtpLookupEntry{
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
		httpClient: &http.Client{Timeout: time.Second * 20},
		state:      ZtpStateInit,
		device:     dev,
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
    hostList = append([]ZtpLookupEntry{controller}, hostList...)
    controller.addClientPrefix = false
    hostList = append([]ZtpLookupEntry{controller}, hostList...)
}

func (zc *ZtpClient) SetDiscoverRetry(retry int) {
	zc.discoverRetry = DefaultRetrySeconds
}

func (zc *ZtpClient) SetPropertyBlock(property *msg.ApPropertyBlock) {
	zc.property = property
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

//This routine performs the 'INIT' state actions.
//
//Initializes the state machine and calls the application's INIT function
//callback.
//
//Adds the ztpclient to the upgrade asset list.
func (zc *ZtpClient) Init() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}

	zc.device.Init()
	return ZtpStateDiscover
}

//This routine performs the 'DISCOVER' state actions.
//
//Discovers the Extreme Control service by performing a DNS lookup of
//'extremecontrol.<local_domain>'.  The returned list of IP addresses
//are scanned one at a time, and the first one to successfully
//communicate with the Extreme Control service is used.
//
//If a local Extreme Control service is not found, the state machine
//tries to discover an Extreme Control service in the cloud via
//'devices.extremenetworks.com'.
//
//The application may configure a fallback controller if the discovery
//mechanism fails.  The fallback controller is configured by setting the
//self.data.args.controller_addr variable.
func (zc *ZtpClient) Discover() (state ZtpClientState) {
	if DEBUG {
		log.Println("Begin")
	}
	urlString := "https://%s%s/%sdevice/%s/"
	for _, hostPort := range hostList {
		prefix := ""
		if hostPort.addClientPrefix {
			prefix = "Client/"
		}
		url := fmt.Sprintf(urlString, hostPort.host, hostPort.port, prefix, REST_Version)
        if DEBUG {
            log.Println(url + "discovery")
        }
		resp, err := zc.httpClient.Get(url + "discovery")
		if err != nil {
            if DEBUG {
                log.Println(err)
            }
			continue
		}

		switch resp.StatusCode {
		case http.StatusOK, http.StatusNotFound:
			// discovery is successful
			zc.controller = hostPort
			zc.device.DiscoverOK(&zc.controller)
			// build up the URL prefix for all other message types
			zc.urlPrefix = fmt.Sprintf("%s%s/%s", url, zc.devType, zc.devID)
			if DEBUG {
				log.Println(zc.urlPrefix)
			}
			return ZtpStateConnect
		default:
			continue
		}
	}
	// discovery failed
	switch zc.device.DiscoverFail() {
	case DeviceReturnFinish:
		return ZtpStateDone
	}

	// pause here before trying again
	time.Sleep(time.Second * time.Duration(zc.discoverRetry))
	return ZtpStateDiscover
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
