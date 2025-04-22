package models

// ServerStatus - AdGuard Home server status and configuration
type ServerStatus struct {
	DnsAddresses               []string `json:"dns_addresses"`
	DnsPort                    uint16   `json:"dns_port"`
	HttpPort                   uint16   `json:"http_port"`
	ProtectionEnabled          bool     `json:"protection_enabled"`
	ProtectionDisabledDuration int64    `json:"protection_disabled_duration,omitempty"`
	DhcpAvailable              bool     `json:"dhcp_available,omitempty"`
	Running                    bool     `json:"running"`
	Version                    string   `json:"version"`
	Language                   string   `json:"language"`
}

// DNSInfo - model does not formally exist in the upstream API
type DNSInfo struct {
	*DNSConfig
	DefaultLocalPtrUpstreams []string `json:"default_local_ptr_upstreams"`
}

// DNSConfig - DNS server configuration
type DNSConfig struct {
	BootstrapDns                 []string `json:"bootstrap_dns" description:"Bootstrap servers, port is optional after colon. Empty value will reset it to default values."`
	UpstreamDns                  []string `json:"upstream_dns" description:"Upstream servers, port is optional after colon. Empty value will reset it to default values."`
	FallbackDns                  []string `json:"fallback_dns" description:"List of fallback DNS servers used when upstream DNS servers are not responding. Empty value will clear the list."`
	UpstreamDnsFile              string   `json:"upstream_dns_file"`
	ProtectionEnabled            bool     `json:"protection_enabled"`
	RateLimit                    uint     `json:"ratelimit"`
	RateLimitSubnetSubnetLenIpv4 uint     `json:"ratelimit_subnet_len_ipv4" description:"Length of the subnet mask for IPv4 addresses."`
	RateLimitSubnetSubnetLenIpv6 uint     `json:"ratelimit_subnet_len_ipv6" description:"Length of the subnet mask for IPv6 addresses."`
	RateLimitWhitelist           []string `json:"ratelimit_whitelist" description:"List of IP addresses excluded from rate limiting."`
	BlockingMode                 string   `json:"blocking_mode"`
	BlockingIpv4                 string   `json:"blocking_ipv4"`
	BlockingIpv6                 string   `json:"blocking_ipv6"`
	BlockedResponseTtl           uint     `json:"blocked_response_ttl" description:"TTL for blocked responses."`
	ProtectionDisabledUntil      string   `json:"protection_disabled_until,omitempty" description:"Protection is pause until this time. Nullable."`
	EDnsCsEnabled                bool     `json:"edns_cs_enabled"`
	EDnsCsUseCustom              bool     `json:"edns_cs_use_custom"`
	EDnsCsCustomIp               string   `json:"edns_cs_custom_ip"`
	DisableIpv6                  bool     `json:"disable_ipv6"`
	DnsSecEnabled                bool     `json:"dnssec_enabled"`
	CacheSize                    uint     `json:"cache_size"`
	CacheTtlMin                  uint     `json:"cache_ttl_min"`
	CacheTtlMax                  uint     `json:"cache_ttl_max"`
	CacheOptimistic              bool     `json:"cache_optimistic"`
	UpstreamMode                 string   `json:"upstream_mode" description:"Upstream modes enumeration."`
	UsePrivatePtrResolvers       bool     `json:"use_private_ptr_resolvers"`
	ResolveClients               bool     `json:"resolve_clients"`
	LocalPtrUpstreams            []string `json:"local_ptr_upstreams" description:"Upstream servers, port is optional after colon. Empty value will reset it to default values."`
	UpstreamTimeout              uint     `json:"upstream_timeout" description:"The number of seconds to wait for a response from the upstream server"`
}

// SetProtectionRequest - Protection state configuration
type SetProtectionRequest struct {
	Enabled  bool   `json:"enabled"`
	Duration uint64 `json:"duration,omitempty" description:"Duration of a pause, in milliseconds. Enabled should be false."`
}

// UpstreamsConfig - Upstream configuration to be tested
type UpstreamsConfig struct {
	BootstrapDns    []string `json:"bootstrap_dns" description:"Bootstrap DNS servers, port is optional after colon."`
	UpstreamDns     []string `json:"upstream_dns" description:"Upstream DNS servers, port is optional after colon."`
	FallbackDns     []string `json:"fallback_dns,omitempty" description:"Fallback DNS servers, port is optional after colon."`
	PrivateUpstream []string `json:"private_upstream" description:"Local PTR resolvers, port is optional after colon."`
}

// UpstreamsConfigResponse - Upstreams configuration response
type UpstreamsConfigResponse map[string]string

// GetVersionRequest - /version.json request data
type GetVersionRequest struct {
	RecheckNow bool `json:"recheck_now,omitempty" description:"If false, server will check for a new version data only once in several hours."`
}

// VersionInfo - Information about the latest available version of AdGuard Home
type VersionInfo struct {
	Disabled        bool   `json:"disabled" description:"If true then other fields don't appear."`
	NewVersion      string `json:"new_version,omitempty"`
	Announcement    string `json:"announcement,omitempty"`
	AnnouncementUrl string `json:"announcement_url,omitempty"`
	CanAutoupdate   bool   `json:"can_autoupdate,omitempty"`
}

// Login - Login request data
type Login struct {
	Name     string `json:"name" description:"User name"`
	Password string `json:"password" description:"Password"`
}

// ProfileInfo - Information about the current user
type ProfileInfo struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Theme    string `json:"theme" description:"Interface theme"`
}
