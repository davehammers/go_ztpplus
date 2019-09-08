package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"log"
	msg "ztp"
	fsm "ztp/client/fsm"
)

//This interface defines the functions that every feature must implement
type Feature interface {
	getCapability(*msg.Capabilities) error
	getConnect(*msg.Connect) error
	getConfig(*msg.ConfigBlock) error
	setConfig(*msg.ConfigBlock) error
}

// The Device defines the information required by the FSM to manage a ZTP device instance
// Additional information may be added below the comment for a device specific implementtion.
type DeviceData struct {
	devID         string
	fsm           *fsm.ZtpClient
	controller    *fsm.ZtpLookupEntry
	property      *msg.ApPropertyBlock // this block gets used in every message
	capabilities  *msg.Capabilities
	upgradeAssets []msg.Asset
	events        []msg.Event

	// add device specific data elements here
}
type Device struct {
	data *DeviceData
}

var (
	des3Key []byte = []byte("99a3d014b62293332b85c3ab")
)

//Create a new device instance that implements the FSM Device interface.
//
//The devID is the unique device identifier string provided to the controller.
//The devID must be unique for all devices reporting to a single controller. Typically this is a serial number or a device MAC address.
func NewDevice(devID string) Device {
	dev := Device{
		data: &DeviceData{
			devID:         devID,
			property:      &msg.ApPropertyBlock{},
			capabilities:  &msg.Capabilities{},
			upgradeAssets: make([]msg.Asset, 0),
			events:        make([]msg.Event, 0),
			// new FSM for our device instance
			fsm: fsm.NewZtpClient(),
		},
	}
	dev.data.fsm.SetDeviceID(dev.data.devID)
	dev.data.fsm.Device = dev
	dev.data.fsm.SetDeviceType("switch")
	return dev
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
	dev.data.fsm.StateMachine()
}

func (dev Device) AddAsset(asset *msg.Asset) {
	msg.DumpJson(asset)
	dev.data.upgradeAssets = append(dev.data.upgradeAssets, *asset)
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
