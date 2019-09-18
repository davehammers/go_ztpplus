package main

import (
	"sync"
	msg "ztp"
	"ztp/client/device"
)

//This is a feature definition template. The methods are required for all features and will be
//called by the device message controll logic to populate data structures or to update the device
//
//To create a new feature, copy this file and replace all references to "Lldp" with the name of the feature.
//
//e.g. s/Lldp/Vlan/
type devLldp struct{}

const lldp_lldpLocalSystemInfoC = "debug cfgmgr show next lldp.lldpLocalSystemInfo"

type lldp_lldpLocalSystemInfoT struct {
	Class string `json:"class"`
	Data  []struct {
		Module                  string `json:"_module"`
		LldpLocChassisID        string `json:"lldpLocChassisId"`
		LldpLocChassisIDSubtype string `json:"lldpLocChassisIdSubtype"`
		LldpLocSysCapEnabled    string `json:"lldpLocSysCapEnabled"`
		LldpLocSysCapSupported  string `json:"lldpLocSysCapSupported"`
		LldpLocSysDesc          string `json:"lldpLocSysDesc"`
		LldpLocSysName          string `json:"lldpLocSysName"`
		Message                 string `json:"message"`
		Status                  string `json:"status"`
	} `json:"data"`
	Module string `json:"module"`
}

var lldp_lldpLocalSystemInfo *lldp_lldpLocalSystemInfoT

const lldp_lldpLocMgmtAddrC = "debug cfgmgr show next lldp.lldpLocMgmtAddr lldpLocManAddr=None lldpLocManAddrSubtype=None"

type lldp_lldpLocMgmtAddrT struct {
	Class string `json:"class"`
	Data  []struct {
		Module                         string `json:"_module"`
		LldpConfigManAddrPortsTxEnable string `json:"lldpConfigManAddrPortsTxEnable"`
		LldpLocManAddr                 string `json:"lldpLocManAddr"`
		LldpLocManAddrIfID             string `json:"lldpLocManAddrIfId"`
		LldpLocManAddrIfSubtype        string `json:"lldpLocManAddrIfSubtype"`
		LldpLocManAddrLen              string `json:"lldpLocManAddrLen"`
		LldpLocManAddrOID              string `json:"lldpLocManAddrOID"`
		LldpLocManAddrSubtype          string `json:"lldpLocManAddrSubtype"`
		Message                        string `json:"message"`
		Status                         string `json:"status"`
	} `json:"data"`
	Module string `json:"module"`
}

var lldp_lldpLocMgmtAddr *lldp_lldpLocMgmtAddrT

const lldp_lldpPortConfigurationC = "debug cfgmgr show next lldp.lldpPortConfiguration port=None"

type lldp_lldpPortConfigurationT struct {
	Class string `json:"class"`
	Data  []struct {
		Module                          string      `json:"_module"`
		Action                          interface{} `json:"action"`
		AppID                           interface{} `json:"appID"`
		DcbxEnabled                     string      `json:"dcbxEnabled"`
		DllClassification               string      `json:"dllClassification"`
		FaAuthStatus                    string      `json:"faAuthStatus"`
		FaEncryptedKey                  string      `json:"faEncryptedKey"`
		FaKey                           interface{} `json:"faKey"`
		FaKeyInUse                      string      `json:"faKeyInUse"`
		LldpAdminStatus                 string      `json:"lldpAdminStatus"`
		LldpMedNotificationEnable       string      `json:"lldpMedNotificationEnable"`
		LldpNotificationEnable          string      `json:"lldpNotificationEnable"`
		LldpPortNumber                  string      `json:"lldpPortNumber"`
		LldpTLVsTxEnable                string      `json:"lldpTLVsTxEnable"`
		LldpXAvExLocPrioFramingMode     string      `json:"lldpXAvExLocPrioFramingMode"`
		LldpXAvExLocXPoEPSEPortReqLevel string      `json:"lldpXAvExLocXPoEPSEPortReqLevel"`
		LldpXAvExPortTLVEnabled         string      `json:"lldpXAvExPortTLVEnabled"`
		LldpXMedPortCapSupported        string      `json:"lldpXMedPortCapSupported"`
		LldpXMedPortConfigTLVsTxEnable  string      `json:"lldpXMedPortConfigTLVsTxEnable"`
		LldpXdot1ConfigPortVlanTxEnable string      `json:"lldpXdot1ConfigPortVlanTxEnable"`
		LldpXdot3TLVsTxEnable           string      `json:"lldpXdot3TLVsTxEnable"`
		Message                         string      `json:"message"`
		Port                            string      `json:"port"`
		PortList                        []string    `json:"portList"`
		Priority                        interface{} `json:"priority"`
		Status                          string      `json:"status"`
		TlvAll                          interface{} `json:"tlvAll"`
		TlvCallServer                   string      `json:"tlvCallServer"`
		TlvCallServerIP                 []string    `json:"tlvCallServerIP"`
		TlvDot1QFraming                 string      `json:"tlvDot1qFraming"`
		TlvDot1QFramingFormat           string      `json:"tlvDot1qFramingFormat"`
		TlvFileServer                   string      `json:"tlvFileServer"`
		TlvFileServerIP                 []string    `json:"tlvFileServerIP"`
		TlvLinkAggr                     string      `json:"tlvLinkAggr"`
		TlvMEDCapabilities              string      `json:"tlvMEDCapabilities"`
		TlvMEDLCI                       string      `json:"tlvMEDLCI"`
		TlvMEDPolicy                    string      `json:"tlvMEDPolicy"`
		TlvMEDPowerViaMDI               string      `json:"tlvMEDPowerViaMDI"`
		TlvMacPhyConfig                 string      `json:"tlvMacPhyConfig"`
		TlvMaxFrameSize                 string      `json:"tlvMaxFrameSize"`
		TlvMgmtAddr                     string      `json:"tlvMgmtAddr"`
		TlvPoEConservationReq           string      `json:"tlvPoEConservationReq"`
		TlvPortDesc                     string      `json:"tlvPortDesc"`
		TlvPortProtID                   string      `json:"tlvPortProtId"`
		TlvPortVlanID                   string      `json:"tlvPortVlanId"`
		TlvPowerMdi                     string      `json:"tlvPowerMdi"`
		TlvProtocolIdent                string      `json:"tlvProtocolIdent"`
		TlvSysCap                       string      `json:"tlvSysCap"`
		TlvSysDesc                      string      `json:"tlvSysDesc"`
		TlvSysName                      string      `json:"tlvSysName"`
		TlvVlanName                     string      `json:"tlvVlanName"`
		Type                            interface{} `json:"type"`
		Version                         interface{} `json:"version"`
	} `json:"data"`
	Module string `json:"module"`
}

var lldp_lldpPortConfiguration *lldp_lldpPortConfigurationT

const lldp_lldpLocalPortEntryC = "debug cfgmgr show next lldp.lldpLocalPortEntry"

type lldp_lldpLocalPortEntryT struct {
	Class string `json:"class"`
	Data  []struct {
		Module               string `json:"_module"`
		LldpLocPortDesc      string `json:"lldpLocPortDesc"`
		LldpLocPortID        string `json:"lldpLocPortId"`
		LldpLocPortIDSubtype string `json:"lldpLocPortIdSubtype"`
		LldpLocPortNum       string `json:"lldpLocPortNum"`
		Message              string `json:"message"`
		Status               string `json:"status"`
	} `json:"data"`
	Module string `json:"module"`
}

var lldp_lldpLocalPortEntry *lldp_lldpLocalPortEntryT
var lldp_portMap map[string]string

const lldp_lldpRemTableEntryC = "debug cfgmgr show next lldp.lldpRemTableEntry"

type lldp_lldpRemTableEntryT struct {
	Class string `json:"class"`
	Data  []struct {
		Module                  string `json:"_module"`
		LldpRemChassisID        string `json:"lldpRemChassisId"`
		LldpRemChassisIDSubtype string `json:"lldpRemChassisIdSubtype"`
		LldpRemIndex            string `json:"lldpRemIndex"`
		LldpRemLocalPortNum     string `json:"lldpRemLocalPortNum"`
		LldpRemPortDesc         string `json:"lldpRemPortDesc"`
		LldpRemPortID           string `json:"lldpRemPortId"`
		LldpRemPortIDSubtype    string `json:"lldpRemPortIdSubtype"`
		LldpRemSysCapEnabled    string `json:"lldpRemSysCapEnabled"`
		LldpRemSysCapSupported  string `json:"lldpRemSysCapSupported"`
		LldpRemSysDesc          string `json:"lldpRemSysDesc"`
		LldpRemSysName          string `json:"lldpRemSysName"`
		LldpRemTimeMark         string `json:"lldpRemTimeMark"`
		Message                 string `json:"message"`
		Status                  string `json:"status"`
	} `json:"data"`
	Module string `json:"module"`
}

var lldp_lldpRemTableEntry *lldp_lldpRemTableEntryT

//Return an instance of the feature interface
func NewDevLldp() (f device.Feature) {
	f = devLldp{}
	return
}

// feature collects and updates the device specific DB with config data
func (p devLldp) GetDBConfig() (err error) {
	var wg sync.WaitGroup
	cmds := map[string]interface{}{
		lldp_lldpLocalSystemInfoC:   &lldp_lldpLocalSystemInfo,
		lldp_lldpLocMgmtAddrC:       &lldp_lldpLocMgmtAddr,
		lldp_lldpPortConfigurationC: &lldp_lldpPortConfiguration,
		lldp_lldpLocalPortEntryC:    &lldp_lldpLocalPortEntry,
		lldp_lldpRemTableEntryC:     &lldp_lldpRemTableEntry,
	}

	for k, v := range cmds {
		wg.Add(1)
		go func(k string, v interface{}) {
			exosIO.Get(k, v)
			wg.Done()
		}(k, v)
	}
	wg.Wait()
	// data fixup
	lldp_portMap = make(map[string]string)
	for _, l := range lldp_lldpLocalPortEntry.Data {
		lldp_portMap[l.LldpLocPortNum] = msg.DecodeHexString(l.LldpLocPortID)
	}
	/*
		for _, l := range &lldp_lldpRemTableEntry.Data {
			l.LldpRemPortID = msg.DecodeHexString(l.LldpRemPortID)
			l.LldpRemSysDesc = msg.DecodeHexString(l.LldpRemSysDesc)
			l.LldpRemSysCapSupported = msg.DecodeHexString(l.LldpRemSysCapSupported)
			l.LldpRemSysName = msg.DecodeHexString(l.LldpRemSysName)
			l.LldpRemLocalPortNum = portMap[l.LldpRemLocalPortNum]
		}
	*/

	return
}

// feature collects and updates the device specific DB with config statistics data
func (p devLldp) GetDBStats() (err error) {
	return
}

//Update the feature capability in the Capabilities part of a message
//The feature should update any fields necessary to represent it's capabilities
func (p devLldp) GetConnect(m *msg.Connect) (err error) {
	//m.DeviceInfo.PortsInfo = make(msg.Connect.DeviceInfo.PortsInfo, 0)
	m.DeviceInfo.PortsInfo = make([]msg.PortsInfo, 0)
	for _, l := range lldp_lldpRemTableEntry.Data {
		x := msg.PortsInfo{
			PortName: lldp_portMap[l.LldpRemLocalPortNum],
			Lldp: msg.Lldp{
				PortID:     msg.DecodeHexString(l.LldpRemPortID),
				SystemName: msg.DecodeHexString(l.LldpRemSysName),
				//MgmtAddress:        l.NbrMgmtAddr,
				ChassisID:          l.LldpRemChassisID,
				SystemCapabilities: msg.DecodeHexString(l.LldpRemSysCapSupported),
			},
		}
		m.DeviceInfo.PortsInfo = append(m.DeviceInfo.PortsInfo, x)
	}
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devLldp) GetConfig(m *msg.Configuration) (err error) {
	m.Capabilities.Lldp.FeatureAvailable = true
	//m.ConfigBlock.Lldp.<somefield> = ""
	return
}

//update the feature information from the informatin received from the controller
func (p devLldp) SetConfig(m *msg.ConfigurationResponse) (err error) {
	return
}

//Update the feature informaiton in the Configuration message before it is sent to the controller
func (p devLldp) GetStats(m *msg.Stats) (err error) {
	//m.ConfigBlock.Lldp.<somefield> = ""
	return
}

//update the feature information from the information received from the controller
func (p devLldp) SetStats(m *msg.StatsResponse) (err error) {
	// create a config response to pass to the SetConfig function
	c := msg.ConfigurationResponse{}
	c.ConfigBlock = m.ConfigBlock
	err = p.SetConfig(&c)

	return
}
