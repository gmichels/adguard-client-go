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
	Interval         int      `json:"interval"`
	Filters          []Filter `json:"filters"`
	WhitelistFilters []Filter `json:"whitelist_filters"`
	UserRules        []string `json:"user_rules"`
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
