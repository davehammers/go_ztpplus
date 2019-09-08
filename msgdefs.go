package ztp

type ApPropertyBlock struct {
    RuSerialNumber string `json:"ruSerialNumber"`
    BpWiredMacaddr string `json:"bpWiredMacaddr"`
    RuSwVersion    string `json:"ruSwVersion"`
    RuModel        string `json:"ruModel"`
}

type Capabilities struct {
    ConfigDownload struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"configDownload"`
    Dot1X struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"dot1x"`
    Lacp struct {
        FeatureAvailable bool `json:"featureAvailable"`
        MaxLagCount      int  `json:"maxLagCount"`
    } `json:"lacp"`
    License struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"license"`
    Lldp struct {
        FeatureAvailable bool  `json:"featureAvailable"`
        StandardTlvs     []int `json:"standardTlvs"`
        CustomTlvs       []struct {
            Oui     string `json:"oui"`
            Subtype int    `json:"subtype"`
        } `json:"customTlvs"`
    } `json:"lldp"`
    Logins struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"logins"`
    MacAuth struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"macAuth"`
    Mlag struct {
        FeatureAvailable bool `json:"featureAvailable"`
        MaxMlagCount     int  `json:"maxMlagCount"`
    } `json:"mlag"`
    MgmtAccess struct {
        Telnet struct {
            FeatureAvailable bool `json:"featureAvailable"`
        } `json:"telnet"`
        SSH struct {
            FeatureAvailable bool `json:"featureAvailable"`
        } `json:"ssh"`
        HTTP struct {
            FeatureAvailable bool `json:"featureAvailable"`
        } `json:"http"`
        HTTPS struct {
            FeatureAvailable bool `json:"featureAvailable"`
        } `json:"https"`
    } `json:"mgmtAccess"`
    Mvrp struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"mvrp"`
    Poe struct {
        FeatureAvailable bool `json:"featureAvailable"`
        PowerBudget      int  `json:"powerBudget"`
    } `json:"poe"`
    Ports struct {
        FeatureAvailable bool `json:"featureAvailable"`
        PortCapabilities []struct {
            Ports           []string `json:"ports"`
            PortSpeed       string   `json:"portSpeed"`
            PortDuplex      string   `json:"portDuplex"`
            AutoNegotiation bool     `json:"autoNegotiation"`
        } `json:"portCapabilities"`
        DesiredPortSettings []struct {
            Ports           []string `json:"ports"`
            PortSpeed       string   `json:"portSpeed"`
            PortDuplex      string   `json:"portDuplex"`
            AutoNegotiation bool     `json:"autoNegotiation"`
        } `json:"desiredPortSettings"`
    } `json:"ports"`
    RadiusServers struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"radiusServers"`
    Snmp struct {
        SnmpV1 struct {
            FeatureAvailable bool `json:"featureAvailable"`
        } `json:"snmpV1"`
        SnmpV2C struct {
            FeatureAvailable bool `json:"featureAvailable"`
        } `json:"snmpV2c"`
        SnmpV3 struct {
            FeatureAvailable bool `json:"featureAvailable"`
        } `json:"snmpV3"`
    } `json:"snmp"`
    SpanningTree struct {
        FeatureAvailable bool `json:"featureAvailable"`
        SpanGuard        bool `json:"spanGuard"`
        LoopProtect      bool `json:"loopProtect"`
    } `json:"spanningTree"`
    Syslog struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"syslog"`
    Stacking struct {
        FeatureAvailable bool     `json:"featureAvailable"`
        Ports            []string `json:"ports"`
    } `json:"stacking"`
    Telemetry struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"telemetry"`
    Vlans struct {
        FeatureAvailable bool `json:"featureAvailable"`
        MaxNtpServers    int  `json:"maxNtpServers"`
        MaxDNSServers    int  `json:"maxDnsServers"`
    } `json:"vlans"`
    Vxlan struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"vxlan"`
    Vpex struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"vpex"`
    Eee struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"eee"`
    LocatorLed struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"locatorLed"`
    Mlagv2 struct {
        FeatureAvailable bool `json:"featureAvailable"`
    } `json:"mlagv2"`
}
type FeaturePack   struct {
    FeaturePackName    string `json:"featurePackName"`
    FeaturePackLicense string `json:"featurePackLicense"`
}
type ConfigBlock struct {
    License struct {
        SystemLicense  string `json:"systemLicense"`
        EffectiveLevel string `json:"effectiveLevel"`
        FeaturePacks   []FeaturePack `json:"featurePacks"`
    } `json:"license"`
    Poe struct {
        Enabled              bool   `json:"enabled"`
        Logging              string `json:"logging"`
        DisconnectPrecedence string `json:"disconnectPrecedence"`
        Ports                []struct {
            Enabled       bool   `json:"enabled"`
            PortName      string `json:"portName"`
            PoePortAlias  string `json:"poePortAlias"`
            Detection     string `json:"detection"`
            Priority      string `json:"priority"`
            OperatorLimit string `json:"operatorLimit"`
        } `json:"ports"`
    } `json:"poe"`
    Dot1X struct {
        Enabled      bool   `json:"enabled"`
        Logging      string `json:"logging"`
        SharedSecret string `json:"sharedSecret"`
        RadiusServer string `json:"radiusServer"`
        Dot1XConfig  []struct {
            Dot1XEnable bool     `json:"dot1XEnable"`
            Dot1XPorts  []string `json:"dot1XPorts"`
        } `json:"dot1XConfig"`
    } `json:"dot1X"`
    Lacp struct {
        Enabled bool     `json:"enabled"`
        Logging string   `json:"logging"`
        Lags    []string `json:"lags"`
    } `json:"lacp"`
    Lldp struct {
        Enabled        bool   `json:"enabled"`
        Logging        string `json:"logging"`
        MgmtAddress    string `json:"mgmtAddress"`
        Location       string `json:"location"`
        SystemName     string `json:"systemName"`
        TlvMgmtAddress bool   `json:"tlvMgmtAddress"`
        TlvSystemName  bool   `json:"tlvSystemName"`
        TlvLocation    bool   `json:"tlvLocation"`
        LldpConfig     []struct {
            LldpEnable bool     `json:"lldpEnable"`
            LldpPorts  []string `json:"lldpPorts"`
        } `json:"lldpConfig"`
    } `json:"lldp"`
    Logins struct {
        LoginConfig []struct {
            AccessLevel string `json:"accessLevel"`
            Password    string `json:"password"`
            Username    string `json:"username"`
        } `json:"loginConfig"`
    } `json:"logins"`
    MacAuth struct {
        Enabled       string `json:"enabled"`
        MacAuthConfig []struct {
            MacAuthEnable bool   `json:"macAuthEnable"`
            MacAuthPorts  string `json:"macAuthPorts"`
        } `json:"macAuthConfig"`
    } `json:"macAuth"`
    Ports []struct {
        PortName         string   `json:"portName"`
        PortAlias        string   `json:"portAlias"`
        AdminStatus      string   `json:"adminStatus"`
        Pvid             int      `json:"pvid"`
        NodeAlias        bool     `json:"nodeAlias"`
        Mvrp             bool     `json:"mvrp"`
        AdminRules       string   `json:"adminRules"`
        IngressFiltering string   `json:"ingressFiltering"`
        LinkDuplex       string   `json:"linkDuplex"`
        LinkSpeed        string   `json:"linkSpeed"`
        LinkDuplexList   []string `json:"linkDuplexList"`
        LinkSpeedList    []string `json:"linkSpeedList"`
    } `json:"ports"`
    Snmp struct {
        SysName       string   `json:"sysName"`
        SysLocation   string   `json:"sysLocation"`
        SysContact    string   `json:"sysContact"`
        Notifications []string `json:"notifications"`
        SnmpV1        struct {
            Enabled     bool     `json:"enabled"`
            Logging     string   `json:"logging"`
            Communities []string `json:"communities"`
        } `json:"snmpV1"`
        SnmpV3 struct {
            Enabled bool     `json:"enabled"`
            Logging string   `json:"logging"`
            Users   []string `json:"users"`
        } `json:"snmpV3"`
        SnmpV2C struct {
            Enabled     bool     `json:"enabled"`
            Logging     string   `json:"logging"`
            Communities []string `json:"communities"`
        } `json:"snmpV2c"`
    } `json:"snmp"`
    SpanningTree struct {
        Enabled        bool   `json:"enabled"`
        Logging        string `json:"logging"`
        BridgePriority int    `json:"bridgePriority"`
        Version        string `json:"version"`
        MstiRegionName string `json:"mstiRegionName"`
        PortLinkType   []struct {
            LinkType      string   `json:"linkType"`
            LinkTypePorts []string `json:"linkTypePorts"`
        } `json:"portLinkType"`
        SpanGuard []struct {
            SpanGuardEnabled bool     `json:"spanGuardEnabled"`
            SpanGuardPorts   []string `json:"spanGuardPorts"`
        } `json:"spanGuard"`
        LoopProtect []struct {
            LoopProtectEnabled bool     `json:"loopProtectEnabled"`
            LoopProtectPorts   []string `json:"loopProtectPorts"`
        } `json:"loopProtect"`
    } `json:"spanningTree"`
    Syslog struct {
        Enabled      bool `json:"enabled"`
        SyslogConfig []struct {
            ServerNetworkAddress string `json:"serverNetworkAddress"`
            ServerUDPPort        int    `json:"serverUdpPort"`
        } `json:"syslogConfig"`
    } `json:"syslog"`
    Vlans struct {
        VlanConfig []struct {
            Name                string   `json:"name"`
            NetworkAddress      string   `json:"networkAddress"`
            NetworkPrefixLength int      `json:"networkPrefixLength"`
            TaggedEgressPorts   string   `json:"taggedEgressPorts"`
            UntaggedEgressPorts []string `json:"untaggedEgressPorts"`
            VlanIds             string   `json:"vlanIds"`
        } `json:"vlanConfig"`
    } `json:"vlans"`
    MgmtAccess struct {
        TelnetConfig struct {
            TelnetEnabled bool `json:"telnetEnabled"`
        } `json:"telnetConfig"`
        SSHConfig struct {
            SSHEnabled bool `json:"sshEnabled"`
        } `json:"sshConfig"`
        HTTPConfig struct {
            HTTPEnabled bool `json:"httpEnabled"`
        } `json:"httpConfig"`
        HTTPSConfig struct {
            HTTPSEnabled bool `json:"httpsEnabled"`
        } `json:"httpsConfig"`
    } `json:"mgmtAccess"`
    Appliance struct {
        PurviewConfig struct {
            NisEnabled        bool   `json:"nisEnabled"`
            NisServerAddress  string `json:"nisServerAddress"`
            NisDomainName     string `json:"nisDomainName"`
            GreEnabled        bool   `json:"greEnabled"`
            TapModeEnabled    bool   `json:"tapModeEnabled"`
            TunnelModeEnabled bool   `json:"tunnelModeEnabled"`
            GreConfig         []struct {
                ApplianceNetworkAddress string `json:"applianceNetworkAddress"`
                CoreFlow2NetworkAddress string `json:"coreFlow2NetworkAddress"`
            } `json:"greConfig"`
            TapModeConfig []struct {
                TapInterface string `json:"tapInterface"`
            } `json:"tapModeConfig"`
            TunnelModeConfig []struct {
                TunnelInterface           string `json:"tunnelInterface"`
                TunnelNetworkAddress      string `json:"tunnelNetworkAddress"`
                TunnelNetworkGateway      string `json:"tunnelNetworkGateway"`
                TunnelNetworkPrefixLength int    `json:"tunnelNetworkPrefixLength"`
            } `json:"tunnelModeConfig"`
        } `json:"purviewConfig"`
        NacConfig struct {
            ApplianceGroup string `json:"applianceGroup"`
        } `json:"nacConfig"`
    } `json:"appliance"`
    MLag struct {
        Enabled    bool `json:"enabled"`
        MlagGroups []struct {
            MlagDetect bool     `json:"mlagDetect"`
            MlagPeer   string   `json:"mlagPeer"`
            MlagPorts  []string `json:"mlagPorts"`
        } `json:"mlagGroups"`
    } `json:"mLag"`
    MLagv2 struct {
        Logging   string `json:"logging"`
        MlagPeers []struct {
            MlagPeerID     string `json:"mlagPeerId"`
            MlagPeerIPAddr string `json:"mlagPeerIpAddr"`
            MlagpeerVr     string `json:"mlagpeerVr"`
            MlagPeerVlanID int    `json:"mlagPeerVlanId"`
        } `json:"mlagPeers"`
        MlagMembers []struct {
            MlagPort       string `json:"mlagPort"`
            MlagPortMlagID int    `json:"mlagPortMlagId"`
            MlagPortPeerID string `json:"mlagPortPeerId"`
        } `json:"mlagMembers"`
    } `json:"mLagv2"`
    Stack struct {
        Enabled  bool     `json:"enabled"`
        Logging  string   `json:"logging"`
        Topology string   `json:"topology"`
        Ports    []string `json:"ports"`
        Peers    struct {
            StackPeer string `json:"stackPeer"`
        } `json:"peers"`
    } `json:"stack"`
    Mvrp struct {
        Enabled    string `json:"enabled"`
        MvrpConfig []struct {
            MvrpEnable bool     `json:"mvrpEnable"`
            MvrpPorts  []string `json:"mvrpPorts"`
        } `json:"mvrpConfig"`
    } `json:"mvrp"`
    Ospf struct {
        Enabled bool   `json:"enabled"`
        Logging string `json:"logging"`
    } `json:"ospf"`
    Timestamp   string `json:"timestamp"`
    BpRequestID int    `json:"bpRequestId"`
    Status      string `json:"status"`
}
type Event struct {
    Type        int    `json:"type"`
    Severity    string `json:"severity"`
    Timestamp   string `json:"timestamp"`
    Target      string `json:"target"`
    Description string `json:"description"`
}
type Asset struct {
    AssetName    string `json:"assetName"`
    AssetVersion string `json:"assetVersion"`
    AssetType    string `json:"assetType"`
}

// C O N N E C T
type Connect struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock"`
	DeviceInfo struct {
		SysDescr        string `json:"sysDescr"`
		SysUpTime       int    `json:"sysUpTime"`
		SysContact      string `json:"sysContact"`
		SysName         string `json:"sysName"`
		SysObjectID     string `json:"sysObjectID"`
		OperatingSystem string `json:"operatingSystem"`
		SerialNumber    string `json:"serialNumber"`
		MacAddr         string `json:"macAddr"`
		MgmtIPAddr      string `json:"mgmtIpAddr"`
		MgmtPort        string `json:"mgmtPort"`
		License         struct {
			FeaturePacks        []FeaturePack `json:"featurePacks"`
			PortCapacityLicense string   `json:"portCapacityLicense"`
			EffectiveLicense    string   `json:"effectiveLicense"`
			EnabledLicense      string   `json:"enabledLicense"`
			SystemLicense       string   `json:"systemLicense"`
			EffectiveLevel      string   `json:"effectiveLevel"`
		} `json:"license"`
		Ports []struct {
			PortType  string   `json:"portType"`
			PortSpeed int      `json:"portSpeed"`
			PortList  []string `json:"portList"`
		} `json:"ports"`
		IfTable []struct {
			IfIndex           int    `json:"ifIndex"`
			IfName            string `json:"ifName"`
			IfDescr           string `json:"ifDescr"`
			IfType            int    `json:"ifType"`
			IfOutDiscards     int    `json:"ifOutDiscards"`
			IfAdminStatus     int    `json:"ifAdminStatus"`
			IfMtu             int    `json:"ifMtu"`
			IfInDiscards      int    `json:"ifInDiscards"`
			IfInErrors        int    `json:"ifInErrors"`
			IfOperStatus      int    `json:"ifOperStatus"`
			IfOutUcastPkts    int    `json:"ifOutUcastPkts"`
			IfInOctets        int    `json:"ifInOctets"`
			IfLastChange      int    `json:"ifLastChange"`
			IfPhysAddress     string `json:"ifPhysAddress"`
			IfInUcastPkts     int    `json:"ifInUcastPkts"`
			IfSpecific        string `json:"ifSpecific"`
			IfOutNUcastPkts   int    `json:"ifOutNUcastPkts"`
			IfOutQLen         int    `json:"ifOutQLen"`
			IfSpeed           int    `json:"ifSpeed"`
			IfOutOctets       int    `json:"ifOutOctets"`
			IfInUnknownProtos int    `json:"ifInUnknownProtos"`
			IfInNUcastPkts    int    `json:"ifInNUcastPkts"`
			IfOutErrors       int    `json:"ifOutErrors"`
		} `json:"ifTable"`
		PortsInfo []struct {
			PortName string `json:"portName"`
			Lldp     struct {
				PortID             string `json:"portId"`
				SystemName         string `json:"systemName"`
				MgmtAddress        string `json:"mgmtAddress"`
				ChassisID          string `json:"chassisId"`
				SystemCapabilities string `json:"systemCapabilities"`
			} `json:"lldp"`
		} `json:"portsInfo"`
	} `json:"deviceInfo"`
}

type ConnectResponse struct {
	Action        string `json:"action,omitempty"`
	StandyTimeout int    `json:"standyTimeout,omitempty"`
	Credentials   struct {
		Login    string `json:"login,omitempty"`
		Password string `json:"password,omitempty"`
	} `json:"credentials,omitempty"`
	Redirect struct {
		Type string `json:"type,omitempty"`
		URI  string `json:"uri,omitempty"`
	} `json:"redirect,omitempty"`
}

// I M A G E U P G R A D E
type ImageUpgrade struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock"`
	Assets []Asset `json:"assets"`
}

type ImageUpgradeResponse struct {
	ImageUpgradeBlock []struct {
		Upgrade      bool   `json:"upgrade,omitempty"`
		URI          string `json:"uri,omitempty"`
		Timeout      int    `json:"timeout,omitempty"`
		AssetName    string `json:"assetName,omitempty"`
		AssetVersion string `json:"assetVersion,omitempty"`
		AssetType    string `json:"assetType,omitempty"`
		AssetSize    int    `json:"assetSize,omitempty"`
	} `json:"imageUpgradeBlock,omitempty"`
}

// C O N F I G U R A T I O N
type Configuration struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock"`
	Capabilities Capabilities `json:"capabilities"`
	ConfigBlock ConfigBlock `json:"configBlock"`
}

// C O N F I G U R A T I O N   R E S P O N S E
type ConfigurationResponse struct {
	CliBlock    string `json:"cliBlock"`
	ConfigBlock ConfigBlock `json:"configBlock"`
}

// E V E N T
type EventMsg struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock"`
	Event []Event `json:"event"`
	ConfigACK ConfigBlock `json:"ConfigACK"`
}

// S T A T S
type Stats struct {
	ApPropertyBlock     ApPropertyBlock     `json:"apPropertyBlock,omitempty"`
	Capabilities        Capabilities        `json:"capabilities,omitempty"`
	Assets []Asset `json:"assets"`
	UpgradeAssets []Asset `json:"upgradeAssets"`
	ConfigBlock         ConfigBlock         `json:"configBlock,omitempty"`
	Cooling struct {
		Trays []struct {
			State string `json:"state"`
			Name  string `json:"name"`
			Fans  []struct {
				Name  string `json:"name"`
				State string `json:"state"`
				Rpm   int    `json:"rpm"`
			} `json:"fans"`
		} `json:"trays"`
	} `json:"cooling"`
	IfTable []struct {
		IfIndex           int    `json:"ifIndex"`
		IfName            string `json:"ifName"`
		IfDescr           string `json:"ifDescr"`
		IfType            int    `json:"ifType"`
		IfOutDiscards     int    `json:"ifOutDiscards"`
		IfAdminStatus     int    `json:"ifAdminStatus"`
		IfMtu             int    `json:"ifMtu"`
		IfInDiscards      int    `json:"ifInDiscards"`
		IfInErrors        int    `json:"ifInErrors"`
		IfOperStatus      int    `json:"ifOperStatus"`
		IfOutUcastPkts    int    `json:"ifOutUcastPkts"`
		IfInOctets        int    `json:"ifInOctets"`
		IfLastChange      int    `json:"ifLastChange"`
		IfPhysAddress     string `json:"ifPhysAddress"`
		IfInUcastPkts     int    `json:"ifInUcastPkts"`
		IfSpecific        string `json:"ifSpecific"`
		IfOutNUcastPkts   int    `json:"ifOutNUcastPkts"`
		IfOutQLen         int    `json:"ifOutQLen"`
		IfSpeed           int    `json:"ifSpeed"`
		IfOutOctets       int    `json:"ifOutOctets"`
		IfInUnknownProtos int    `json:"ifInUnknownProtos"`
		IfInNUcastPkts    int    `json:"ifInNUcastPkts"`
		IfOutErrors       int    `json:"ifOutErrors"`
	} `json:"ifTable"`
	PortsInfo []struct {
		PortName   string `json:"portName"`
		StatsCheck int    `json:"statsCheck"`
		StpStatus  struct {
			StpState string `json:"stpState"`
			StpRole  string `json:"stpRole"`
		} `json:"stpStatus"`
		Poe struct {
			Curr    int    `json:"curr"`
			Power   int    `json:"power"`
			Volts   int    `json:"volts"`
			State   string `json:"state"`
			Fault   string `json:"fault"`
			PdClass string `json:"pdClass"`
		} `json:"poe"`
		Lldp struct {
			PortID             string   `json:"portId"`
			SystemName         string   `json:"systemName"`
			SystemCapabilities []string `json:"systemCapabilities"`
			MgmtAddress        string   `json:"mgmtAddress"`
			ChassisID          string   `json:"chassisId"`
			Med                struct {
				ModelName    string   `json:"modelName"`
				SwVersion    string   `json:"swVersion"`
				Capabilities []string `json:"capabilities"`
				Serial       string   `json:"serial"`
				HwVersion    string   `json:"hwVersion"`
				MdiPdValue   string   `json:"mdiPdValue"`
			} `json:"med"`
		} `json:"lldp"`
		Lag struct {
			Enabled   bool   `json:"enabled"`
			Lag       string `json:"lag"`
			RxState   string `json:"rxState"`
			SelLogic  string `json:"selLogic"`
			MuxState  string `json:"muxState"`
			RxPackets int    `json:"rxPackets"`
			TxPackets int    `json:"txPackets"`
		} `json:"lag"`
		Link struct {
			AdminStatus      string `json:"adminStatus"`
			LinkStatus       string `json:"linkStatus"`
			LinkDuplexConfig string `json:"linkDuplexConfig"`
			LinkDuplexActual string `json:"linkDuplexActual"`
			LinkSpeedActual  string `json:"linkSpeedActual"`
			LinkSpeedConfig  string `json:"linkSpeedConfig"`
			RxBytes          int    `json:"rxBytes"`
			RxPackets        int    `json:"rxPackets"`
			RxErrors         int    `json:"rxErrors"`
			TxBytes          int    `json:"txBytes"`
			TxPackets        int    `json:"txPackets"`
			TxErrors         int    `json:"txErrors"`
		} `json:"link"`
	} `json:"portsInfo"`
	PowerSupply []struct {
		Module   string `json:"module"`
		State    string `json:"state"`
		PartInfo string `json:"partInfo"`
		Input    int    `json:"input"`
	} `json:"powerSupply"`
	Temperatue []struct {
		Temp   int    `json:"temp"`
		Status string `json:"status"`
		Name   string `json:"name"`
	} `json:"temperatue"`
	Telemetry []struct {
		Path           string `json:"path"`
		HTTPStatusCode int    `json:"httpStatusCode"`
		Repsonse       struct {
		} `json:"repsonse"`
	} `json:"telemetry"`
	IetfVxlanVxlan struct {
		GlobalEnable  bool `json:"global-enable"`
		VxlanInstance []struct {
			VxlanID           int    `json:"vxlan-id"`
			Description       string `json:"description"`
			UnknowUnicastDrop string `json:"unknow-unicast-drop"`
			FilterVrrp        string `json:"filter-vrrp"`
			VxlanEvpn         struct {
				RouteDistinguisher string `json:"route-distinguisher"`
				VpnTargets         []struct {
					RtValue string `json:"rt-value"`
					RtType  string `json:"rt-type"`
				} `json:"vpn-targets"`
			} `json:"vxlan-evpn"`
		} `json:"vxlan-instance"`
		VtepInstances []struct {
			VtepID                int    `json:"vtep-id"`
			VtepName              string `json:"vtep-name"`
			SourceInterface       string `json:"source-interface"`
			MulticastIP           string `json:"multicast-ip"`
			InnerVlanHandlingMode int    `json:"inner-vlan-handling-mode"`
			BindVxlanID           []int  `json:"bind-vxlan-id"`
			StaticVxlanTunnel     []struct {
				VxlanTunnelID   int    `json:"vxlan-tunnel-id"`
				VxlanTunnelName string `json:"vxlan-tunnel-name"`
				AddressFamily   []struct {
					Af                  string `json:"af"`
					TunnelSourceIP      string `json:"tunnel-source-ip"`
					TunnelDestinationIP string `json:"tunnel-destination-ip"`
					BindVxlanID         []int  `json:"bind-vxlan-id"`
				} `json:"address-family"`
			} `json:"static-vxlan-tunnel"`
			RedundancyGroupBind []struct {
				VxlanID         int `json:"vxlan-id"`
				RedundancyGroup int `json:"redundancy-group"`
			} `json:"redundancy-group-bind"`
		} `json:"vtep-instances"`
	} `json:"ietf-vxlan:vxlan"`
	IetfVxlanVxlanState struct {
		VxlanTunnel []struct {
			LocalIP        string `json:"local-ip"`
			RemoteIP       string `json:"remote-ip"`
			StaticTunnelID int    `json:"static-tunnel-id"`
			EvpnTunnelID   int    `json:"evpn-tunnel-id"`
			Statistics     struct {
				InBytes            int `json:"in-bytes"`
				OutBytes           int `json:"out-bytes"`
				InPackets          int `json:"in-packets"`
				OutPackets         int `json:"out-packets"`
				TunnelVniStatistic []struct {
					VxlanID    int `json:"vxlan-id"`
					InBytes    int `json:"in-bytes"`
					OutBytes   int `json:"out-bytes"`
					InPackets  int `json:"in-packets"`
					OutPackets int `json:"out-packets"`
				} `json:"tunnel-vni-statistic"`
			} `json:"statistics"`
		} `json:"vxlan-tunnel"`
	} `json:"ietf-vxlan:vxlan-state"`
	CPUUtilization struct {
		CPUMonitorInterval         int `json:"cpuMonitorInterval"`
		CPUMonitorTotalUtilization int `json:"cpuMonitorTotalUtilization"`
		CPUMonitorTable            []struct {
			Module            string `json:"module"`
			ProcessName       string `json:"processName"`
			ProcessID         int    `json:"processId"`
			ProcessState      string `json:"processState"`
			Utilization5Secs  int    `json:"utilization5secs"`
			Utilization10Secs int    `json:"utilization10secs"`
			Utilization30Secs int    `json:"utilization30secs"`
			Utilization1Min   int    `json:"utilization1min"`
			Utilization5Mins  int    `json:"utilization5mins"`
			Utilization30Mins int    `json:"utilization30mins"`
			Utilization1Hour  int    `json:"utilization1hour"`
			MaxUtilization    int    `json:"maxUtilization"`
			UserTime          int    `json:"userTime"`
			SystemTime        int    `json:"systemTime"`
		} `json:"cpuMonitorTable"`
		CPUMonitorSystemTable []struct {
			Module            string `json:"module"`
			Utilization5Secs  int    `json:"utilization5secs"`
			Utilization10Secs int    `json:"utilization10secs"`
			Utilization30Secs int    `json:"utilization30secs"`
			Utilization1Min   int    `json:"utilization1min"`
			Utilization5Mins  int    `json:"utilization5mins"`
			Utilization30Mins int    `json:"utilization30mins"`
			Utilization1Hour  int    `json:"utilization1hour"`
			MaxUtilization    int    `json:"maxUtilization"`
		} `json:"cpuMonitorSystemTable"`
	} `json:"cpuUtilization"`
	MemoryUtilization struct {
		MemoryMonitorTable []struct {
			Module      string `json:"module"`
			ProcessName string `json:"processName"`
			Usage       int    `json:"usage"`
			Limit       int    `json:"limit"`
		} `json:"memoryMonitorTable"`
		MemoryMonitorSystemTable []struct {
			Module      string `json:"module"`
			SystemTotal int    `json:"systemTotal"`
			SystemFree  int    `json:"systemFree"`
			SystemUsage int    `json:"systemUsage"`
			UserUsage   int    `json:"userUsage"`
		} `json:"memoryMonitorSystemTable"`
	} `json:"memoryUtilization"`
	Mlagv2 struct {
		MLagPeers []struct {
			MLagPeerID     string `json:"mLagPeerId"`
			MLagPeerStatus string `json:"mLagPeerStatus"`
			MLagPeerUpTime int    `json:"mLagPeerUpTime"`
		} `json:"mLagPeers"`
	} `json:"mlagv2"`
	Timestamp   int    `json:"timestamp"`
	SysUpTime   int    `json:"sysUpTime"`
	CheckInTime string `json:"checkInTime"`
	UpgradeTime string `json:"upgradeTime"`
}

type StatsResponse struct {
	StatsTimer            int                   `json:"statsTimer,omitempty"`
	StatsInterval         int                   `json:"statsInterval,omitempty"`
	//Action                Action                `json:"action,omitempty"`
	//Actions               []Actions             `json:"actions,omitempty"`
	Action                string                `json:"action,omitempty"`
	Actions               []string             `json:"actions,omitempty"`
	//ImageUpgradeBlock     ImageUpgradeBlock     `json:"imageUpgradeBlock,omitempty"`
	ImageUpgradeBlock     []ImageUpgradeBlock     `json:"imageUpgradeBlock,omitempty"`
	CliBlock              string                `json:"cliBlock,omitempty"`
	ConfigBlock           ConfigBlock           `json:"configBlock,omitempty"`
	TracesBlock           TracesBlock           `json:"tracesBlock,omitempty"`
	ConfigUpload          ConfigUpload          `json:"configUpload,omitempty"`
	ConfigReporting       ConfigReporting       `json:"configReporting,omitempty"`
	RebootDevices         RebootDevices         `json:"rebootDevices,omitempty"`
	ExosBpeReconciliation ExosBpeReconciliation `json:"exosBpeReconciliation,omitempty"`
}
type Action struct {
	Action     string `json:"action,omitempty"`
	ActionTime string `json:"actionTime,omitempty"`
}
type Actions struct {
	Action     string `json:"action,omitempty"`
	ActionTime string `json:"actionTime,omitempty"`
}
type ImageUpgradeBlock struct {
	Upgrade      bool   `json:"upgrade,omitempty"`
	URI          string `json:"uri,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	AssetName    string `json:"assetName,omitempty"`
	AssetVersion string `json:"assetVersion,omitempty"`
	AssetType    string `json:"assetType,omitempty"`
}
type TracesBlock struct {
	URI    string `json:"uri,omitempty"`
	Delete bool   `json:"delete,omitempty"`
}
type ConfigUpload struct {
	URI string `json:"uri,omitempty"`
}
type ConfigReporting struct {
	Enabled       bool     `json:"enabled,omitempty"`
	Monitorchange bool     `json:"monitorchange,omitempty"`
	Include       []string `json:"include,omitempty"`
	Exclude       []string `json:"exclude,omitempty"`
	Persistent    bool     `json:"persistent,omitempty"`
}
type RebootDevices struct {
	DeviceList []string `json:"deviceList,omitempty"`
}
type ReconfigureBpe struct {
	SerialNum  string `json:"serialNum,omitempty"`
	Slot       string `json:"slot,omitempty"`
	Port       string `json:"port,omitempty"`
	ModuleName string `json:"moduleName,omitempty"`
}
type RemoveLagPort struct {
	Port string `json:"port,omitempty"`
}
type ExosBpeReconciliation struct {
	Macro          string           `json:"macro,omitempty"`
	ReconfigureBpe []ReconfigureBpe `json:"reconfigureBpe,omitempty"`
	RemoveLagPort  []RemoveLagPort  `json:"removeLagPort,omitempty"`
}
