package ztp

type ApPropertyBlock struct {
    RuSerialNumber string `json:"ruSerialNumber,omitempty"`
    BpWiredMacaddr string `json:"bpWiredMacaddr,omitempty"`
    RuSwVersion    string `json:"ruSwVersion,omitempty"`
    RuModel        string `json:"ruModel,omitempty"`
}

type Capabilities struct {
    ConfigDownload struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"configDownload,omitempty"`
    Dot1X struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"dot1x,omitempty"`
    Lacp struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
        MaxLagCount      int  `json:"maxLagCount,omitempty"`
    } `json:"lacp,omitempty"`
    License struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"license,omitempty"`
    Lldp struct {
        FeatureAvailable bool  `json:"featureAvailable,omitempty"`
        StandardTlvs     []int `json:"standardTlvs,omitempty"`
        CustomTlvs       []struct {
            Oui     string `json:"oui,omitempty"`
            Subtype int    `json:"subtype,omitempty"`
        } `json:"customTlvs,omitempty"`
    } `json:"lldp,omitempty"`
    Logins struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"logins,omitempty"`
    MacAuth struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"macAuth,omitempty"`
    Mlag struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
        MaxMlagCount     int  `json:"maxMlagCount,omitempty"`
    } `json:"mlag,omitempty"`
    MgmtAccess struct {
        Telnet struct {
            FeatureAvailable bool `json:"featureAvailable,omitempty"`
        } `json:"telnet,omitempty"`
        SSH struct {
            FeatureAvailable bool `json:"featureAvailable,omitempty"`
        } `json:"ssh,omitempty"`
        HTTP struct {
            FeatureAvailable bool `json:"featureAvailable,omitempty"`
        } `json:"http,omitempty"`
        HTTPS struct {
            FeatureAvailable bool `json:"featureAvailable,omitempty"`
        } `json:"https,omitempty"`
    } `json:"mgmtAccess,omitempty"`
    Mvrp struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"mvrp,omitempty"`
    Poe struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
        PowerBudget      int  `json:"powerBudget,omitempty"`
    } `json:"poe,omitempty"`
    Ports struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
        PortCapabilities []struct {
            Ports           []string `json:"ports,omitempty"`
            PortSpeed       string   `json:"portSpeed,omitempty"`
            PortDuplex      string   `json:"portDuplex,omitempty"`
            AutoNegotiation bool     `json:"autoNegotiation,omitempty"`
        } `json:"portCapabilities,omitempty"`
        DesiredPortSettings []struct {
            Ports           []string `json:"ports,omitempty"`
            PortSpeed       string   `json:"portSpeed,omitempty"`
            PortDuplex      string   `json:"portDuplex,omitempty"`
            AutoNegotiation bool     `json:"autoNegotiation,omitempty"`
        } `json:"desiredPortSettings,omitempty"`
    } `json:"ports,omitempty"`
    RadiusServers struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"radiusServers,omitempty"`
    Snmp struct {
        SnmpV1 struct {
            FeatureAvailable bool `json:"featureAvailable,omitempty"`
        } `json:"snmpV1,omitempty"`
        SnmpV2C struct {
            FeatureAvailable bool `json:"featureAvailable,omitempty"`
        } `json:"snmpV2c,omitempty"`
        SnmpV3 struct {
            FeatureAvailable bool `json:"featureAvailable,omitempty"`
        } `json:"snmpV3,omitempty"`
    } `json:"snmp,omitempty"`
    SpanningTree struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
        SpanGuard        bool `json:"spanGuard,omitempty"`
        LoopProtect      bool `json:"loopProtect,omitempty"`
    } `json:"spanningTree,omitempty"`
    Syslog struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"syslog,omitempty"`
    Stacking struct {
        FeatureAvailable bool     `json:"featureAvailable,omitempty"`
        Ports            []string `json:"ports,omitempty"`
    } `json:"stacking,omitempty"`
    Telemetry struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"telemetry,omitempty"`
    Vlans struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
        MaxNtpServers    int  `json:"maxNtpServers,omitempty"`
        MaxDNSServers    int  `json:"maxDnsServers,omitempty"`
    } `json:"vlans,omitempty"`
    Vxlan struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"vxlan,omitempty"`
    Vpex struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"vpex,omitempty"`
    Eee struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"eee,omitempty"`
    LocatorLed struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"locatorLed,omitempty"`
    Mlagv2 struct {
        FeatureAvailable bool `json:"featureAvailable,omitempty"`
    } `json:"mlagv2,omitempty"`
}
type FeaturePack   struct {
    FeaturePackName    string `json:"featurePackName,omitempty"`
    FeaturePackLicense string `json:"featurePackLicense,omitempty"`
}
type ConfigBlock struct {
    License struct {
        SystemLicense  string `json:"systemLicense,omitempty"`
        EffectiveLevel string `json:"effectiveLevel,omitempty"`
        FeaturePacks   []FeaturePack `json:"featurePacks,omitempty"`
    } `json:"license,omitempty"`
    Poe struct {
        Enabled              bool   `json:"enabled,omitempty"`
        Logging              string `json:"logging,omitempty"`
        DisconnectPrecedence string `json:"disconnectPrecedence,omitempty"`
        Ports                []struct {
            Enabled       bool   `json:"enabled,omitempty"`
            PortName      string `json:"portName,omitempty"`
            PoePortAlias  string `json:"poePortAlias,omitempty"`
            Detection     string `json:"detection,omitempty"`
            Priority      string `json:"priority,omitempty"`
            OperatorLimit string `json:"operatorLimit,omitempty"`
        } `json:"ports,omitempty"`
    } `json:"poe,omitempty"`
    Dot1X struct {
        Enabled      bool   `json:"enabled,omitempty"`
        Logging      string `json:"logging,omitempty"`
        SharedSecret string `json:"sharedSecret,omitempty"`
        RadiusServer string `json:"radiusServer,omitempty"`
        Dot1XConfig  []struct {
            Dot1XEnable bool     `json:"dot1XEnable,omitempty"`
            Dot1XPorts  []string `json:"dot1XPorts,omitempty"`
        } `json:"dot1XConfig,omitempty"`
    } `json:"dot1X,omitempty"`
    Lacp struct {
        Enabled bool     `json:"enabled,omitempty"`
        Logging string   `json:"logging,omitempty"`
        Lags    []string `json:"lags,omitempty"`
    } `json:"lacp,omitempty"`
    Lldp struct {
        Enabled        bool   `json:"enabled,omitempty"`
        Logging        string `json:"logging,omitempty"`
        MgmtAddress    string `json:"mgmtAddress,omitempty"`
        Location       string `json:"location,omitempty"`
        SystemName     string `json:"systemName,omitempty"`
        TlvMgmtAddress bool   `json:"tlvMgmtAddress,omitempty"`
        TlvSystemName  bool   `json:"tlvSystemName,omitempty"`
        TlvLocation    bool   `json:"tlvLocation,omitempty"`
        LldpConfig     []struct {
            LldpEnable bool     `json:"lldpEnable,omitempty"`
            LldpPorts  []string `json:"lldpPorts,omitempty"`
        } `json:"lldpConfig,omitempty"`
    } `json:"lldp,omitempty"`
    Logins struct {
        LoginConfig []struct {
            AccessLevel string `json:"accessLevel,omitempty"`
            Password    string `json:"password,omitempty"`
            Username    string `json:"username,omitempty"`
        } `json:"loginConfig,omitempty"`
    } `json:"logins,omitempty"`
    MacAuth struct {
        Enabled       string `json:"enabled,omitempty"`
        MacAuthConfig []struct {
            MacAuthEnable bool   `json:"macAuthEnable,omitempty"`
            MacAuthPorts  string `json:"macAuthPorts,omitempty"`
        } `json:"macAuthConfig,omitempty"`
    } `json:"macAuth,omitempty"`
    Ports []struct {
        PortName         string   `json:"portName,omitempty"`
        PortAlias        string   `json:"portAlias,omitempty"`
        AdminStatus      string   `json:"adminStatus,omitempty"`
        Pvid             int      `json:"pvid,omitempty"`
        NodeAlias        bool     `json:"nodeAlias,omitempty"`
        Mvrp             bool     `json:"mvrp,omitempty"`
        AdminRules       string   `json:"adminRules,omitempty"`
        IngressFiltering string   `json:"ingressFiltering,omitempty"`
        LinkDuplex       string   `json:"linkDuplex,omitempty"`
        LinkSpeed        string   `json:"linkSpeed,omitempty"`
        LinkDuplexList   []string `json:"linkDuplexList,omitempty"`
        LinkSpeedList    []string `json:"linkSpeedList,omitempty"`
    } `json:"ports,omitempty"`
    Snmp struct {
        SysName       string   `json:"sysName,omitempty"`
        SysLocation   string   `json:"sysLocation,omitempty"`
        SysContact    string   `json:"sysContact,omitempty"`
        Notifications []string `json:"notifications,omitempty"`
        SnmpV1        struct {
            Enabled     bool     `json:"enabled,omitempty"`
            Logging     string   `json:"logging,omitempty"`
            Communities []string `json:"communities,omitempty"`
        } `json:"snmpV1,omitempty"`
        SnmpV3 struct {
            Enabled bool     `json:"enabled,omitempty"`
            Logging string   `json:"logging,omitempty"`
            Users   []string `json:"users,omitempty"`
        } `json:"snmpV3,omitempty"`
        SnmpV2C struct {
            Enabled     bool     `json:"enabled,omitempty"`
            Logging     string   `json:"logging,omitempty"`
            Communities []string `json:"communities,omitempty"`
        } `json:"snmpV2c,omitempty"`
    } `json:"snmp,omitempty"`
    SpanningTree struct {
        Enabled        bool   `json:"enabled,omitempty"`
        Logging        string `json:"logging,omitempty"`
        BridgePriority int    `json:"bridgePriority,omitempty"`
        Version        string `json:"version,omitempty"`
        MstiRegionName string `json:"mstiRegionName,omitempty"`
        PortLinkType   []struct {
            LinkType      string   `json:"linkType,omitempty"`
            LinkTypePorts []string `json:"linkTypePorts,omitempty"`
        } `json:"portLinkType,omitempty"`
        SpanGuard []struct {
            SpanGuardEnabled bool     `json:"spanGuardEnabled,omitempty"`
            SpanGuardPorts   []string `json:"spanGuardPorts,omitempty"`
        } `json:"spanGuard,omitempty"`
        LoopProtect []struct {
            LoopProtectEnabled bool     `json:"loopProtectEnabled,omitempty"`
            LoopProtectPorts   []string `json:"loopProtectPorts,omitempty"`
        } `json:"loopProtect,omitempty"`
    } `json:"spanningTree,omitempty"`
    Syslog struct {
        Enabled      bool `json:"enabled,omitempty"`
        SyslogConfig []struct {
            ServerNetworkAddress string `json:"serverNetworkAddress,omitempty"`
            ServerUDPPort        int    `json:"serverUdpPort,omitempty"`
        } `json:"syslogConfig,omitempty"`
    } `json:"syslog,omitempty"`
    Vlans struct {
        VlanConfig []struct {
            Name                string   `json:"name,omitempty"`
            NetworkAddress      string   `json:"networkAddress,omitempty"`
            NetworkPrefixLength int      `json:"networkPrefixLength,omitempty"`
            TaggedEgressPorts   string   `json:"taggedEgressPorts,omitempty"`
            UntaggedEgressPorts []string `json:"untaggedEgressPorts,omitempty"`
            VlanIds             string   `json:"vlanIds,omitempty"`
        } `json:"vlanConfig,omitempty"`
    } `json:"vlans,omitempty"`
    MgmtAccess struct {
        TelnetConfig struct {
            TelnetEnabled bool `json:"telnetEnabled,omitempty"`
        } `json:"telnetConfig,omitempty"`
        SSHConfig struct {
            SSHEnabled bool `json:"sshEnabled,omitempty"`
        } `json:"sshConfig,omitempty"`
        HTTPConfig struct {
            HTTPEnabled bool `json:"httpEnabled,omitempty"`
        } `json:"httpConfig,omitempty"`
        HTTPSConfig struct {
            HTTPSEnabled bool `json:"httpsEnabled,omitempty"`
        } `json:"httpsConfig,omitempty"`
    } `json:"mgmtAccess,omitempty"`
    Appliance struct {
        PurviewConfig struct {
            NisEnabled        bool   `json:"nisEnabled,omitempty"`
            NisServerAddress  string `json:"nisServerAddress,omitempty"`
            NisDomainName     string `json:"nisDomainName,omitempty"`
            GreEnabled        bool   `json:"greEnabled,omitempty"`
            TapModeEnabled    bool   `json:"tapModeEnabled,omitempty"`
            TunnelModeEnabled bool   `json:"tunnelModeEnabled,omitempty"`
            GreConfig         []struct {
                ApplianceNetworkAddress string `json:"applianceNetworkAddress,omitempty"`
                CoreFlow2NetworkAddress string `json:"coreFlow2NetworkAddress,omitempty"`
            } `json:"greConfig,omitempty"`
            TapModeConfig []struct {
                TapInterface string `json:"tapInterface,omitempty"`
            } `json:"tapModeConfig,omitempty"`
            TunnelModeConfig []struct {
                TunnelInterface           string `json:"tunnelInterface,omitempty"`
                TunnelNetworkAddress      string `json:"tunnelNetworkAddress,omitempty"`
                TunnelNetworkGateway      string `json:"tunnelNetworkGateway,omitempty"`
                TunnelNetworkPrefixLength int    `json:"tunnelNetworkPrefixLength,omitempty"`
            } `json:"tunnelModeConfig,omitempty"`
        } `json:"purviewConfig,omitempty"`
        NacConfig struct {
            ApplianceGroup string `json:"applianceGroup,omitempty"`
        } `json:"nacConfig,omitempty"`
    } `json:"appliance,omitempty"`
    MLag struct {
        Enabled    bool `json:"enabled,omitempty"`
        MlagGroups []struct {
            MlagDetect bool     `json:"mlagDetect,omitempty"`
            MlagPeer   string   `json:"mlagPeer,omitempty"`
            MlagPorts  []string `json:"mlagPorts,omitempty"`
        } `json:"mlagGroups,omitempty"`
    } `json:"mLag,omitempty"`
    MLagv2 struct {
        Logging   string `json:"logging,omitempty"`
        MlagPeers []struct {
            MlagPeerID     string `json:"mlagPeerId,omitempty"`
            MlagPeerIPAddr string `json:"mlagPeerIpAddr,omitempty"`
            MlagpeerVr     string `json:"mlagpeerVr,omitempty"`
            MlagPeerVlanID int    `json:"mlagPeerVlanId,omitempty"`
        } `json:"mlagPeers,omitempty"`
        MlagMembers []struct {
            MlagPort       string `json:"mlagPort,omitempty"`
            MlagPortMlagID int    `json:"mlagPortMlagId,omitempty"`
            MlagPortPeerID string `json:"mlagPortPeerId,omitempty"`
        } `json:"mlagMembers,omitempty"`
    } `json:"mLagv2,omitempty"`
    Stack struct {
        Enabled  bool     `json:"enabled,omitempty"`
        Logging  string   `json:"logging,omitempty"`
        Topology string   `json:"topology,omitempty"`
        Ports    []string `json:"ports,omitempty"`
        Peers    struct {
            StackPeer string `json:"stackPeer,omitempty"`
        } `json:"peers,omitempty"`
    } `json:"stack,omitempty"`
    Mvrp struct {
        Enabled    string `json:"enabled,omitempty"`
        MvrpConfig []struct {
            MvrpEnable bool     `json:"mvrpEnable,omitempty"`
            MvrpPorts  []string `json:"mvrpPorts,omitempty"`
        } `json:"mvrpConfig,omitempty"`
    } `json:"mvrp,omitempty"`
    Ospf struct {
        Enabled bool   `json:"enabled,omitempty"`
        Logging string `json:"logging,omitempty"`
    } `json:"ospf,omitempty"`
    Timestamp   string `json:"timestamp,omitempty"`
    BpRequestID int    `json:"bpRequestId,omitempty"`
    Status      string `json:"status,omitempty"`
}
type Event struct {
    Type        int    `json:"type,omitempty"`
    Severity    string `json:"severity,omitempty"`
    Timestamp   string `json:"timestamp,omitempty"`
    Target      string `json:"target,omitempty"`
    Description string `json:"description,omitempty"`
}
type Asset struct {
    AssetName    string `json:"assetName,omitempty"`
    AssetVersion string `json:"assetVersion,omitempty"`
    AssetType    string `json:"assetType,omitempty"`
}

// C O N N E C T
type Connect struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock,omitempty"`
	DeviceInfo struct {
		SysDescr        string `json:"sysDescr,omitempty"`
		SysUpTime       int    `json:"sysUpTime,omitempty"`
		SysContact      string `json:"sysContact,omitempty"`
		SysName         string `json:"sysName,omitempty"`
		SysObjectID     string `json:"sysObjectID,omitempty"`
		OperatingSystem string `json:"operatingSystem,omitempty"`
		SerialNumber    string `json:"serialNumber,omitempty"`
		MacAddr         string `json:"macAddr,omitempty"`
		MgmtIPAddr      string `json:"mgmtIpAddr,omitempty"`
		MgmtPort        string `json:"mgmtPort,omitempty"`
		License         struct {
			FeaturePacks        []FeaturePack `json:"featurePacks,omitempty"`
			PortCapacityLicense string   `json:"portCapacityLicense,omitempty"`
			EffectiveLicense    string   `json:"effectiveLicense,omitempty"`
			EnabledLicense      string   `json:"enabledLicense,omitempty"`
			SystemLicense       string   `json:"systemLicense,omitempty"`
			EffectiveLevel      string   `json:"effectiveLevel,omitempty"`
		} `json:"license,omitempty"`
		Ports []struct {
			PortType  string   `json:"portType,omitempty"`
			PortSpeed int      `json:"portSpeed,omitempty"`
			PortList  []string `json:"portList,omitempty"`
		} `json:"ports,omitempty"`
		IfTable []struct {
			IfIndex           int    `json:"ifIndex,omitempty"`
			IfName            string `json:"ifName,omitempty"`
			IfDescr           string `json:"ifDescr,omitempty"`
			IfType            int    `json:"ifType,omitempty"`
			IfOutDiscards     int    `json:"ifOutDiscards,omitempty"`
			IfAdminStatus     int    `json:"ifAdminStatus,omitempty"`
			IfMtu             int    `json:"ifMtu,omitempty"`
			IfInDiscards      int    `json:"ifInDiscards,omitempty"`
			IfInErrors        int    `json:"ifInErrors,omitempty"`
			IfOperStatus      int    `json:"ifOperStatus,omitempty"`
			IfOutUcastPkts    int    `json:"ifOutUcastPkts,omitempty"`
			IfInOctets        int    `json:"ifInOctets,omitempty"`
			IfLastChange      int    `json:"ifLastChange,omitempty"`
			IfPhysAddress     string `json:"ifPhysAddress,omitempty"`
			IfInUcastPkts     int    `json:"ifInUcastPkts,omitempty"`
			IfSpecific        string `json:"ifSpecific,omitempty"`
			IfOutNUcastPkts   int    `json:"ifOutNUcastPkts,omitempty"`
			IfOutQLen         int    `json:"ifOutQLen,omitempty"`
			IfSpeed           int    `json:"ifSpeed,omitempty"`
			IfOutOctets       int    `json:"ifOutOctets,omitempty"`
			IfInUnknownProtos int    `json:"ifInUnknownProtos,omitempty"`
			IfInNUcastPkts    int    `json:"ifInNUcastPkts,omitempty"`
			IfOutErrors       int    `json:"ifOutErrors,omitempty"`
		} `json:"ifTable,omitempty"`
		PortsInfo []struct {
			PortName string `json:"portName,omitempty"`
			Lldp     struct {
				PortID             string `json:"portId,omitempty"`
				SystemName         string `json:"systemName,omitempty"`
				MgmtAddress        string `json:"mgmtAddress,omitempty"`
				ChassisID          string `json:"chassisId,omitempty"`
				SystemCapabilities string `json:"systemCapabilities,omitempty"`
			} `json:"lldp,omitempty"`
		} `json:"portsInfo,omitempty"`
	} `json:"deviceInfo,omitempty"`
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
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock,omitempty"`
	Assets []Asset `json:"assets,omitempty"`
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
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock,omitempty"`
	Capabilities Capabilities `json:"capabilities,omitempty"`
	ConfigBlock ConfigBlock `json:"configBlock,omitempty"`
}

// C O N F I G U R A T I O N   R E S P O N S E
type ConfigurationResponse struct {
	CliBlock    string `json:"cliBlock,omitempty"`
	ConfigBlock ConfigBlock `json:"configBlock,omitempty"`
}

// E V E N T
type EventMsg struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock,omitempty"`
	Event []Event `json:"event,omitempty"`
	ConfigACK ConfigBlock `json:"ConfigACK,omitempty"`
}

// S T A T S
type Stats struct {
	ApPropertyBlock     ApPropertyBlock     `json:"apPropertyBlock,omitempty"`
	Capabilities        Capabilities        `json:"capabilities,omitempty"`
	Assets []Asset `json:"assets,omitempty"`
	UpgradeAssets []Asset `json:"upgradeAssets,omitempty"`
	ConfigBlock         ConfigBlock         `json:"configBlock,omitempty"`
	Cooling struct {
		Trays []struct {
			State string `json:"state,omitempty"`
			Name  string `json:"name,omitempty"`
			Fans  []struct {
				Name  string `json:"name,omitempty"`
				State string `json:"state,omitempty"`
				Rpm   int    `json:"rpm,omitempty"`
			} `json:"fans,omitempty"`
		} `json:"trays,omitempty"`
	} `json:"cooling,omitempty"`
	IfTable []struct {
		IfIndex           int    `json:"ifIndex,omitempty"`
		IfName            string `json:"ifName,omitempty"`
		IfDescr           string `json:"ifDescr,omitempty"`
		IfType            int    `json:"ifType,omitempty"`
		IfOutDiscards     int    `json:"ifOutDiscards,omitempty"`
		IfAdminStatus     int    `json:"ifAdminStatus,omitempty"`
		IfMtu             int    `json:"ifMtu,omitempty"`
		IfInDiscards      int    `json:"ifInDiscards,omitempty"`
		IfInErrors        int    `json:"ifInErrors,omitempty"`
		IfOperStatus      int    `json:"ifOperStatus,omitempty"`
		IfOutUcastPkts    int    `json:"ifOutUcastPkts,omitempty"`
		IfInOctets        int    `json:"ifInOctets,omitempty"`
		IfLastChange      int    `json:"ifLastChange,omitempty"`
		IfPhysAddress     string `json:"ifPhysAddress,omitempty"`
		IfInUcastPkts     int    `json:"ifInUcastPkts,omitempty"`
		IfSpecific        string `json:"ifSpecific,omitempty"`
		IfOutNUcastPkts   int    `json:"ifOutNUcastPkts,omitempty"`
		IfOutQLen         int    `json:"ifOutQLen,omitempty"`
		IfSpeed           int    `json:"ifSpeed,omitempty"`
		IfOutOctets       int    `json:"ifOutOctets,omitempty"`
		IfInUnknownProtos int    `json:"ifInUnknownProtos,omitempty"`
		IfInNUcastPkts    int    `json:"ifInNUcastPkts,omitempty"`
		IfOutErrors       int    `json:"ifOutErrors,omitempty"`
	} `json:"ifTable,omitempty"`
	PortsInfo []struct {
		PortName   string `json:"portName,omitempty"`
		StatsCheck int    `json:"statsCheck,omitempty"`
		StpStatus  struct {
			StpState string `json:"stpState,omitempty"`
			StpRole  string `json:"stpRole,omitempty"`
		} `json:"stpStatus,omitempty"`
		Poe struct {
			Curr    int    `json:"curr,omitempty"`
			Power   int    `json:"power,omitempty"`
			Volts   int    `json:"volts,omitempty"`
			State   string `json:"state,omitempty"`
			Fault   string `json:"fault,omitempty"`
			PdClass string `json:"pdClass,omitempty"`
		} `json:"poe,omitempty"`
		Lldp struct {
			PortID             string   `json:"portId,omitempty"`
			SystemName         string   `json:"systemName,omitempty"`
			SystemCapabilities []string `json:"systemCapabilities,omitempty"`
			MgmtAddress        string   `json:"mgmtAddress,omitempty"`
			ChassisID          string   `json:"chassisId,omitempty"`
			Med                struct {
				ModelName    string   `json:"modelName,omitempty"`
				SwVersion    string   `json:"swVersion,omitempty"`
				Capabilities []string `json:"capabilities,omitempty"`
				Serial       string   `json:"serial,omitempty"`
				HwVersion    string   `json:"hwVersion,omitempty"`
				MdiPdValue   string   `json:"mdiPdValue,omitempty"`
			} `json:"med,omitempty"`
		} `json:"lldp,omitempty"`
		Lag struct {
			Enabled   bool   `json:"enabled,omitempty"`
			Lag       string `json:"lag,omitempty"`
			RxState   string `json:"rxState,omitempty"`
			SelLogic  string `json:"selLogic,omitempty"`
			MuxState  string `json:"muxState,omitempty"`
			RxPackets int    `json:"rxPackets,omitempty"`
			TxPackets int    `json:"txPackets,omitempty"`
		} `json:"lag,omitempty"`
		Link struct {
			AdminStatus      string `json:"adminStatus,omitempty"`
			LinkStatus       string `json:"linkStatus,omitempty"`
			LinkDuplexConfig string `json:"linkDuplexConfig,omitempty"`
			LinkDuplexActual string `json:"linkDuplexActual,omitempty"`
			LinkSpeedActual  string `json:"linkSpeedActual,omitempty"`
			LinkSpeedConfig  string `json:"linkSpeedConfig,omitempty"`
			RxBytes          int    `json:"rxBytes,omitempty"`
			RxPackets        int    `json:"rxPackets,omitempty"`
			RxErrors         int    `json:"rxErrors,omitempty"`
			TxBytes          int    `json:"txBytes,omitempty"`
			TxPackets        int    `json:"txPackets,omitempty"`
			TxErrors         int    `json:"txErrors,omitempty"`
		} `json:"link,omitempty"`
	} `json:"portsInfo,omitempty"`
	PowerSupply []struct {
		Module   string `json:"module,omitempty"`
		State    string `json:"state,omitempty"`
		PartInfo string `json:"partInfo,omitempty"`
		Input    int    `json:"input,omitempty"`
	} `json:"powerSupply,omitempty"`
	Temperatue []struct {
		Temp   int    `json:"temp,omitempty"`
		Status string `json:"status,omitempty"`
		Name   string `json:"name,omitempty"`
	} `json:"temperatue,omitempty"`
	Telemetry []struct {
		Path           string `json:"path,omitempty"`
		HTTPStatusCode int    `json:"httpStatusCode,omitempty"`
		Repsonse       struct {
		} `json:"repsonse,omitempty"`
	} `json:"telemetry,omitempty"`
	IetfVxlanVxlan struct {
		GlobalEnable  bool `json:"global-enable,omitempty"`
		VxlanInstance []struct {
			VxlanID           int    `json:"vxlan-id,omitempty"`
			Description       string `json:"description,omitempty"`
			UnknowUnicastDrop string `json:"unknow-unicast-drop,omitempty"`
			FilterVrrp        string `json:"filter-vrrp,omitempty"`
			VxlanEvpn         struct {
				RouteDistinguisher string `json:"route-distinguisher,omitempty"`
				VpnTargets         []struct {
					RtValue string `json:"rt-value,omitempty"`
					RtType  string `json:"rt-type,omitempty"`
				} `json:"vpn-targets,omitempty"`
			} `json:"vxlan-evpn,omitempty"`
		} `json:"vxlan-instance,omitempty"`
		VtepInstances []struct {
			VtepID                int    `json:"vtep-id,omitempty"`
			VtepName              string `json:"vtep-name,omitempty"`
			SourceInterface       string `json:"source-interface,omitempty"`
			MulticastIP           string `json:"multicast-ip,omitempty"`
			InnerVlanHandlingMode int    `json:"inner-vlan-handling-mode,omitempty"`
			BindVxlanID           []int  `json:"bind-vxlan-id,omitempty"`
			StaticVxlanTunnel     []struct {
				VxlanTunnelID   int    `json:"vxlan-tunnel-id,omitempty"`
				VxlanTunnelName string `json:"vxlan-tunnel-name,omitempty"`
				AddressFamily   []struct {
					Af                  string `json:"af,omitempty"`
					TunnelSourceIP      string `json:"tunnel-source-ip,omitempty"`
					TunnelDestinationIP string `json:"tunnel-destination-ip,omitempty"`
					BindVxlanID         []int  `json:"bind-vxlan-id,omitempty"`
				} `json:"address-family,omitempty"`
			} `json:"static-vxlan-tunnel,omitempty"`
			RedundancyGroupBind []struct {
				VxlanID         int `json:"vxlan-id,omitempty"`
				RedundancyGroup int `json:"redundancy-group,omitempty"`
			} `json:"redundancy-group-bind,omitempty"`
		} `json:"vtep-instances,omitempty"`
	} `json:"ietf-vxlan:vxlan,omitempty"`
	IetfVxlanVxlanState struct {
		VxlanTunnel []struct {
			LocalIP        string `json:"local-ip,omitempty"`
			RemoteIP       string `json:"remote-ip,omitempty"`
			StaticTunnelID int    `json:"static-tunnel-id,omitempty"`
			EvpnTunnelID   int    `json:"evpn-tunnel-id,omitempty"`
			Statistics     struct {
				InBytes            int `json:"in-bytes,omitempty"`
				OutBytes           int `json:"out-bytes,omitempty"`
				InPackets          int `json:"in-packets,omitempty"`
				OutPackets         int `json:"out-packets,omitempty"`
				TunnelVniStatistic []struct {
					VxlanID    int `json:"vxlan-id,omitempty"`
					InBytes    int `json:"in-bytes,omitempty"`
					OutBytes   int `json:"out-bytes,omitempty"`
					InPackets  int `json:"in-packets,omitempty"`
					OutPackets int `json:"out-packets,omitempty"`
				} `json:"tunnel-vni-statistic,omitempty"`
			} `json:"statistics,omitempty"`
		} `json:"vxlan-tunnel,omitempty"`
	} `json:"ietf-vxlan:vxlan-state,omitempty"`
	CPUUtilization struct {
		CPUMonitorInterval         int `json:"cpuMonitorInterval,omitempty"`
		CPUMonitorTotalUtilization int `json:"cpuMonitorTotalUtilization,omitempty"`
		CPUMonitorTable            []struct {
			Module            string `json:"module,omitempty"`
			ProcessName       string `json:"processName,omitempty"`
			ProcessID         int    `json:"processId,omitempty"`
			ProcessState      string `json:"processState,omitempty"`
			Utilization5Secs  int    `json:"utilization5secs,omitempty"`
			Utilization10Secs int    `json:"utilization10secs,omitempty"`
			Utilization30Secs int    `json:"utilization30secs,omitempty"`
			Utilization1Min   int    `json:"utilization1min,omitempty"`
			Utilization5Mins  int    `json:"utilization5mins,omitempty"`
			Utilization30Mins int    `json:"utilization30mins,omitempty"`
			Utilization1Hour  int    `json:"utilization1hour,omitempty"`
			MaxUtilization    int    `json:"maxUtilization,omitempty"`
			UserTime          int    `json:"userTime,omitempty"`
			SystemTime        int    `json:"systemTime,omitempty"`
		} `json:"cpuMonitorTable,omitempty"`
		CPUMonitorSystemTable []struct {
			Module            string `json:"module,omitempty"`
			Utilization5Secs  int    `json:"utilization5secs,omitempty"`
			Utilization10Secs int    `json:"utilization10secs,omitempty"`
			Utilization30Secs int    `json:"utilization30secs,omitempty"`
			Utilization1Min   int    `json:"utilization1min,omitempty"`
			Utilization5Mins  int    `json:"utilization5mins,omitempty"`
			Utilization30Mins int    `json:"utilization30mins,omitempty"`
			Utilization1Hour  int    `json:"utilization1hour,omitempty"`
			MaxUtilization    int    `json:"maxUtilization,omitempty"`
		} `json:"cpuMonitorSystemTable,omitempty"`
	} `json:"cpuUtilization,omitempty"`
	MemoryUtilization struct {
		MemoryMonitorTable []struct {
			Module      string `json:"module,omitempty"`
			ProcessName string `json:"processName,omitempty"`
			Usage       int    `json:"usage,omitempty"`
			Limit       int    `json:"limit,omitempty"`
		} `json:"memoryMonitorTable,omitempty"`
		MemoryMonitorSystemTable []struct {
			Module      string `json:"module,omitempty"`
			SystemTotal int    `json:"systemTotal,omitempty"`
			SystemFree  int    `json:"systemFree,omitempty"`
			SystemUsage int    `json:"systemUsage,omitempty"`
			UserUsage   int    `json:"userUsage,omitempty"`
		} `json:"memoryMonitorSystemTable,omitempty"`
	} `json:"memoryUtilization,omitempty"`
	Mlagv2 struct {
		MLagPeers []struct {
			MLagPeerID     string `json:"mLagPeerId,omitempty"`
			MLagPeerStatus string `json:"mLagPeerStatus,omitempty"`
			MLagPeerUpTime int    `json:"mLagPeerUpTime,omitempty"`
		} `json:"mLagPeers,omitempty"`
	} `json:"mlagv2,omitempty"`
	Timestamp   int    `json:"timestamp,omitempty"`
	SysUpTime   int    `json:"sysUpTime,omitempty"`
	CheckInTime string `json:"checkInTime,omitempty"`
	UpgradeTime string `json:"upgradeTime,omitempty"`
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
