package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// FilteringStatus - Get filtering parameters
func (c *ADG) FilteringStatus() (*FilterStatus, error) {
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

	// convert response to an object
	var allFilters FilterStatus
	err = json.Unmarshal(body, &allFilters)
	if err != nil {
		return nil, err
	}

	// retrurn the object
	return &allFilters, nil
}

// FilteringConfig - Set filtering parameters
func (c *ADG) FilteringConfig(filterConfig FilterConfig) error {
	// convert provided object to JSON
	rb, err := json.Marshal(filterConfig)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/config", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	// perform request
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	// return nothing
	return nil
}

// FilteringAddUrl - Add filter URL or an absolute file path
func (c *ADG) FilteringAddUrl(filterData AddUrlRequest) error {
	// convert provided object to JSON
	rb, err := json.Marshal(filterData)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/add_url", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	// perform request
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	// return nothing
	return nil
}

// FilteringRemoveUrl - Remove filter URL
func (c *ADG) FilteringRemoveUrl(filterDelete RemoveUrlRequest) error {
	// convert provided object to JSON
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
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	// return nothing
	return nil
}

// FilteringSetUrl - Set URL parameters
func (c *ADG) FilteringSetUrl(filterUpdate FilterSetUrl) error {
	// convert provided object to JSON
	rb, err := json.Marshal(filterUpdate)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/set_url", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	// perform request
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	// return nothing
	return nil
}

// FilteringRefresh - Set URL parameters
func (c *ADG) FilteringRefresh(filterRefreshRequest FilterRefreshRequest) (*FilterRefreshResponse, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(filterRefreshRequest)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/refresh", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to an object
	var filterRefreshResponse FilterRefreshResponse
	err = json.Unmarshal(body, &filterRefreshResponse)
	if err != nil {
		return nil, err
	}

	// return the object
	return &filterRefreshResponse, nil
}

// FilteringSetRules - Set user-defined filter rules
func (c *ADG) FilteringSetRules(rules SetRulesRequest) error {
	// convert provided object to JSON
	rb, err := json.Marshal(rules)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/filtering/set_rules", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	// perform request
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	// return nothing
	return nil
}

// FilteringCheckHost - Check if host name is filtered
func (c *ADG) FilteringCheckHost(name string, client *string, qtype *string) (*FilterCheckHostResponse, error) {
	// create query parameters dynamically
	queryParams := fmt.Sprintf("?name=%s", name)
	if client != nil && *client != "" {
		queryParams += fmt.Sprintf("&client=%s", *client)
	}
	if qtype != nil && *qtype != "" {
		queryParams += fmt.Sprintf("&qtype=%s", *qtype)
	}

	// initialize request with query parameters
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/filtering/check_host%s", c.HostURL, queryParams), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to an object
	var filterCheckHostResponse FilterCheckHostResponse
	err = json.Unmarshal(body, &filterCheckHostResponse)
	if err != nil {
		return nil, err
	}

	// Return the object
	return &filterCheckHostResponse, nil
}
