package models

// QueryLog - Query log
type QueryLog struct {
	Oldest string         `json:"oldest"`
	Data   []QueryLogItem `json:"data"`
}

// QueryLogItem - Query log item
type QueryLogItem struct {
	Answer         []DnsAnswer        `json:"answer"`
	OriginalAnswer []DnsAnswer        `json:"original_answer" description:"Answer from upstream server (optional)"`
	Cached         bool               `json:"cached" description:"Defines if the response has been served from cache"`
	Upstream       string             `json:"upstream" description:"Upstream URL starting with tcp://, tls://, https://, or with an IP address."`
	AnswerDnssec   bool               `json:"answer_dnssec" description:"If true, the response had the Authenticated Data (AD) flag set."`
	Client         string             `json:"client" description:"The client's IP address."`
	ClientId       string             `json:"client_id" description:"The ClientID, if provided in DoH, DoQ, or DoT."`
	ClientInfo     QueryLogItemClient `json:"client_info"`
	ClientProto    string             `json:"client_proto"`
	Ecs            string             `json:"ecs" description:"The IP network defined by an EDNS Client-Subnet option in the request message if any."`
	ElapsedMs      string             `json:"elapsedMs"`
	Question       DnsQuestion        `json:"question"`
	Rules          []ResultRule       `json:"rules" description:"Applied rules."`
	Reason         string             `json:"reason" description:"Request filtering status."`
	ServiceName    string             `json:"service_name" description:"Set if reason=FilteredBlockedService"`
	Status         string             `json:"status" description:"DNS response status"`
	Time           string             `json:"time" description:"DNS request processing start time"`
}

// DnsAnswer - DNS answer section
type DnsAnswer struct {
	Ttl   uint32 `json:"ttl"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

// QueryLogItemClient - Client information for a query log item.
type QueryLogItemClient struct {
	Disallowed     bool                    `json:"disallowed" description:"Whether the client's IP is blocked or not."`
	DisallowedRule string                  `json:"disallowed_rule" description:"The rule due to which the client is allowed or blocked."`
	Name           string                  `json:"name" description:"Persistent client's name or runtime client's hostname. May be empty."`
	Whois          QueryLogItemClientWhois `json:"whois"`
}

// QueryLogItemClientWhois - Client WHOIS information, if any.
type QueryLogItemClientWhois struct {
	City    string `json:"city" description:"City, if any."`
	Country string `json:"country" description:"Country, if any."`
	Orgname string `json:"orgname" description:"Organization name, if any."`
}

// DnsQuestion - DNS question section
type DnsQuestion struct {
	Class       string `json:"class"`
	Name        string `json:"name"`
	UnicodeName string `json:"unicode_name"`
	Type        string `json:"type"`
}

// GetQueryLogConfigResponse - Query log configuration
type GetQueryLogConfigResponse struct {
	Enabled           bool     `json:"enabled" description:"Is query log enabled"`
	Interval          uint64   `json:"interval" description:"Time period for query log rotation in milliseconds"`
	AnonymizeClientIp bool     `json:"anonymize_client_ip" description:"Anonymize clients' IP address"`
	Ignored           []string `json:"ignored" description:"List of host names, which should not be written to log"`
}
