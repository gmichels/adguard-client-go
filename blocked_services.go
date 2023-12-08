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
func (c *ADG) GetBlockedServices() (*BlockedServicesSchedule, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/blocked_services/get", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to BlockedServicesSchedule object
	var blockedServicesSchedule BlockedServicesSchedule
	err = json.Unmarshal(body, &blockedServicesSchedule)
	if err != nil {
		return nil, err
	}

	return &blockedServicesSchedule, nil
}

// UpdateBlockedServices - Update blocked services
func (c *ADG) SetBlockedServices(blockedServicesSchedule BlockedServicesSchedule) (*BlockedServicesSchedule, error) {
	// convert provided blocked services to JSON
	rb, err := json.Marshal(blockedServicesSchedule)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/blocked_services/update", c.HostURL), strings.NewReader(string(rb)))
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
	return &blockedServicesSchedule, nil
}
