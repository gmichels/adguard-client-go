package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetAllFilter - Returns all filters
func (c *ADG) GetAllFilters() (*FilterStatus, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/filtering/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to a FilterStatus object
	var allFilters FilterStatus
	err = json.Unmarshal(body, &allFilters)
	if err != nil {
		return nil, err
	}

	return &allFilters, nil
}

// GetUserRules - Returns a list of all user rules
func (c *ADG) GetUserRules() (*[]string, error) {
	allFilters, err := c.GetAllFilters()
	if err != nil {
		return nil, err
	}

	return &allFilters.UserRules, nil
}

// UpdateUserRules - Update user-provided rules, returning the list of all user rules
func (c *ADG) UpdateUserRules(rules SetRulesRequest) (*[]string, error) {
	// convert provided filter to JSON
	rb, err := json.Marshal(rules)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/set_rules", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// appease Go as returned body is useless
	_ = body

	// return the list of rules that was passed
	return &rules.Rules, nil
}

// GetListFilterById - Returns a list filter based on its name and whether it's a whitelist filter
func (c *ADG) GetListFilterById(id int64) (*Filter, bool, error) {
	allFilters, err := c.GetAllFilters()
	if err != nil {
		return nil, false, err
	}

	// go through the blacklist filters in the response until we find the one we want
	for _, filterInfo := range allFilters.Filters {
		if filterInfo.Id == id {
			return &filterInfo, false, nil
		}
	}
	// go through the whitelist filters in the response until we find the one we want
	for _, filterInfo := range allFilters.WhitelistFilters {
		if filterInfo.Id == id {
			return &filterInfo, true, nil
		}
	}

	// when no matches are found
	return nil, false, nil
}

// GetListFilterByName - Returns a list filter based on its name and whether it's a whitelist filter
func (c *ADG) GetListFilterByName(listName string) (*Filter, bool, error) {
	allFilters, err := c.GetAllFilters()
	if err != nil {
		return nil, false, err
	}

	// go through the blacklist filters in the response until we find the one we want
	for _, filterInfo := range allFilters.Filters {
		if filterInfo.Name == listName {
			return &filterInfo, false, nil
		}
	}
	// go through the whitelist filters in the response until we find the one we want
	for _, filterInfo := range allFilters.WhitelistFilters {
		if filterInfo.Name == listName {
			return &filterInfo, true, nil
		}
	}

	// when no matches are found
	return nil, false, nil
}

// CreateListFilter - Create a list filter, returning the created filter and whether it's a whitelist filter
func (c *ADG) CreateListFilter(filterData AddUrlRequest) (*Filter, bool, error) {
	// convert provided filter to JSON
	rb, err := json.Marshal(filterData)
	if err != nil {
		return nil, false, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/add_url", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, false, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, false, err
	}

	// appease Go as returned body is useless
	_ = body

	// retrieve the filter from the all filters list
	filter, whitelist, err := c.GetListFilterByName(filterData.Name)
	if err != nil {
		return nil, false, err
	}

	// return the filter
	return filter, whitelist, nil
}

// UpdateListFilter - Update a list filter, returning the updated filter and whether it's a whitelist filter
func (c *ADG) UpdateListFilter(filterUpdate FilterSetUrl) (*Filter, bool, error) {
	// convert provided update list info to JSON
	rb, err := json.Marshal(filterUpdate)
	if err != nil {
		return nil, false, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/set_url", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, false, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, false, err
	}

	// appease Go
	_ = body

	// retrieve the filter from the all filters list
	filter, whitelist, err := c.GetListFilterByName(filterUpdate.Data.Name)
	if err != nil {
		return nil, false, err
	}

	// return the filter
	return filter, whitelist, nil
}

// DeleteListFilter - Deletes a list filter
func (c *ADG) DeleteListFilter(filterDelete RemoveUrlRequest) error {
	// convert provided delete filter to JSON
	rb, err := json.Marshal(filterDelete)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/remove_url", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	// appease Go
	_ = body

	// no need to return anything
	return nil
}

// ConfigureFiltering - Configure DNS server filtering parameters
func (c *ADG) ConfigureFiltering(filterConfig FilterConfig) (*FilterConfig, error) {
	// convert provided filtering config to JSON
	rb, err := json.Marshal(filterConfig)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/config", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// appease Go
	_ = body

	// return what was passed
	return &filterConfig, nil
}
