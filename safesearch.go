package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetSafeSearchStatus - Retrieve safe search configuration
func (c *ADG) GetSafeSearchConfig() (*SafeSearchConfig, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/safesearch/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to a SafeSerchConfig object
	var safeSearchConfig SafeSearchConfig
	err = json.Unmarshal(body, &safeSearchConfig)
	if err != nil {
		return nil, err
	}

	return &safeSearchConfig, nil
}

// SetSafeSearchConfig - Set safe search configuration
func (c *ADG) SetSafeSearchConfig(safeSearchConfig SafeSearchConfig) (*SafeSearchConfig, error) {
	// convert provided safe search config to JSON
	rb, err := json.Marshal(safeSearchConfig)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/safesearch/settings", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same safe search configuration that was passed
	return &safeSearchConfig, nil
}
