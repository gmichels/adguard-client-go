package adguard

// AllClients
type AllClients struct {
	Clients       []Client     `json:"clients"`
	ClientAuto    []ClientAuto `json:"auto_clients"`
	SupportedTags []string     `json:"supported_tags"`
}

// Client
type Client struct {
	Name                     string   `json:"name"`
	Ids                      []string `json:"ids"`
	UseGlobalSettings        bool     `json:"use_global_settings"`
	FilteringEnabled         bool     `json:"filtering_enabled"`
	ParentalEnabled          bool     `json:"parental_enabled"`
	SafebrowsingEnabled      bool     `json:"safebrowsing_enabled"`
	SafesearchEnabled        bool     `json:"safesearch_enabled"`
	UseGlobalBlockedServices bool     `json:"use_global_blocked_services"`
	BlockedServices          []string `json:"blocked_services"`
	Upstreams                []string `json:"upstreams"`
	Tags                     []string `json:"tags"`
}

// ClientAuto
type ClientAuto struct {
	Name      string            `json:"name"`
	Ip        string            `json:"ip"`
	Source    string            `json:"source"`
	WhoisInfo map[string]string `json:"whois_info"`
}

// ClientUpdate
type ClientUpdate struct {
	Name string `json:"name"`
	Data Client `json:"data"`
}

// ClientDelete
type ClientDelete struct {
	Name string `json:"name"`
}

// FilterStatus
type FilterStatus struct {
	Enabled          bool     `json:"enabled"`
	Interval         uint     `json:"interval"`
	Filters          []Filter `json:"filters"`
	WhitelistFilters []Filter `json:"whitelist_filters"`
	UserRules        []string `json:"user_rules"`
}

// FilterConfig
type FilterConfig struct {
	Enabled  bool `json:"enabled"`
	Interval uint `json:"interval"`
}

// Filter
type Filter struct {
	Enabled     bool   `json:"enabled"`
	Id          int64  `json:"id"`
	LastUpdated string `json:"last_updated"`
	Name        string `json:"name"`
	RulesCount  int32  `json:"rules_count"`
	Url         string `json:"url"`
}

// AddUrlRequest
type AddUrlRequest struct {
	Name      string `json:"name"`
	Url       string `json:"url"`
	Whitelist bool   `json:"whitelist"`
}

// FilterSetUrl
type FilterSetUrl struct {
	Data      FilterSetUrlData `json:"data"`
	Url       string           `json:"url"`
	Whitelist bool             `json:"whitelist"`
}

// FilterSetUrlData
type FilterSetUrlData struct {
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

// RemoveUrlRequest
type RemoveUrlRequest struct {
	Url       string `json:"url"`
	Whitelist bool   `json:"whitelist"`
}

// SetRulesRequest
type SetRulesRequest struct {
	Rules []string `json:"rules"`
}

// RewriteEntry
type RewriteEntry struct {
	Domain string `json:"domain"`
	Answer string `json:"answer"`
}

// DNSConfig
type DNSConfig struct {
	BootstrapDns           []string `json:"bootstrap_dns"`
	UpstreamDns            []string `json:"upstream_dns"`
	UpstreamDnsFile        string   `json:"upstream_dns_file"`
	RateLimit              uint     `json:"ratelimit"`
	BlockingMode           string   `json:"blocking_mode"`
	BlockingIpv4           string   `json:"blocking_ipv4"`
	BlockingIpv6           string   `json:"blocking_ipv6"`
	EDnsCsEnabled          bool     `json:"edns_cs_enabled"`
	DisableIpv6            bool     `json:"disable_ipv6"`
	DnsSecEnabled          bool     `json:"dnssec_enabled"`
	CacheSize              uint     `json:"cache_size"`
	CacheTtlMin            uint     `json:"cache_ttl_min"`
	CacheTtlMax            uint     `json:"cache_ttl_max"`
	CacheOptimistic        bool     `json:"cache_optimistic"`
	UpstreamMode           string   `json:"upstream_mode"`
	UsePrivatePtrResolvers bool     `json:"use_private_ptr_resolvers"`
	ResolveClients         bool     `json:"resolve_clients"`
	LocalPtrUpstreams      []string `json:"local_ptr_upstreams"`
}

// DNSInfo
type DNSInfo struct {
	*DNSConfig
	DefaultLocalPtrUpstreams []string `json:"default_local_ptr_upstreams"`
}

// AccessList
type AccessList struct {
	AllowedClients    []string `json:"allowed_clients"`
	DisallowedClients []string `json:"disallowed_clients"`
	BlockedHosts      []string `json:"blocked_hosts"`
}

// GetQueryLogConfigResponse
type GetQueryLogConfigResponse struct {
	Enabled           bool     `json:"enabled"`
	Interval          uint64   `json:"interval"`
	AnonymizeClientIp bool     `json:"anonymize_client_ip"`
	Ignored           []string `json:"ignored"`
}

// GetStatsConfigResponse
type GetStatsConfigResponse struct {
	Enabled  bool     `json:"enabled"`
	Interval uint64   `json:"interval"`
	Ignored  []string `json:"ignored"`
}

// Enabled
type Enabled struct {
	Enabled bool `json:"enabled"`
}

// SafeSearchConfig
type SafeSearchConfig struct {
	Enabled    bool `json:"enabled"`
	Bing       bool `json:"bing"`
	Duckduckgo bool `json:"duckduckgo"`
	Google     bool `json:"google"`
	Pixabay    bool `json:"pixabay"`
	Yandex     bool `json:"yandex"`
	Youtube    bool `json:"youtube"`
}

// BlockedService
type BlockedService struct {
	IconSvg string   `json:"icon_svg"`
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Rules   []string `json:"rules"`
}

// BlockedServicesAll
type BlockedServicesAll struct {
	BlockedServices []BlockedService `json:"blocked_services"`
}

// DhcpStatus
type DhcpStatus struct {
	Enabled       bool              `json:"enabled"`
	InterfaceName string            `json:"interface_name"`
	V4            DhcpConfigV4      `json:"v4"`
	V6            DhcpConfigV6      `json:"v6"`
	Leases        []DhcpLease       `json:"leases"`
	StaticLeases  []DhcpStaticLease `json:"static_leases"`
}

// DhcpConfig
type DhcpConfig struct {
	Enabled       bool         `json:"enabled"`
	InterfaceName string       `json:"interface_name"`
	V4            DhcpConfigV4 `json:"v4,omitempty"`
	V6            DhcpConfigV6 `json:"v6,omitempty"`
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

// DhcpStaticLease
type DhcpStaticLease struct {
	Mac      string `json:"mac"`
	Ip       string `json:"ip"`
	Hostname string `json:"hostname"`
}

// NetInterface
type NetInterface struct {
	Flags           string   `json:"flags"`
	HardwareAddress string   `json:"hardware_address"`
	Name            string   `json:"name"`
	IpAddresses     []string `json:"ip_addresses"`
	Mtu             int      `json:"mtu"`
}

// NetInterfaces
type NetInterfaces struct {
	Name NetInterface
}
