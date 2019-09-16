package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"log"
	msg "ztp"
	"ztp/client/device"
	fsm "ztp/client/fsm"
)

//The Device defines the information required by the FSM to manage a ZTP device instance
//Additional information may be added below the comment for a device specific implementtion.
//When adding additional fields to the Device struct, keep in mind this is a Go interface. The methods that implement the interface recieve a copy of this struct as it was created. To add data elements, they must be initialized pointers to some other memeory location. Updating the *pointer will update the provided memory location. Storing data directly in this struct after it has been created will have no effect.
type Device struct {
	devID      string               // device ID for this instance
	simulation bool                 // true if simulating multiple devices
	fsm        *fsm.ZtpClient       // HTTP client instance
	controller *fsm.ZtpLookupEntry  // created by device. controler DNS names
	property   *msg.ApPropertyBlock // this block Gets used in every message
	events     *[]msg.Event         // dynamic list of events during runtime
	features   *[]device.Feature    // each feature is added to this table
	// add device specific data elements here
}

var (
	des3Key []byte = []byte("99a3d014b62293332b85c3ab")
)

//Create a new device instance that implements the FSM Device interface.
//
//The devID is the unique device identifier string provided to the controller.
//The devID must be unique for all devices reporting to a single controller. Typically this is a serial number or a device MAC address.
func NewDevice(devID string, simulation bool) (i fsm.Device) {
	events := make([]msg.Event, 0)
	features := make([]device.Feature, 0)
	dev := Device{
		devID:      devID,
		simulation: simulation,
		property:   &msg.ApPropertyBlock{},
		events:     &events,
		features:   &features,
		// new FSM for our device instance
		fsm: fsm.NewZtpClient(),
	}
	dev.fsm.SetDeviceID(dev.devID)
	dev.fsm.Device = dev
	dev.fsm.SetDeviceType("switch")

	// These are the features implemented by the device
	*dev.features = append(*dev.features, NewDevFeature())
	*dev.features = append(*dev.features, NewDevConfigDownload())
	*dev.features = append(*dev.features, NewDevDot1X())
	*dev.features = append(*dev.features, NewDevLacp())
	*dev.features = append(*dev.features, NewDevLicense())
	*dev.features = append(*dev.features, NewDevLldp())
	*dev.features = append(*dev.features, NewDevLogins())
	*dev.features = append(*dev.features, NewDevMacAuth())
	*dev.features = append(*dev.features, NewDevMlag())
	*dev.features = append(*dev.features, NewDevMgmtAccess())
	*dev.features = append(*dev.features, NewDevMvrp())
	*dev.features = append(*dev.features, NewDevPoe())
	*dev.features = append(*dev.features, NewDevPorts())
	*dev.features = append(*dev.features, NewDevRadiusServers())
	*dev.features = append(*dev.features, NewDevSnmp())
	*dev.features = append(*dev.features, NewDevSpanningTree())
	*dev.features = append(*dev.features, NewDevSyslog())
	*dev.features = append(*dev.features, NewDevStacking())
	*dev.features = append(*dev.features, NewDevTelemetry())
	*dev.features = append(*dev.features, NewDevVlans())
	*dev.features = append(*dev.features, NewDevVpex())
	*dev.features = append(*dev.features, NewDevEee())
	*dev.features = append(*dev.features, NewDevLocatorLed())
	*dev.features = append(*dev.features, NewDevMlagv2())

	i = dev
	return
}

//This function allows the application to override the default
//routine for retrieving the source IP address of the socket used
//to communicate with Extreme Control.
//
//The override routine SHALL take two parameters, an IP address and
//a port number.
func (dev Device) GetSourceIP() {
}

func (dev Device) StartFSM() {
	// this function doesn't return until the FSM is completed
	dev.fsm.StateMachine()
}

func (dev Device) AddAsset(asset *msg.Asset) {
	msg.DumpJson(asset)
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
