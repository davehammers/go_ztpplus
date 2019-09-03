package ztp

type ApPropertyBlock struct {
	RuSerialNumber string `json:"ruSerialNumber"`
	BpWiredMacaddr string `json:"bpWiredMacaddr"`
	RuSwVersion    string `json:"ruSwVersion"`
	RuModel        string `json:"ruModel"`
}
type License struct {
	FeaturePacks        []string `json:"featurePacks"`
	PortCapacityLicense string   `json:"portCapacityLicense"`
	EffectiveLicense    string   `json:"effectiveLicense"`
	EnabledLicense      string   `json:"enabledLicense"`
	SystemLicense       string   `json:"systemLicense"`
	EffectiveLevel      string   `json:"effectiveLevel"`
}
type Ports struct {
	PortType  string   `json:"portType"`
	PortSpeed int      `json:"portSpeed"`
	PortList  []string `json:"portList"`
}
type IfTable struct {
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
}
type Lldp struct {
	PortID             string `json:"portId"`
	SystemName         string `json:"systemName"`
	MgmtAddress        string `json:"mgmtAddress"`
	ChassisID          string `json:"chassisId"`
	SystemCapabilities string `json:"systemCapabilities"`
}
type PortsInfo struct {
	PortName string `json:"portName"`
	Lldp     Lldp   `json:"lldp"`
}
type DeviceInfo struct {
	SysDescr        string      `json:"sysDescr"`
	SysUpTime       int         `json:"sysUpTime"`
	SysContact      string      `json:"sysContact"`
	SysName         string      `json:"sysName"`
	SysObjectID     string      `json:"sysObjectID"`
	OperatingSystem string      `json:"operatingSystem"`
	SerialNumber    string      `json:"serialNumber"`
	MacAddr         string      `json:"macAddr"`
	MgmtIPAddr      string      `json:"mgmtIpAddr"`
	MgmtPort        string      `json:"mgmtPort"`
	License         License     `json:"license"`
	Ports           []Ports     `json:"ports"`
	IfTable         []IfTable   `json:"ifTable"`
	PortsInfo       []PortsInfo `json:"portsInfo"`
}

type Assets struct {
	AssetName    string `json:"assetName"`
	AssetVersion string `json:"assetVersion"`
	AssetType    string `json:"assetType"`
}

type ConfigDownload struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Dot1X struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Lacp struct {
	FeatureAvailable bool `json:"featureAvailable"`
	MaxLagCount      int  `json:"maxLagCount"`
}
type CustomTlvs struct {
	Oui     string `json:"oui"`
	Subtype int    `json:"subtype"`
}
type Logins struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type MacAuth struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Mlag struct {
	FeatureAvailable bool `json:"featureAvailable"`
	MaxMlagCount     int  `json:"maxMlagCount"`
}
type Telnet struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type SSH struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type HTTP struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type HTTPS struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type MgmtAccess struct {
	Telnet Telnet `json:"telnet"`
	SSH    SSH    `json:"ssh"`
	HTTP   HTTP   `json:"http"`
	HTTPS  HTTPS  `json:"https"`
}
type Mvrp struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Poe struct {
	FeatureAvailable bool `json:"featureAvailable"`
	PowerBudget      int  `json:"powerBudget"`
}
type PortCapabilities struct {
	Ports           []string `json:"ports"`
	PortSpeed       string   `json:"portSpeed"`
	PortDuplex      string   `json:"portDuplex"`
	AutoNegotiation bool     `json:"autoNegotiation"`
}
type DesiredPortSettings struct {
	Ports           []string `json:"ports"`
	PortSpeed       string   `json:"portSpeed"`
	PortDuplex      string   `json:"portDuplex"`
	AutoNegotiation bool     `json:"autoNegotiation"`
}
type RadiusServers struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type SnmpV1 struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type SnmpV2C struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type SnmpV3 struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Snmp struct {
	SnmpV1  SnmpV1  `json:"snmpV1"`
	SnmpV2C SnmpV2C `json:"snmpV2c"`
	SnmpV3  SnmpV3  `json:"snmpV3"`
}
type SpanningTree struct {
	FeatureAvailable bool `json:"featureAvailable"`
	SpanGuard        bool `json:"spanGuard"`
	LoopProtect      bool `json:"loopProtect"`
}
type Syslog struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Stacking struct {
	FeatureAvailable bool     `json:"featureAvailable"`
	Ports            []string `json:"ports"`
}
type Telemetry struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Vlans struct {
	FeatureAvailable bool `json:"featureAvailable"`
	MaxNtpServers    int  `json:"maxNtpServers"`
	MaxDNSServers    int  `json:"maxDnsServers"`
}
type Vxlan struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Vpex struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Eee struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type LocatorLed struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Mlagv2 struct {
	FeatureAvailable bool `json:"featureAvailable"`
}
type Capabilities struct {
	ConfigDownload ConfigDownload `json:"configDownload"`
	Dot1X          Dot1X          `json:"dot1x"`
	Lacp           Lacp           `json:"lacp"`
	License        License        `json:"license"`
	Lldp           Lldp           `json:"lldp"`
	Logins         Logins         `json:"logins"`
	MacAuth        MacAuth        `json:"macAuth"`
	Mlag           Mlag           `json:"mlag"`
	MgmtAccess     MgmtAccess     `json:"mgmtAccess"`
	Mvrp           Mvrp           `json:"mvrp"`
	Poe            Poe            `json:"poe"`
	Ports          Ports          `json:"ports"`
	RadiusServers  RadiusServers  `json:"radiusServers"`
	Snmp           Snmp           `json:"snmp"`
	SpanningTree   SpanningTree   `json:"spanningTree"`
	Syslog         Syslog         `json:"syslog"`
	Stacking       Stacking       `json:"stacking"`
	Telemetry      Telemetry      `json:"telemetry"`
	Vlans          Vlans          `json:"vlans"`
	Vxlan          Vxlan          `json:"vxlan"`
	Vpex           Vpex           `json:"vpex"`
	Eee            Eee            `json:"eee"`
	LocatorLed     LocatorLed     `json:"locatorLed"`
	Mlagv2         Mlagv2         `json:"mlagv2"`
}
type FeaturePacks struct {
	FeaturePackName    string `json:"featurePackName"`
	FeaturePackLicense string `json:"featurePackLicense"`
}
type Dot1XConfig struct {
	Dot1XEnable bool     `json:"dot1XEnable"`
	Dot1XPorts  []string `json:"dot1XPorts"`
}
type LldpConfig struct {
	LldpEnable bool     `json:"lldpEnable"`
	LldpPorts  []string `json:"lldpPorts"`
}
type LoginConfig struct {
	AccessLevel string `json:"accessLevel"`
	Password    string `json:"password"`
	Username    string `json:"username"`
}
type MacAuthConfig struct {
	MacAuthEnable bool   `json:"macAuthEnable"`
	MacAuthPorts  string `json:"macAuthPorts"`
}
type PortLinkType struct {
	LinkType      string   `json:"linkType"`
	LinkTypePorts []string `json:"linkTypePorts"`
}
type SpanGuard struct {
	SpanGuardEnabled bool     `json:"spanGuardEnabled"`
	SpanGuardPorts   []string `json:"spanGuardPorts"`
}
type LoopProtect struct {
	LoopProtectEnabled bool     `json:"loopProtectEnabled"`
	LoopProtectPorts   []string `json:"loopProtectPorts"`
}
type SyslogConfig struct {
	ServerNetworkAddress string `json:"serverNetworkAddress"`
	ServerUDPPort        int    `json:"serverUdpPort"`
}
type VlanConfig struct {
	Name                string   `json:"name"`
	NetworkAddress      string   `json:"networkAddress"`
	NetworkPrefixLength int      `json:"networkPrefixLength"`
	TaggedEgressPorts   string   `json:"taggedEgressPorts"`
	UntaggedEgressPorts []string `json:"untaggedEgressPorts"`
	VlanIds             string   `json:"vlanIds"`
}
type TelnetConfig struct {
	TelnetEnabled bool `json:"telnetEnabled"`
}
type SSHConfig struct {
	SSHEnabled bool `json:"sshEnabled"`
}
type HTTPConfig struct {
	HTTPEnabled bool `json:"httpEnabled"`
}
type HTTPSConfig struct {
	HTTPSEnabled bool `json:"httpsEnabled"`
}
type GreConfig struct {
	ApplianceNetworkAddress string `json:"applianceNetworkAddress"`
	CoreFlow2NetworkAddress string `json:"coreFlow2NetworkAddress"`
}
type TapModeConfig struct {
	TapInterface string `json:"tapInterface"`
}
type TunnelModeConfig struct {
	TunnelInterface           string `json:"tunnelInterface"`
	TunnelNetworkAddress      string `json:"tunnelNetworkAddress"`
	TunnelNetworkGateway      string `json:"tunnelNetworkGateway"`
	TunnelNetworkPrefixLength int    `json:"tunnelNetworkPrefixLength"`
}
type PurviewConfig struct {
	NisEnabled        bool               `json:"nisEnabled"`
	NisServerAddress  string             `json:"nisServerAddress"`
	NisDomainName     string             `json:"nisDomainName"`
	GreEnabled        bool               `json:"greEnabled"`
	TapModeEnabled    bool               `json:"tapModeEnabled"`
	TunnelModeEnabled bool               `json:"tunnelModeEnabled"`
	GreConfig         []GreConfig        `json:"greConfig"`
	TapModeConfig     []TapModeConfig    `json:"tapModeConfig"`
	TunnelModeConfig  []TunnelModeConfig `json:"tunnelModeConfig"`
}
type NacConfig struct {
	ApplianceGroup string `json:"applianceGroup"`
}
type Appliance struct {
	PurviewConfig PurviewConfig `json:"purviewConfig"`
	NacConfig     NacConfig     `json:"nacConfig"`
}
type MlagGroups struct {
	MlagDetect bool     `json:"mlagDetect"`
	MlagPeer   string   `json:"mlagPeer"`
	MlagPorts  []string `json:"mlagPorts"`
}
type MLag struct {
	Enabled    bool         `json:"enabled"`
	MlagGroups []MlagGroups `json:"mlagGroups"`
}
type MlagPeers struct {
	MlagPeerID     string `json:"mlagPeerId"`
	MlagPeerIPAddr string `json:"mlagPeerIpAddr"`
	MlagpeerVr     string `json:"mlagpeerVr"`
	MlagPeerVlanID int    `json:"mlagPeerVlanId"`
}
type MlagMembers struct {
	MlagPort       string `json:"mlagPort"`
	MlagPortMlagID int    `json:"mlagPortMlagId"`
	MlagPortPeerID string `json:"mlagPortPeerId"`
}
type MLagv2 struct {
	Logging     string        `json:"logging"`
	MlagPeers   []MlagPeers   `json:"mlagPeers"`
	MlagMembers []MlagMembers `json:"mlagMembers"`
}
type Peers struct {
	StackPeer string `json:"stackPeer"`
}
type Stack struct {
	Enabled  bool     `json:"enabled"`
	Logging  string   `json:"logging"`
	Topology string   `json:"topology"`
	Ports    []string `json:"ports"`
	Peers    Peers    `json:"peers"`
}
type MvrpConfig struct {
	MvrpEnable bool     `json:"mvrpEnable"`
	MvrpPorts  []string `json:"mvrpPorts"`
}
type Ospf struct {
	Enabled bool   `json:"enabled"`
	Logging string `json:"logging"`
}
type ConfigBlock struct {
	License      License      `json:"license"`
	Poe          Poe          `json:"poe"`
	Dot1X        Dot1X        `json:"dot1X"`
	Lacp         Lacp         `json:"lacp"`
	Lldp         Lldp         `json:"lldp"`
	Logins       Logins       `json:"logins"`
	MacAuth      MacAuth      `json:"macAuth"`
	Ports        []Ports      `json:"ports"`
	Snmp         Snmp         `json:"snmp"`
	SpanningTree SpanningTree `json:"spanningTree"`
	Syslog       Syslog       `json:"syslog"`
	Vlans        Vlans        `json:"vlans"`
	MgmtAccess   MgmtAccess   `json:"mgmtAccess"`
	Appliance    Appliance    `json:"appliance"`
	MLag         MLag         `json:"mLag"`
	MLagv2       MLagv2       `json:"mLagv2"`
	Stack        Stack        `json:"stack"`
	Mvrp         Mvrp         `json:"mvrp"`
	Ospf         Ospf         `json:"ospf"`
	Timestamp    string       `json:"timestamp"`
	BpRequestID  int          `json:"bpRequestId"`
	Status       string       `json:"status"`
}

type Event struct {
	Type        int    `json:"type"`
	Severity    string `json:"severity"`
	Timestamp   string `json:"timestamp"`
	Target      string `json:"target"`
	Description string `json:"description"`
}

type StackPeerConfig struct {
	StackPeer string `json:"stackPeer"`
}
type StackingPeer struct {
	Enabled         bool              `json:"enabled"`
	StackPeerConfig []StackPeerConfig `json:"stackPeerConfig"`
	StackPortConfig []string          `json:"stackPortConfig"`
	StackTopology   string            `json:"stackTopology"`
}
type Authentication struct {
	Type     string `json:"type"`
	Password string `json:"password"`
}
type DhcpRelayAgents struct {
	IPAddress string `json:"ipAddress"`
}
type StaticRoutes struct {
	Network   string `json:"network"`
	NextHop   string `json:"nextHop"`
	IPVersion string `json:"ipVersion"`
}
type Hosts struct {
	Vrf          string         `json:"vrf"`
	IPAddress    string         `json:"ipAddress"`
	IPVersion    string         `json:"ipVersion"`
	StaticRoutes []StaticRoutes `json:"staticRoutes"`
}
type ExtremeFabric struct {
	Enabled         bool              `json:"enabled"`
	FabricDomain    string            `json:"fabricDomain"`
	Authentication  Authentication    `json:"authentication"`
	DhcpRelayAgents []DhcpRelayAgents `json:"dhcpRelayAgents"`
	Hosts           []Hosts           `json:"hosts"`
}
type ConfigACK struct {
	License       License         `json:"license"`
	Poe           Poe             `json:"poe"`
	Dot1X         Dot1X           `json:"dot1X"`
	Lacp          Lacp            `json:"lacp"`
	Lldp          Lldp            `json:"lldp"`
	Logins        Logins          `json:"logins"`
	MacAuth       MacAuth         `json:"macAuth"`
	Ports         []Ports         `json:"ports"`
	Snmp          Snmp            `json:"snmp"`
	SpanningTree  SpanningTree    `json:"spanningTree"`
	Syslog        Syslog          `json:"syslog"`
	Vlans         Vlans           `json:"vlans"`
	MgmtAccess    MgmtAccess      `json:"mgmtAccess"`
	Mlag          Mlag            `json:"mlag"`
	Stacking      StackingPeer    `json:"stacking"`
	Mvrp          Mvrp            `json:"mvrp"`
	ExtremeFabric ExtremeFabric   `json:"extremeFabric"`
	Vxlan         Vxlan           `json:"vxlan"`
	RadiusServers []RadiusServers `json:"radiusServers"`
	Timestamp     string          `json:"timestamp"`
	BpRequestID   int             `json:"bpRequestId"`
	Status        string          `json:"status"`
}

type UpgradeAssets struct {
	AssetName    string `json:"assetName"`
	AssetVersion string `json:"assetVersion"`
	AssetType    string `json:"assetType"`
}
type Fans struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Rpm   int    `json:"rpm"`
}
type Trays struct {
	State string `json:"state"`
	Name  string `json:"name"`
	Fans  []Fans `json:"fans"`
}
type Cooling struct {
	Trays []Trays `json:"trays"`
}
type StpStatus struct {
	StpState string `json:"stpState"`
	StpRole  string `json:"stpRole"`
}
type Med struct {
	ModelName    string   `json:"modelName"`
	SwVersion    string   `json:"swVersion"`
	Capabilities []string `json:"capabilities"`
	Serial       string   `json:"serial"`
	HwVersion    string   `json:"hwVersion"`
	MdiPdValue   string   `json:"mdiPdValue"`
}
type Lag struct {
	Enabled   bool   `json:"enabled"`
	Lag       string `json:"lag"`
	RxState   string `json:"rxState"`
	SelLogic  string `json:"selLogic"`
	MuxState  string `json:"muxState"`
	RxPackets int    `json:"rxPackets"`
	TxPackets int    `json:"txPackets"`
}
type Link struct {
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
}
type PowerSupply struct {
	Module   string `json:"module"`
	State    string `json:"state"`
	PartInfo string `json:"partInfo"`
	Input    int    `json:"input"`
}
type Temperatue struct {
	Temp   int    `json:"temp"`
	Status string `json:"status"`
	Name   string `json:"name"`
}
type Repsonse struct {
}
type VpnTargets struct {
	RtValue string `json:"rt-value"`
	RtType  string `json:"rt-type"`
}
type VxlanEvpn struct {
	RouteDistinguisher string       `json:"route-distinguisher"`
	VpnTargets         []VpnTargets `json:"vpn-targets"`
}
type VxlanInstance struct {
	VxlanID           int       `json:"vxlan-id"`
	Description       string    `json:"description"`
	UnknowUnicastDrop string    `json:"unknow-unicast-drop"`
	FilterVrrp        string    `json:"filter-vrrp"`
	VxlanEvpn         VxlanEvpn `json:"vxlan-evpn"`
}
type AddressFamily struct {
	Af                  string `json:"af"`
	TunnelSourceIP      string `json:"tunnel-source-ip"`
	TunnelDestinationIP string `json:"tunnel-destination-ip"`
	BindVxlanID         []int  `json:"bind-vxlan-id"`
}
type StaticVxlanTunnel struct {
	VxlanTunnelID   int             `json:"vxlan-tunnel-id"`
	VxlanTunnelName string          `json:"vxlan-tunnel-name"`
	AddressFamily   []AddressFamily `json:"address-family"`
}
type RedundancyGroupBind struct {
	VxlanID         int `json:"vxlan-id"`
	RedundancyGroup int `json:"redundancy-group"`
}
type VtepInstances struct {
	VtepID                int                   `json:"vtep-id"`
	VtepName              string                `json:"vtep-name"`
	SourceInterface       string                `json:"source-interface"`
	MulticastIP           string                `json:"multicast-ip"`
	InnerVlanHandlingMode int                   `json:"inner-vlan-handling-mode"`
	BindVxlanID           []int                 `json:"bind-vxlan-id"`
	StaticVxlanTunnel     []StaticVxlanTunnel   `json:"static-vxlan-tunnel"`
	RedundancyGroupBind   []RedundancyGroupBind `json:"redundancy-group-bind"`
}
type IetfVxlanVxlan struct {
	GlobalEnable  bool            `json:"global-enable"`
	VxlanInstance []VxlanInstance `json:"vxlan-instance"`
	VtepInstances []VtepInstances `json:"vtep-instances"`
}
type TunnelVniStatistic struct {
	VxlanID    int `json:"vxlan-id"`
	InBytes    int `json:"in-bytes"`
	OutBytes   int `json:"out-bytes"`
	InPackets  int `json:"in-packets"`
	OutPackets int `json:"out-packets"`
}
type Statistics struct {
	InBytes            int                  `json:"in-bytes"`
	OutBytes           int                  `json:"out-bytes"`
	InPackets          int                  `json:"in-packets"`
	OutPackets         int                  `json:"out-packets"`
	TunnelVniStatistic []TunnelVniStatistic `json:"tunnel-vni-statistic"`
}
type VxlanTunnel struct {
	LocalIP        string     `json:"local-ip"`
	RemoteIP       string     `json:"remote-ip"`
	StaticTunnelID int        `json:"static-tunnel-id"`
	EvpnTunnelID   int        `json:"evpn-tunnel-id"`
	Statistics     Statistics `json:"statistics"`
}
type IetfVxlanVxlanState struct {
	VxlanTunnel []VxlanTunnel `json:"vxlan-tunnel"`
}
type CPUMonitorTable struct {
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
}
type CPUMonitorSystemTable struct {
	Module            string `json:"module"`
	Utilization5Secs  int    `json:"utilization5secs"`
	Utilization10Secs int    `json:"utilization10secs"`
	Utilization30Secs int    `json:"utilization30secs"`
	Utilization1Min   int    `json:"utilization1min"`
	Utilization5Mins  int    `json:"utilization5mins"`
	Utilization30Mins int    `json:"utilization30mins"`
	Utilization1Hour  int    `json:"utilization1hour"`
	MaxUtilization    int    `json:"maxUtilization"`
}
type CPUUtilization struct {
	CPUMonitorInterval         int                     `json:"cpuMonitorInterval"`
	CPUMonitorTotalUtilization int                     `json:"cpuMonitorTotalUtilization"`
	CPUMonitorTable            []CPUMonitorTable       `json:"cpuMonitorTable"`
	CPUMonitorSystemTable      []CPUMonitorSystemTable `json:"cpuMonitorSystemTable"`
}
type MemoryMonitorTable struct {
	Module      string `json:"module"`
	ProcessName string `json:"processName"`
	Usage       int    `json:"usage"`
	Limit       int    `json:"limit"`
}
type MemoryMonitorSystemTable struct {
	Module      string `json:"module"`
	SystemTotal int    `json:"systemTotal"`
	SystemFree  int    `json:"systemFree"`
	SystemUsage int    `json:"systemUsage"`
	UserUsage   int    `json:"userUsage"`
}
type MemoryUtilization struct {
	MemoryMonitorTable       []MemoryMonitorTable       `json:"memoryMonitorTable"`
	MemoryMonitorSystemTable []MemoryMonitorSystemTable `json:"memoryMonitorSystemTable"`
}
type MLagPeers struct {
	MLagPeerID     string `json:"mLagPeerId"`
	MLagPeerStatus string `json:"mLagPeerStatus"`
	MLagPeerUpTime int    `json:"mLagPeerUpTime"`
}

// C O N N E C T
type Connect struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock"`
	DeviceInfo      DeviceInfo      `json:"deviceInfo"`
}

type ConnectResponse struct {
	Action        string `json:"action"`
	StandyTimeout int    `json:"standyTimeout"`
	Credentials   []struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	} `json:"credentials"`
	Redirect struct {
		Type string `json:"type"`
		URI  string `json:"uri"`
	} `json:"redirect"`
}

// I M A G E U P G R A D E
type ImageUpgrade struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock"`
	Assets          []Assets        `json:"assets"`
}

// C O N F I G U R A T I O N
type Configuration struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock"`
	Capabilities    Capabilities    `json:"capabilities"`
	ConfigBlock     ConfigBlock     `json:"configBlock"`
}

// E V E N T
type EventMsg struct {
	ApPropertyBlock ApPropertyBlock `json:"apPropertyBlock"`
	Event           []Event         `json:"event"`
	ConfigACK       ConfigACK       `json:"ConfigACK"`
}

// S T A T S
type Stats struct {
	ApPropertyBlock     ApPropertyBlock     `json:"apPropertyBlock"`
	Capabilities        Capabilities        `json:"capabilities"`
	Assets              []Assets            `json:"assets"`
	UpgradeAssets       []UpgradeAssets     `json:"upgradeAssets"`
	ConfigBlock         ConfigBlock         `json:"configBlock"`
	Cooling             Cooling             `json:"cooling"`
	IfTable             []IfTable           `json:"ifTable"`
	PortsInfo           []PortsInfo         `json:"portsInfo"`
	PowerSupply         []PowerSupply       `json:"powerSupply"`
	Temperatue          []Temperatue        `json:"temperatue"`
	Telemetry           []Telemetry         `json:"telemetry"`
	IetfVxlanVxlan      IetfVxlanVxlan      `json:"ietf-vxlan:vxlan"`
	IetfVxlanVxlanState IetfVxlanVxlanState `json:"ietf-vxlan:vxlan-state"`
	CPUUtilization      CPUUtilization      `json:"cpuUtilization"`
	MemoryUtilization   MemoryUtilization   `json:"memoryUtilization"`
	Mlagv2              Mlagv2              `json:"mlagv2"`
	Timestamp           int                 `json:"timestamp"`
	SysUpTime           int                 `json:"sysUpTime"`
	CheckInTime         string              `json:"checkInTime"`
	UpgradeTime         string              `json:"upgradeTime"`
}
