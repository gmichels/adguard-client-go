package models

// AccessList - Client and host access list. Each of the lists should contain only unique elements. In addition, allowed and disallowed lists cannot contain the same elements.
type AccessList struct {
	AllowedClients    []string `json:"allowed_clients" description:"The allowlist of clients: IP addresses, CIDRs, or ClientIDs."`
	DisallowedClients []string `json:"disallowed_clients" description:"The blocklist of clients: IP addresses, CIDRs, or ClientIDs."`
	BlockedHosts      []string `json:"blocked_hosts" description:"The blocklist of hosts."`
}
