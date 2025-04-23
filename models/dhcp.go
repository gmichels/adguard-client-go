package models

// DhcpStatus - Built-in DHCP server configuration and status
type DhcpStatus struct {
	Enabled       bool              `json:"enabled"`
	InterfaceName string            `json:"interface_name"`
	V4            DhcpConfigV4      `json:"v4"`
	V6            DhcpConfigV6      `json:"v6"`
	Leases        []DhcpLease       `json:"leases"`
	StaticLeases  []DhcpStaticLease `json:"static_leases"`
}

// DhcpConfigV4
type DhcpConfigV4 struct {
	GatewayIp     string `json:"gateway_ip,omitempty"`
	SubnetMask    string `json:"subnet_mask,omitempty"`
	RangeStart    string `json:"range_start,omitempty"`
	RangeEnd      string `json:"range_end,omitempty"`
	LeaseDuration uint64 `json:"lease_duration,omitempty"`
}

// DhcpConfigV6
type DhcpConfigV6 struct {
	RangeStart    string `json:"range_start,omitempty"`
	LeaseDuration uint64 `json:"lease_duration,omitempty"`
}

// DhcpLease
type DhcpLease struct {
	Mac      string `json:"mac"`
	Ip       string `json:"ip"`
	Hostname string `json:"hostname"`
	Expires  string `json:"expires"`
}

// DhcpStaticLease - DHCP static lease information
type DhcpStaticLease struct {
	Mac      string `json:"mac"`
	Ip       string `json:"ip"`
	Hostname string `json:"hostname"`
}

// NetInterfaces - Network interfaces dictionary, keys are interface names.
type NetInterfaces map[string]NetInterface

// NetInterface - Network interface info
type NetInterface struct {
	Flags           string   `json:"flags" description:"Flags could be any combination of the following values, divided by the \"|\" character: \"up\", \"broadcast\", \"loopback\", \"pointtopoint\" and \"multicast\"."`
	GatewayIp       string   `json:"gateway_ip" description:"The IP address of the gateway."`
	HardwareAddress string   `json:"hardware_address"`
	Name            string   `json:"name"`
	IpV4Addresses   []string `json:"ipv4_addresses" description:"The addresses of the interface of v4 family."`
	IpV6Addresses   []string `json:"ipv6_addresses" description:"The addresses of the interface of v6 family."`
}

// DhcpConfig
type DhcpConfig struct {
	Enabled       bool         `json:"enabled"`
	InterfaceName string       `json:"interface_name"`
	V4            DhcpConfigV4 `json:"v4,omitempty"`
	V6            DhcpConfigV6 `json:"v6,omitempty"`
}

// DhcpFindActiveReq - Request for checking for other DHCP servers in the network.
type DhcpFindActiveReq struct {
	Interface string `json:"interface" description:"The name of the network interface"`
}

// DhcpSearchResult - Information about a DHCP server discovered in the current network.
type DhcpSearchResult struct {
	V4 DhcpSearchV4 `json:"v4"`
	V6 DhcpSearchV6 `json:"v6"`
}

// DhcpSearchV4
type DhcpSearchV4 struct {
	OtherServer DhcpSearchResultOtherServer `json:"other_server"`
	StaticIp    DhcpSearchResultStaticIp    `json:"static_ip"`
}

// DhcpSearchV6
type DhcpSearchV6 struct {
	OtherServer DhcpSearchResultOtherServer `json:"other_server"`
}

// DhcpSearchResultOtherServer
type DhcpSearchResultOtherServer struct {
	Found string `json:"found" description:"The result of searching the other DHCP server."`
	Error string `json:"error"`
}

// DhcpSearchResultStaticIp
type DhcpSearchResultStaticIp struct {
	Static string `json:"static" description:"The result of determining static IP address."`
	Ip     string `json:"ip"`
}
