package models

// Clients
type Clients struct {
	Clients       ClientsArray     `json:"clients"`
	ClientAuto    ClientsAutoArray `json:"auto_clients"`
	SupportedTags []string         `json:"supported_tags"`
}

// ClientsArray - Clients array
type ClientsArray []Client

// ClientsAutoArray - Auto-Clients array
type ClientsAutoArray []ClientAuto

// ClientAuto - Auto-Client information
type ClientAuto struct {
	Name      string    `json:"name" description:"Name"`
	Ip        string    `json:"ip" description:"IP address"`
	Mac       string    `json:"mac" description:"MAC address"`
	Source    string    `json:"source" description:"The source of this information"`
	WhoisInfo WhoisInfo `json:"whois_info"`
}

type WhoisInfo map[string]string

// Client - Client information
type Client struct {
	Name                string   `json:"name" description:"Name"`
	Ids                 []string `json:"ids" description:"IP, CIDR, MAC, or ClientID."`
	UseGlobalSettings   bool     `json:"use_global_settings"`
	FilteringEnabled    bool     `json:"filtering_enabled"`
	ParentalEnabled     bool     `json:"parental_enabled"`
	SafebrowsingEnabled bool     `json:"safebrowsing_enabled"`
	// Deprecated: SafesearchEnabled, use SafeSearch instead
	SafesearchEnabled        bool             `json:"safesearch_enabled"`
	SafeSearch               SafeSearchConfig `json:"safe_search"`
	UseGlobalBlockedServices bool             `json:"use_global_blocked_services"`
	BlockedServicesSchedule  Schedule         `json:"blocked_services_schedule"`
	BlockedServices          []string         `json:"blocked_services"`
	Upstreams                []string         `json:"upstreams"`
	Tags                     []string         `json:"tags"`
	IgnoreQuerylog           bool             `json:"ignore_querylog"`
	IgnoreStatistics         bool             `json:"ignore_statistics"`
	UpstreamsCacheEnabled    bool             `json:"upstreams_cache_enabled"`
	UpstreamsCacheSize       uint             `json:"upstreams_cache_size"`
}

// ClientDelete - Client delete request
type ClientDelete struct {
	Name string `json:"name"`
}

// ClientUpdate - Client update request
type ClientUpdate struct {
	Name string `json:"name"`
	Data Client `json:"data"`
}

// ClientsSearchRequest - Client search request
type ClientsSearchRequest struct {
	Clients []ClientSearchRequestItem `json:"clients"`
}

// ClientSearchRequestItem
type ClientSearchRequestItem struct {
	Id string `json:"id" description:"Client IP address, CIDR, MAC address, or ClientID"`
}

// ClientsFindResponse - Client search results
type ClientsFindResponse []ClientFindEntry

// ClientFindEntry
type ClientFindEntry map[string]ClientFindSubEntry

// ClientFindSubEntry - Client information
type ClientFindSubEntry struct {
	Name                string   `json:"name" description:"Name"`
	Ids                 []string `json:"ids" decription:"IP, CIDR, MAC, or ClientID."`
	UseGlobalSettings   bool     `json:"use_global_settings"`
	FilteringEnabled    bool     `json:"filtering_enabled"`
	ParentalEnabled     bool     `json:"parental_enabled"`
	SafebrowsingEnabled bool     `json:"safebrowsing_enabled"`
	// Deprecated: SafesearchEnabled, use SafeSearch instead
	SafesearchEnabled        bool             `json:"safesearch_enabled"` // deprecated
	SafeSearch               SafeSearchConfig `json:"safe_search"`
	UseGlobalBlockedServices bool             `json:"use_global_blocked_services"`
	BlockedServices          []string         `json:"blocked_services"`
	Upstreams                []string         `json:"upstreams"`
	WhoisInfo                WhoisInfo        `json:"whois_info"`
	Disallowed               bool             `json:"disallowed" description:"Whether the client's IP is blocked or not."`
	DisallowedRule           string           `json:"disallowed_rule" description:"The rule due to which the client is disallowed. If disallowed is set to true, and this string is empty, then the client IP is disallowed by the \"allowed IP list\", that is it is not included in the allowed list."`
	IgnoreQuerylog           bool             `json:"ignore_querylog"`
	IgnoreStatistics         bool             `json:"ignore_statistics"`
}
