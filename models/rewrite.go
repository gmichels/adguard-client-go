package models

// RewriteList - Rewrite rules array
type RewriteList []RewriteEntry

// RewriteEntry - Rewrite rule
type RewriteEntry struct {
	Domain  string `json:"domain" description:"Domain name"`
	Answer  string `json:"answer" description:"value of A, AAAA or CNAME DNS record"`
	Enabled bool   `json:"enabled,omitempty" description:"Optional. If omitted on add, defaults to true. On update, omitted preserves previous value."`
}

// RewriteUpdate - Rewrite rule update object
type RewriteUpdate struct {
	Target RewriteEntry `json:"target"`
	Update RewriteEntry `json:"update"`
}

// RewriteSettings - DNS rewrite settings
type RewriteSettings struct {
	Enabled bool `json:"enabled" description:"indicates whether rewrites are applied"`
}
