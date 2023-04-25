package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetBlockedServicesList - Returns the list of all available services to be blocked
func (c *ADG) GetBlockedServicesList() (*BlockedServicesAll, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blocked_services/all", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to BlockedServicesAll object
	var blockedServicesAll BlockedServicesAll
	err = json.Unmarshal(body, &blockedServicesAll)
	if err != nil {
		return nil, err
	}

	return &blockedServicesAll, nil
}

// GetBlockedServices - Returns the services that are blocked globally
func (c *ADG) GetBlockedServices() (*[]string, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blocked_services/list", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to a list of strings
	var blockedServices []string
	err = json.Unmarshal(body, &blockedServices)
	if err != nil {
		return nil, err
	}

	return &blockedServices, nil
}

// SetBlockedServices - Sets the services to be blocked globally
func (c *ADG) SetBlockedServices(blockedServices []string) (*[]string, error) {
	// convert provided blocked services to JSON
	rb, err := json.Marshal(blockedServices)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/blocked_services/set", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same blocked services that were passed
	return &blockedServices, nil
}
