package adguard

// AllClients
type AllClients struct {
	Clients       []Client     `json:"clients"`
	ClientAuto    []ClientAuto `json:"auto_clients"`
	SupportedTags []string     `json:"supported_tags"`
}

// Client
type Client struct {
	Name                string   `json:"name"`
	Ids                 []string `json:"ids"`
	UseGlobalSettings   bool     `json:"use_global_settings"`
	FilteringEnabled    bool     `json:"filtering_enabled"`
	ParentalEnabled     bool     `json:"parental_enabled"`
	SafebrowsingEnabled bool     `json:"safebrowsing_enabled"`
	BlockedServices     []string `json:"blocked_services"`
	Upstreams           []string `json:"upstreams"`
	Tags                []string `json:"tags"`
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
