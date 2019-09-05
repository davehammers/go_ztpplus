package ztp

type ApPropertyBlock struct {
	RuSerialNumber string `json:"ruSerialNumber,omitempty"`
	BpWiredMacaddr string `json:"bpWiredMacaddr,omitempty"`
	RuSwVersion    string `json:"ruSwVersion,omitempty"`
	RuModel        string `json:"ruModel,omitempty"`
}
type License struct {
	FeaturePacks        []string `json:"featurePacks,omitempty"`
	PortCapacityLicense string   `json:"portCapacityLicense,omitempty"`
	EffectiveLicense    string   `json:"effectiveLicense,omitempty"`
	EnabledLicense      string   `json:"enabledLicense,omitempty"`
	SystemLicense       string   `json:"systemLicense,omitempty"`
	EffectiveLevel      string   `json:"effectiveLevel,omitempty"`
}
type Ports struct {
	PortType  string   `json:"portType,omitempty"`
	PortSpeed int      `json:"portSpeed,omitempty"`
	PortList  []string `json:"portList,omitempty"`
}
type IfTable struct {
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
}
type Lldp struct {
	PortID             string `json:"portId,omitempty"`
	SystemName         string `json:"systemName,omitempty"`
	MgmtAddress        string `json:"mgmtAddress,omitempty"`
	ChassisID          string `json:"chassisId,omitempty"`
	SystemCapabilities string `json:"systemCapabilities,omitempty"`
}
type PortsInfo struct {
	PortName string `json:"portName,omitempty"`
	Lldp     Lldp   `json:"lldp,omitempty"`
}
type DeviceInfo struct {
	SysDescr        string      `json:"sysDescr,omitempty"`
	SysUpTime       int         `json:"sysUpTime,omitempty"`
	SysContact      string      `json:"sysContact,omitempty"`
	SysName         string      `json:"sysName,omitempty"`
	SysObjectID     string      `json:"sysObjectID,omitempty"`
	OperatingSystem string      `json:"operatingSystem,omitempty"`
	SerialNumber    string      `json:"serialNumber,omitempty"`
	MacAddr         string      `json:"macAddr,omitempty"`
	MgmtIPAddr      string      `json:"mgmtIpAddr,omitempty"`
	MgmtPort        string      `json:"mgmtPort,omitempty"`
	License         License     `json:"license,omitempty"`
	Ports           []Ports     `json:"ports,omitempty"`
	IfTable         []IfTable   `json:"ifTable,omitempty"`
	PortsInfo       []PortsInfo `json:"portsInfo,omitempty"`
}

type Assets struct {
	AssetName    string `json:"assetName,omitempty"`
	AssetVersion string `json:"assetVersion,omitempty"`
	AssetType    string `json:"assetType,omitempty"`
}

type ConfigDownload struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Dot1X struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Lacp struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
	MaxLagCount      int  `json:"maxLagCount,omitempty"`
}
type CustomTlvs struct {
	Oui     string `json:"oui,omitempty"`
	Subtype int    `json:"subtype,omitempty"`
}
type Logins struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type MacAuth struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Mlag struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
	MaxMlagCount     int  `json:"maxMlagCount,omitempty"`
}
type Telnet struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type SSH struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type HTTP struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type HTTPS struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type MgmtAccess struct {
	Telnet Telnet `json:"telnet,omitempty"`
	SSH    SSH    `json:"ssh,omitempty"`
	HTTP   HTTP   `json:"http,omitempty"`
	HTTPS  HTTPS  `json:"https,omitempty"`
}
type Mvrp struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Poe struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
	PowerBudget      int  `json:"powerBudget,omitempty"`
}
type PortCapabilities struct {
	Ports           []string `json:"ports,omitempty"`
	PortSpeed       string   `json:"portSpeed,omitempty"`
	PortDuplex      string   `json:"portDuplex,omitempty"`
	AutoNegotiation bool     `json:"autoNegotiation,omitempty"`
}
type DesiredPortSettings struct {
	Ports           []string `json:"ports,omitempty"`
	PortSpeed       string   `json:"portSpeed,omitempty"`
	PortDuplex      string   `json:"portDuplex,omitempty"`
	AutoNegotiation bool     `json:"autoNegotiation,omitempty"`
}
type RadiusServers struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type SnmpV1 struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type SnmpV2C struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type SnmpV3 struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Snmp struct {
	SnmpV1  SnmpV1  `json:"snmpV1,omitempty"`
	SnmpV2C SnmpV2C `json:"snmpV2c,omitempty"`
	SnmpV3  SnmpV3  `json:"snmpV3,omitempty"`
}
type SpanningTree struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
	SpanGuard        bool `json:"spanGuard,omitempty"`
	LoopProtect      bool `json:"loopProtect,omitempty"`
}
type Syslog struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Stacking struct {
	FeatureAvailable bool     `json:"featureAvailable,omitempty"`
	Ports            []string `json:"ports,omitempty"`
}
type Telemetry struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Vlans struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
	MaxNtpServers    int  `json:"maxNtpServers,omitempty"`
	MaxDNSServers    int  `json:"maxDnsServers,omitempty"`
}
type Vxlan struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Vpex struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Eee struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type LocatorLed struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Mlagv2 struct {
	FeatureAvailable bool `json:"featureAvailable,omitempty"`
}
type Capabilities struct {
	ConfigDownload ConfigDownload `json:"configDownload,omitempty"`
	Dot1X          Dot1X          `json:"dot1x,omitempty"`
	Lacp           Lacp           `json:"lacp,omitempty"`
	License        License        `json:"license,omitempty"`
	Lldp           Lldp           `json:"lldp,omitempty"`
	Logins         Logins         `json:"logins,omitempty"`
	MacAuth        MacAuth        `json:"macAuth,omitempty"`
	Mlag           Mlag           `json:"mlag,omitempty"`
	MgmtAccess     MgmtAccess     `json:"mgmtAccess,omitempty"`
	Mvrp           Mvrp           `json:"mvrp,omitempty"`
	Poe            Poe            `json:"poe,omitempty"`
	Ports          Ports          `json:"ports,omitempty"`
	RadiusServers  RadiusServers  `json:"radiusServers,omitempty"`
	Snmp           Snmp           `json:"snmp,omitempty"`
	SpanningTree   SpanningTree   `json:"spanningTree,omitempty"`
	Syslog         Syslog         `json:"syslog,omitempty"`
	Stacking       Stacking       `json:"stacking,omitempty"`
	Telemetry      Telemetry      `json:"telemetry,omitempty"`
	Vlans          Vlans          `json:"vlans,omitempty"`
	Vxlan          Vxlan          `json:"vxlan,omitempty"`
	Vpex           Vpex           `json:"vpex,omitempty"`
	Eee            Eee            `json:"eee,omitempty"`
	LocatorLed     LocatorLed     `json:"locatorLed,omitempty"`
	Mlagv2         Mlagv2         `json:"mlagv2,omitempty"`
}
type FeaturePacks struct {
	FeaturePackName    string `json:"featurePackName,omitempty"`
	FeaturePackLicense string `json:"featurePackLicense,omitempty"`
}
type Dot1XConfig struct {
	Dot1XEnable bool     `json:"dot1XEnable,omitempty"`
	Dot1XPorts  []string `json:"dot1XPorts,omitempty"`
}
type LldpConfig struct {
	LldpEnable bool     `json:"lldpEnable,omitempty"`
	LldpPorts  []string `json:"lldpPorts,omitempty"`
}
type LoginConfig struct {
	AccessLevel string `json:"accessLevel,omitempty"`
	Password    string `json:"password,omitempty"`
	Username    string `json:"username,omitempty"`
}
type MacAuthConfig struct {
	MacAuthEnable bool   `json:"macAuthEnable,omitempty"`
	MacAuthPorts  string `json:"macAuthPorts,omitempty"`
}
type PortLinkType struct {
	LinkType      string   `json:"linkType,omitempty"`
	LinkTypePorts []string `json:"linkTypePorts,omitempty"`
}
type SpanGuard struct {
	SpanGuardEnabled bool     `json:"spanGuardEnabled,omitempty"`
	SpanGuardPorts   []string `json:"spanGuardPorts,omitempty"`
}
type LoopProtect struct {
	LoopProtectEnabled bool     `json:"loopProtectEnabled,omitempty"`
	LoopProtectPorts   []string `json:"loopProtectPorts,omitempty"`
}
type SyslogConfig struct {
	ServerNetworkAddress string `json:"serverNetworkAddress,omitempty"`
	ServerUDPPort        int    `json:"serverUdpPort,omitempty"`
}
type VlanConfig struct {
	Name                string   `json:"name,omitempty"`
	NetworkAddress      string   `json:"networkAddress,omitempty"`
	NetworkPrefixLength int      `json:"networkPrefixLength,omitempty"`
	TaggedEgressPorts   string   `json:"taggedEgressPorts,omitempty"`
	UntaggedEgressPorts []string `json:"untaggedEgressPorts,omitempty"`
	VlanIds             string   `json:"vlanIds,omitempty"`
}
type TelnetConfig struct {
	TelnetEnabled bool `json:"telnetEnabled,omitempty"`
}
type SSHConfig struct {
	SSHEnabled bool `json:"sshEnabled,omitempty"`
}
type HTTPConfig struct {
	HTTPEnabled bool `json:"httpEnabled,omitempty"`
}
type HTTPSConfig struct {
	HTTPSEnabled bool `json:"httpsEnabled,omitempty"`
}
type GreConfig struct {
	ApplianceNetworkAddress string `json:"applianceNetworkAddress,omitempty"`
	CoreFlow2NetworkAddress string `json:"coreFlow2NetworkAddress,omitempty"`
}
type TapModeConfig struct {
	TapInterface string `json:"tapInterface,omitempty"`
}
type TunnelModeConfig struct {
	TunnelInterface           string `json:"tunnelInterface,omitempty"`
	TunnelNetworkAddress      string `json:"tunnelNetworkAddress,omitempty"`
	TunnelNetworkGateway      string `json:"tunnelNetworkGateway,omitempty"`
	TunnelNetworkPrefixLength int    `json:"tunnelNetworkPrefixLength,omitempty"`
}
type PurviewConfig struct {
	NisEnabled        bool               `json:"nisEnabled,omitempty"`
	NisServerAddress  string             `json:"nisServerAddress,omitempty"`
	NisDomainName     string             `json:"nisDomainName,omitempty"`
	GreEnabled        bool               `json:"greEnabled,omitempty"`
	TapModeEnabled    bool               `json:"tapModeEnabled,omitempty"`
	TunnelModeEnabled bool               `json:"tunnelModeEnabled,omitempty"`
	GreConfig         []GreConfig        `json:"greConfig,omitempty"`
	TapModeConfig     []TapModeConfig    `json:"tapModeConfig,omitempty"`
	TunnelModeConfig  []TunnelModeConfig `json:"tunnelModeConfig,omitempty"`
}
type NacConfig struct {
	ApplianceGroup string `json:"applianceGroup,omitempty"`
}
type Appliance struct {
	PurviewConfig PurviewConfig `json:"purviewConfig,omitempty"`
	NacConfig     NacConfig     `json:"nacConfig,omitempty"`
}
type MlagGroups struct {
	MlagDetect bool     `json:"mlagDetect,omitempty"`
	MlagPeer   string   `json:"mlagPeer,omitempty"`
	MlagPorts  []string `json:"mlagPorts,omitempty"`
}
type MLag struct {
	Enabled    bool         `json:"enabled,omitempty"`
	MlagGroups []MlagGroups `json:"mlagGroups,omitempty"`
}
type MlagPeers struct {
	MlagPeerID     string `json:"mlagPeerId,omitempty"`
	MlagPeerIPAddr string `json:"mlagPeerIpAddr,omitempty"`
	MlagpeerVr     string `json:"mlagpeerVr,omitempty"`
	MlagPeerVlanID int    `json:"mlagPeerVlanId,omitempty"`
}
type MlagMembers struct {
	MlagPort       string `json:"mlagPort,omitempty"`
	MlagPortMlagID int    `json:"mlagPortMlagId,omitempty"`
	MlagPortPeerID string `json:"mlagPortPeerId,omitempty"`
}
type MLagv2 struct {
	Logging     string        `json:"logging,omitempty"`
	MlagPeers   []MlagPeers   `json:"mlagPeers,omitempty"`
	MlagMembers []MlagMembers `json:"mlagMembers,omitempty"`
}
type Peers struct {
	StackPeer string `json:"stackPeer,omitempty"`
}
type Stack struct {
	Enabled  bool     `json:"enabled,omitempty"`
	Logging  string   `json:"logging,omitempty"`
	Topology string   `json:"topology,omitempty"`
	Ports    []string `json:"ports,omitempty"`
	Peers    Peers    `json:"peers,omitempty"`
}
type MvrpConfig struct {
	MvrpEnable bool     `json:"mvrpEnable,omitempty"`
	MvrpPorts  []string `json:"mvrpPorts,omitempty"`
}
type Ospf struct {
	Enabled bool   `json:"enabled,omitempty"`
	Logging string `json:"logging,omitempty"`
}
type ConfigBlock struct {
	License      License      `json:"license,omitempty"`
	Poe          Poe          `json:"poe,omitempty"`
	Dot1X        Dot1X        `json:"dot1X,omitempty"`
	Lacp         Lacp         `json:"lacp,omitempty"`
	Lldp         Lldp         `json:"lldp,omitempty"`
	Logins       Logins       `json:"logins,omitempty"`
	MacAuth      MacAuth      `json:"macAuth,omitempty"`
	Ports        []Ports      `json:"ports,omitempty"`
	Snmp         Snmp         `json:"snmp,omitempty"`
	SpanningTree SpanningTree `json:"spanningTree,omitempty"`
	Syslog       Syslog       `json:"syslog,omitempty"`
	Vlans        Vlans        `json:"vlans,omitempty"`
	MgmtAccess   MgmtAccess   `json:"mgmtAccess,omitempty"`
	Appliance    Appliance    `json:"appliance,omitempty"`
	MLag         MLag         `json:"mLag,omitempty"`
	MLagv2       MLagv2       `json:"mLagv2,omitempty"`
	Stack        Stack        `json:"stack,omitempty"`
	Mvrp         Mvrp         `json:"mvrp,omitempty"`
	Ospf         Ospf         `json:"ospf,omitempty"`
	Timestamp    string       `json:"timestamp,omitempty"`
	BpRequestID  int          `json:"bpRequestId,omitempty"`
	Status       string       `json:"status,omitempty"`
}

type Event struct {
	Type        int    `json:"type,omitempty"`
	Severity    string `json:"severity,omitempty"`
	Timestamp   string `json:"timestamp,omitempty"`
	Target      string `json:"target,omitempty"`
	Description string `json:"description,omitempty"`
}

type StackPeerConfig struct {
	StackPeer string `json:"stackPeer,omitempty"`
}
type StackingPeer struct {
	Enabled         bool              `json:"enabled,omitempty"`
	StackPeerConfig []StackPeerConfig `json:"stackPeerConfig,omitempty"`
	StackPortConfig []string          `json:"stackPortConfig,omitempty"`
	StackTopology   string            `json:"stackTopology,omitempty"`
}
type Authentication struct {
	Type     string `json:"type,omitempty"`
	Password string `json:"password,omitempty"`
}
type DhcpRelayAgents struct {
	IPAddress string `json:"ipAddress,omitempty"`
}
type StaticRoutes struct {
	Network   string `json:"network,omitempty"`
	NextHop   string `json:"nextHop,omitempty"`
	IPVersion string `json:"ipVersion,omitempty"`
}
type Hosts struct {
	Vrf          string         `json:"vrf,omitempty"`
	IPAddress    string         `json:"ipAddress,omitempty"`
	IPVersion    string         `json:"ipVersion,omitempty"`
	StaticRoutes []StaticRoutes `json:"staticRoutes,omitempty"`
}
type ExtremeFabric struct {
	Enabled         bool              `json:"enabled,omitempty"`
	FabricDomain    string            `json:"fabricDomain,omitempty"`
	Authentication  Authentication    `json:"authentication,omitempty"`
	DhcpRelayAgents []DhcpRelayAgents `json:"dhcpRelayAgents,omitempty"`
	Hosts           []Hosts           `json:"hosts,omitempty"`
}
type ConfigACK struct {
	License       License         `json:"license,omitempty"`
	Poe           Poe             `json:"poe,omitempty"`
	Dot1X         Dot1X           `json:"dot1X,omitempty"`
	Lacp          Lacp            `json:"lacp,omitempty"`
	Lldp          Lldp            `json:"lldp,omitempty"`
	Logins        Logins          `json:"logins,omitempty"`
	MacAuth       MacAuth         `json:"macAuth,omitempty"`
	Ports         []Ports         `json:"ports,omitempty"`
	Snmp          Snmp            `json:"snmp,omitempty"`
	SpanningTree  SpanningTree    `json:"spanningTree,omitempty"`
	Syslog        Syslog          `json:"syslog,omitempty"`
	Vlans         Vlans           `json:"vlans,omitempty"`
	MgmtAccess    MgmtAccess      `json:"mgmtAccess,omitempty"`
	Mlag          Mlag            `json:"mlag,omitempty"`
	Stacking      StackingPeer    `json:"stacking,omitempty"`
	Mvrp          Mvrp            `json:"mvrp,omitempty"`
	ExtremeFabric ExtremeFabric   `json:"extremeFabric,omitempty"`
	Vxlan         Vxlan           `json:"vxlan,omitempty"`
	RadiusServers []RadiusServers `json:"radiusServers,omitempty"`
	Timestamp     string          `json:"timestamp,omitempty"`
	BpRequestID   int             `json:"bpRequestId,omitempty"`
	Status        string          `json:"status,omitempty"`
}

type UpgradeAssets struct {
	AssetName    string `json:"assetName,omitempty"`
	AssetVersion string `json:"assetVersion,omitempty"`
	AssetType    string `json:"assetType,omitempty"`
}
type Fans struct {
	Name  string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
	Rpm   int    `json:"rpm,omitempty"`
}
type Trays struct {
	State string `json:"state,omitempty"`
	Name  string `json:"name,omitempty"`
	Fans  []Fans `json:"fans,omitempty"`
}
type Cooling struct {
	Trays []Trays `json:"trays,omitempty"`
}
type StpStatus struct {
	StpState string `json:"stpState,omitempty"`
	StpRole  string `json:"stpRole,omitempty"`
}
type Med struct {
	ModelName    string   `json:"modelName,omitempty"`
	SwVersion    string   `json:"swVersion,omitempty"`
	Capabilities []string `json:"capabilities,omitempty"`
	Serial       string   `json:"serial,omitempty"`
	HwVersion    string   `json:"hwVersion,omitempty"`
	MdiPdValue   string   `json:"mdiPdValue,omitempty"`
}
type Lag struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Lag       string `json:"lag,omitempty"`
	RxState   string `json:"rxState,omitempty"`
	SelLogic  string `json:"selLogic,omitempty"`
	MuxState  string `json:"muxState,omitempty"`
	RxPackets int    `json:"rxPackets,omitempty"`
	TxPackets int    `json:"txPackets,omitempty"`
}
type Link struct {
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
}
type PowerSupply struct {
	Module   string `json:"module,omitempty"`
	State    string `json:"state,omitempty"`
	PartInfo string `json:"partInfo,omitempty"`
	Input    int    `json:"input,omitempty"`
}
type Temperatue struct {
	Temp   int    `json:"temp,omitempty"`
	Status string `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
}
type Repsonse struct {
}
type VpnTargets struct {
	RtValue string `json:"rt-value,omitempty"`
	RtType  string `json:"rt-type,omitempty"`
}
type VxlanEvpn struct {
	RouteDistinguisher string       `json:"route-distinguisher,omitempty"`
	VpnTargets         []VpnTargets `json:"vpn-targets,omitempty"`
}
type VxlanInstance struct {
	VxlanID           int       `json:"vxlan-id,omitempty"`
	Description       string    `json:"description,omitempty"`
	UnknowUnicastDrop string    `json:"unknow-unicast-drop,omitempty"`
	FilterVrrp        string    `json:"filter-vrrp,omitempty"`
	VxlanEvpn         VxlanEvpn `json:"vxlan-evpn,omitempty"`
}
type AddressFamily struct {
	Af                  string `json:"af,omitempty"`
	TunnelSourceIP      string `json:"tunnel-source-ip,omitempty"`
	TunnelDestinationIP string `json:"tunnel-destination-ip,omitempty"`
	BindVxlanID         []int  `json:"bind-vxlan-id,omitempty"`
}
type StaticVxlanTunnel struct {
	VxlanTunnelID   int             `json:"vxlan-tunnel-id,omitempty"`
	VxlanTunnelName string          `json:"vxlan-tunnel-name,omitempty"`
	AddressFamily   []AddressFamily `json:"address-family,omitempty"`
}
type RedundancyGroupBind struct {
	VxlanID         int `json:"vxlan-id,omitempty"`
	RedundancyGroup int `json:"redundancy-group,omitempty"`
}
type VtepInstances struct {
	VtepID                int                   `json:"vtep-id,omitempty"`
	VtepName              string                `json:"vtep-name,omitempty"`
	SourceInterface       string                `json:"source-interface,omitempty"`
	MulticastIP           string                `json:"multicast-ip,omitempty"`
	InnerVlanHandlingMode int                   `json:"inner-vlan-handling-mode,omitempty"`
	BindVxlanID           []int                 `json:"bind-vxlan-id,omitempty"`
	StaticVxlanTunnel     []StaticVxlanTunnel   `json:"static-vxlan-tunnel,omitempty"`
	RedundancyGroupBind   []RedundancyGroupBind `json:"redundancy-group-bind,omitempty"`
}
type IetfVxlanVxlan struct {
	GlobalEnable  bool            `json:"global-enable,omitempty"`
	VxlanInstance []VxlanInstance `json:"vxlan-instance,omitempty"`
	VtepInstances []VtepInstances `json:"vtep-instances,omitempty"`
}
type TunnelVniStatistic struct {
	VxlanID    int `json:"vxlan-id,omitempty"`
	InBytes    int `json:"in-bytes,omitempty"`
	OutBytes   int `json:"out-bytes,omitempty"`
	InPackets  int `json:"in-packets,omitempty"`
	OutPackets int `json:"out-packets,omitempty"`
}
type Statistics struct {
	InBytes            int                  `json:"in-bytes,omitempty"`
	OutBytes           int                  `json:"out-bytes,omitempty"`
	InPackets          int                  `json:"in-packets,omitempty"`
	OutPackets         int                  `json:"out-packets,omitempty"`
	TunnelVniStatistic []TunnelVniStatistic `json:"tunnel-vni-statistic,omitempty"`
}
type VxlanTunnel struct {
	LocalIP        string     `json:"local-ip,omitempty"`
	RemoteIP       string     `json:"remote-ip,omitempty"`
	StaticTunnelID int        `json:"static-tunnel-id,omitempty"`
	EvpnTunnelID   int        `json:"evpn-tunnel-id,omitempty"`
	Statistics     Statistics `json:"statistics,omitempty"`
}
type IetfVxlanVxlanState struct {
	VxlanTunnel []VxlanTunnel `json:"vxlan-tunnel,omitempty"`
}
type CPUMonitorTable struct {
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
}
type CPUMonitorSystemTable struct {
	Module            string `json:"module,omitempty"`
	Utilization5Secs  int    `json:"utilization5secs,omitempty"`
	Utilization10Secs int    `json:"utilization10secs,omitempty"`
	Utilization30Secs int    `json:"utilization30secs,omitempty"`
	Utilization1Min   int    `json:"utilization1min,omitempty"`
	Utilization5Mins  int    `json:"utilization5mins,omitempty"`
	Utilization30Mins int    `json:"utilization30mins,omitempty"`
	Utilization1Hour  int    `json:"utilization1hour,omitempty"`
	MaxUtilization    int    `json:"maxUtilization,omitempty"`
}
type CPUUtilization struct {
	CPUMonitorInterval         int                     `json:"cpuMonitorInterval,omitempty"`
	CPUMonitorTotalUtilization int                     `json:"cpuMonitorTotalUtilization,omitempty"`
	CPUMonitorTable            []CPUMonitorTable       `json:"cpuMonitorTable,omitempty"`
	CPUMonitorSystemTable      []CPUMonitorSystemTable `json:"cpuMonitorSystemTable,omitempty"`
}
type MemoryMonitorTable struct {
	Module      string `json:"module,omitempty"`
	ProcessName string `json:"processName,omitempty"`
	Usage       int    `json:"usage,omitempty"`
	Limit       int    `json:"limit,omitempty"`
}
type MemoryMonitorSystemTable struct {
	Module      string `json:"module,omitempty"`
	SystemTotal int    `json:"systemTotal,omitempty"`
	SystemFree  int    `json:"systemFree,omitempty"`
	SystemUsage int    `json:"systemUsage,omitempty"`
	UserUsage   int    `json:"userUsage,omitempty"`
}
type MemoryUtilization struct {
	MemoryMonitorTable       []MemoryMonitorTable       `json:"memoryMonitorTable,omitempty"`
	MemoryMonitorSystemTable []MemoryMonitorSystemTable `json:"memoryMonitorSystemTable,omitempty"`
}
type MLagPeers struct {
	MLagPeerID     string `json:"mLagPeerId,omitempty"`
	MLagPeerStatus string `json:"mLagPeerStatus,omitempty"`
	MLagPeerUpTime int    `json:"mLagPeerUpTime,omitempty"`
}

// C O N N E C T
type Connect struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock,omitempty"`
	DeviceInfo      DeviceInfo      `json:"deviceInfo,omitempty"`
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
	Assets          []Assets        `json:"assets,omitempty"`
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
	Capabilities    Capabilities    `json:"capabilities,omitempty"`
	ConfigBlock     ConfigBlock     `json:"configBlock,omitempty"`
}

// E V E N T
type EventMsg struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock,omitempty"`
	Event           []Event         `json:"event,omitempty"`
	ConfigACK       ConfigACK       `json:"ConfigACK,omitempty"`
}

// S T A T S
type Stats struct {
	ApPropertyBlock     ApPropertyBlock     `json:"apPropertyBlock,omitempty"`
	Capabilities        Capabilities        `json:"capabilities,omitempty"`
	Assets              []Assets            `json:"assets,omitempty"`
	UpgradeAssets       []UpgradeAssets     `json:"upgradeAssets,omitempty"`
	ConfigBlock         ConfigBlock         `json:"configBlock,omitempty"`
	Cooling             Cooling             `json:"cooling,omitempty"`
	IfTable             []IfTable           `json:"ifTable,omitempty"`
	PortsInfo           []PortsInfo         `json:"portsInfo,omitempty"`
	PowerSupply         []PowerSupply       `json:"powerSupply,omitempty"`
	Temperatue          []Temperatue        `json:"temperatue,omitempty"`
	Telemetry           []Telemetry         `json:"telemetry,omitempty"`
	IetfVxlanVxlan      IetfVxlanVxlan      `json:"ietf-vxlan:vxlan,omitempty"`
	IetfVxlanVxlanState IetfVxlanVxlanState `json:"ietf-vxlan:vxlan-state,omitempty"`
	CPUUtilization      CPUUtilization      `json:"cpuUtilization,omitempty"`
	MemoryUtilization   MemoryUtilization   `json:"memoryUtilization,omitempty"`
	Mlagv2              Mlagv2              `json:"mlagv2,omitempty"`
	Timestamp           int                 `json:"timestamp,omitempty"`
	SysUpTime           int                 `json:"sysUpTime,omitempty"`
	CheckInTime         string              `json:"checkInTime,omitempty"`
	UpgradeTime         string              `json:"upgradeTime,omitempty"`
}
