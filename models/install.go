package models

// AddressesInfo - AdGuard Home addresses configuration
type AddressesInfo struct {
	DnsPort    uint16                   `json:"dns_port"`
	Interfaces map[string]NetInterfaces `json:"interfaces"`
	Version    string                   `json:"version"`
	WebPort    uint16                   `json:"web_port"`
}

// CheckConfigRequest - Configuration to be checked
type CheckConfigRequest struct {
	Dns         CheckConfigRequestInfo `json:"dns"`
	Web         CheckConfigRequestInfo `json:"web"`
	SetStaticIp bool                   `json:"set_static_ip"`
}

// CheckConfigRequestInfo
type CheckConfigRequestInfo struct {
	Ip      string `json:"ip"`
	Port    uint16 `json:"port"`
	Autofix bool   `json:"autofix"`
}

// CheckConfigResponse
type CheckConfigResponse struct {
	Dns         CheckConfigResponseInfo `json:"dns"`
	Web         CheckConfigResponseInfo `json:"web"`
	SetStaticIp CheckConfigStaticIpInfo `json:"set_static_ip"`
}

// CheckConfigResponseInfo
type CheckConfigResponseInfo struct {
	Status     string `json:"status"`
	CanAutofix bool   `json:"can_autofix"`
}

// CheckConfigStaticIpInfo
type CheckConfigStaticIpInfo struct {
	Static CheckConfigStaticIpInfoStatic `json:"static"`
	Ip     string                        `json:"ip" description:"Current dynamic IP address. Set if static=no"`
	Error  string                        `json:"error" description:"Error text. Set if static=no"`
}

// CheckConfigStaticIpInfoStatic
type CheckConfigStaticIpInfoStatic string

// InitialConfiguration - AdGuard Home initial configuration for the first-install wizard.
type InitialConfiguration struct {
	Dns      AddressInfo `json:"dns"`
	Web      AddressInfo `json:"web"`
	Username string      `json:"username" description:"Basic auth username"`
	Password string      `json:"password" description:"Basic auth password"`
}

// AddressInfo - Port information
type AddressInfo struct {
	Ip   string `json:"ip"`
	Port uint16 `json:"port"`
}
