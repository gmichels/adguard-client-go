package models

// Stats - Server statistics data
type Stats struct {
	TimeUnits               string          `json:"time_units" description:"Time units"`
	NumDnsQueries           int             `json:"num_dns_queries" description:"Total number of DNS queries"`
	NumBlockedFiltering     int             `json:"num_blocked_filtering" description:"Number of requests blocked by filtering rules"`
	NumReplacedSafebrowsing int             `json:"num_replaced_safebrowsing" description:"Number of requests blocked by safebrowsing module"`
	NumReplaceSafesearch    int             `json:"num_replace_safesearch" description:"Number of requests blocked by safesearch module"`
	NumReplacedParental     int             `json:"num_replaced_parental" description:"Number of blocked adult websites"`
	AvgProcessingTime       float32         `json:"avg_processing_time" description:"Average time in seconds on processing a DNS request"`
	TopQueriedDomains       []TopArrayEntry `json:"top_queried_domains"`
	TopClients              []TopArrayEntry `json:"top_clients"`
	TopBlockedDomains       []TopArrayEntry `json:"top_blocked_domains"`
	TopUpstreamResponses    []TopArrayEntry `json:"top_upstream_responses" description:"Total number of responses from each upstream."`
	TopUpstreamAvgTime      []TopArrayEntry `json:"top_upstream_avg_time" description:"Average processing time in seconds of requests from each upstream."`
	DnsQueries              []int           `json:"dns_queries"`
	BlockedFiltering        []int           `json:"blocked_filtering"`
	ReplacedSafebrowsing    []int           `json:"replaced_safebrowsing"`
	ReplacedParental        []int           `json:"replaced_parental"`
}

// TopArrayEntry - Represent the number of hits or time duration per key (url, domain, or client IP).
type TopArrayEntry struct {
	DomainOrIp           float32 `json:"domain_or_ip"`
	AdditionalProperties map[string]float32
}

// GetStatsConfigResponse
type GetStatsConfigResponse struct {
	Enabled  bool     `json:"enabled"`
	Interval uint64   `json:"interval"`
	Ignored  []string `json:"ignored"`
}
