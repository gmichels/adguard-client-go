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

// UpdateUserRules - Returns a list of all user rules
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

// GetListFilterById - Returns a list filter based on its id and whether it's a blacklist/whitelist filter
func (c *ADG) GetListFilterById(id int64, filterType string) (*Filter, error) {
	allFilters, err := c.GetAllFilters()
	if err != nil {
		return nil, err
	}

	// go through the filters in the response until we find the one we want, based on the filter type
	if filterType == "blacklist" {
		for _, filterInfo := range allFilters.Filters {
			if filterInfo.Id == id {
				return &filterInfo, nil
			}
		}
	} else if filterType == "whitelist" {
		for _, filterInfo := range allFilters.WhitelistFilters {
			if filterInfo.Id == id {
				return &filterInfo, nil
			}
		}
	} else {
		// if we got here, someone passed gibberish for filterType
		return nil, fmt.Errorf("unknown value `%s` for `filterType` (allowed values are `blacklist` or `whitelist`)", filterType)
	}

	// when no matches are found
	return nil, nil
}

// GetListFilterByName - Returns a list filter based on its name and whether it's a blacklist/whitelist filter
func (c *ADG) GetListFilterByName(listName string, filterType string) (*Filter, error) {
	allFilters, err := c.GetAllFilters()
	if err != nil {
		return nil, err
	}

	// go through the filters in the response until we find the one we want, based on the filter type
	if filterType == "blacklist" {
		for _, filterInfo := range allFilters.Filters {
			if filterInfo.Name == listName {
				return &filterInfo, nil
			}
		}
	} else if filterType == "whitelist" {
		for _, filterInfo := range allFilters.WhitelistFilters {
			if filterInfo.Name == listName {
				return &filterInfo, nil
			}
		}
	} else {
		// if we got here, someone passed gibberish for filterType
		return nil, fmt.Errorf("unknown value `%s` for `filterType` (allowed values are `blacklist` or `whitelist`)", filterType)
	}

	// when no matches are found
	return nil, nil
}

// CreateListFilter - Create a list filter
func (c *ADG) CreateListFilter(filterData AddUrlRequest) (*Filter, error) {
	// convert provided filter to JSON
	rb, err := json.Marshal(filterData)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/add_url", c.HostURL), strings.NewReader(string(rb)))
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

	// default filter type to blacklist
	filterType := "blacklist"
	if filterData.Whitelist {
		// switch to whitelist
		filterType = "whitelist"
	}

	// retrieve the filter from the all filters list
	filter, err := c.GetListFilterByName(filterData.Name, filterType)
	if err != nil {
		return nil, err
	}

	// return the filter
	return filter, nil
}

// UpdateListFilter - Update a list filter
func (c *ADG) UpdateListFilter(filterUpdate FilterSetUrl) (*FilterSetUrl, error) {
	// convert provided update list info to JSON
	rb, err := json.Marshal(filterUpdate)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/set_url", c.HostURL), strings.NewReader(string(rb)))
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

	// return the data that was passed
	return &filterUpdate, nil
}

// DeleteListFilter - Deletes a client
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
