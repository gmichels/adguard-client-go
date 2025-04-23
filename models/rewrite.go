package models

// RewriteList - Rewrite rules array
type RewriteList []RewriteEntry

// RewriteEntry - Rewrite rule
type RewriteEntry struct {
	Domain string `json:"domain" description:"Domain name"`
	Answer string `json:"answer" description:"value of A, AAAA or CNAME DNS record"`
}

// RewriteUpdate - Rewrite rule update object
type RewriteUpdate struct {
	Target RewriteEntry `json:"target"`
	Update RewriteEntry `json:"update"`
}
