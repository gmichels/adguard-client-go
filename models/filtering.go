package models

// FilterStatus - Filtering settings
type FilterStatus struct {
	Enabled          bool     `json:"enabled"`
	Interval         uint     `json:"interval"`
	Filters          []Filter `json:"filters"`
	WhitelistFilters []Filter `json:"whitelist_filters"`
	UserRules        []string `json:"user_rules"`
}

// Filter - Filter subscription info
type Filter struct {
	Enabled     bool   `json:"enabled"`
	Id          int64  `json:"id"`
	LastUpdated string `json:"last_updated"`
	Name        string `json:"name"`
	RulesCount  uint32 `json:"rules_count"`
	Url         string `json:"url"`
}

// FilterConfig - Filtering settings
type FilterConfig struct {
	Enabled  bool `json:"enabled"`
	Interval uint `json:"interval"`
}

// AddUrlRequest - /add_url request data
type AddUrlRequest struct {
	Name      string `json:"name"`
	Url       string `json:"url" description:"URL or an absolute path to the file containing filtering rules."`
	Whitelist bool   `json:"whitelist"`
}

// RemoveUrlRequest - /remove_url request data
type RemoveUrlRequest struct {
	Url       string `json:"url" description:"Previously added URL containing filtering rules"`
	Whitelist bool   `json:"whitelist"`
}

// FilterSetUrl - Filtering URL settings
type FilterSetUrl struct {
	Data      FilterSetUrlData `json:"data"`
	Url       string           `json:"url"`
	Whitelist bool             `json:"whitelist"`
}

// FilterSetUrlData - Filter update data
type FilterSetUrlData struct {
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
	Url     string `json:"url"`
}

// FilterRefreshRequest - Refresh Filters request data
type FilterRefreshRequest struct {
	Whitelist bool `json:"whitelist"`
}

// FilterRefreshResponse - /filtering/refresh response data
type FilterRefreshResponse struct {
	Updated int `json:"updated"`
}

// SetRulesRequest - Custom filtering rules setting request.
type SetRulesRequest struct {
	Rules []string `json:"rules"`
}

// FilterCheckHostResponse - Check Host Result
type FilterCheckHostResponse struct {
	Reason      string       `json:"reason" description:"Request filtering status."`
	Rules       []ResultRule `json:"rules" description:"Applied rules."`
	ServiceName string       `json:"service_name" description:"Set if reason=FilteredBlockedService"`
	Cname       string       `json:"cname" description:"Set if reason=Rewrite"`
	IpAddrs     []string     `json:"ip_addrs" description:"Set if reason=Rewrite"`
}

// ResultRule - Applied rule.
type ResultRule struct {
	FilterListId int64  `json:"filter_list_id" description:"In case if there's a rule applied to this DNS request, this is ID of the filter list that the rule belongs to."`
	Text         string `json:"text" description:"The text of the filtering rule applied to the request (if any)."`
}
