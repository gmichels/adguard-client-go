package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// FilteringStatus - Get filtering parameters
func (c *ADG) FilteringStatus() (*models.FilterStatus, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/filtering/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var allFilters models.FilterStatus
	err = json.Unmarshal(body, &allFilters)
	if err != nil {
		return nil, err
	}

	// retrurn the object
	return &allFilters, nil
}

// FilteringConfig - Set filtering parameters
func (c *ADG) FilteringConfig(filterConfig models.FilterConfig) error {
	// convert provided object to JSON
	rb, err := json.Marshal(filterConfig)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/filtering/config", c.HostURL), strings.NewReader(string(rb)))
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
func (c *ADG) FilteringAddUrl(filterData models.AddUrlRequest) error {
	// convert provided object to JSON
	rb, err := json.Marshal(filterData)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/filtering/add_url", c.HostURL), strings.NewReader(string(rb)))
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
func (c *ADG) FilteringRemoveUrl(filterDelete models.RemoveUrlRequest) error {
	// convert provided object to JSON
	rb, err := json.Marshal(filterDelete)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/filtering/remove_url", c.HostURL), strings.NewReader(string(rb)))
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
func (c *ADG) FilteringSetUrl(filterUpdate models.FilterSetUrl) error {
	// convert provided object to JSON
	rb, err := json.Marshal(filterUpdate)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/filtering/set_url", c.HostURL), strings.NewReader(string(rb)))
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
func (c *ADG) FilteringRefresh(filterRefreshRequest models.FilterRefreshRequest) (*models.FilterRefreshResponse, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(filterRefreshRequest)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/filtering/refresh", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var filterRefreshResponse models.FilterRefreshResponse
	err = json.Unmarshal(body, &filterRefreshResponse)
	if err != nil {
		return nil, err
	}

	// return the object
	return &filterRefreshResponse, nil
}

// FilteringSetRules - Set user-defined filter rules
func (c *ADG) FilteringSetRules(rules models.SetRulesRequest) error {
	// convert provided object to JSON
	rb, err := json.Marshal(rules)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/filtering/set_rules", c.HostURL), strings.NewReader(string(rb)))
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
func (c *ADG) FilteringCheckHost(name *string, client *string, qtype *string) (*models.FilterCheckHostResponse, error) {
	// create query parameters dynamically
	queryParams := fmt.Sprintf("?name=%s", *name)
	if client != nil && *client != "" {
		queryParams += fmt.Sprintf("&client=%s", *client)
	}
	if qtype != nil && *qtype != "" {
		queryParams += fmt.Sprintf("&qtype=%s", *qtype)
	}

	// initialize request with query parameters
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/filtering/check_host%s", c.HostURL, queryParams), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var filterCheckHostResponse models.FilterCheckHostResponse
	err = json.Unmarshal(body, &filterCheckHostResponse)
	if err != nil {
		return nil, err
	}

	// return the object
	return &filterCheckHostResponse, nil
}
