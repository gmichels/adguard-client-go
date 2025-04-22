package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
	// custom json module to allow for omitting zero value structs
	// "github.com/clarketm/json"
)

// BlockedServicesAll - Get available services to use for blocking
func (c *ADG) BlockedServicesAll() (*models.BlockedServicesAll, error) {
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

	// convert response to object
	var blockedServicesAll models.BlockedServicesAll
	err = json.Unmarshal(body, &blockedServicesAll)
	if err != nil {
		return nil, err
	}

	// return the object
	return &blockedServicesAll, nil
}

// BlockedServicesGet - Get blocked services
func (c *ADG) BlockedServicesGet() (*models.BlockedServicesSchedule, error) {
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

	// convert response to object
	var blockedServicesSchedule models.BlockedServicesSchedule
	err = json.Unmarshal(body, &blockedServicesSchedule)
	if err != nil {
		return nil, err
	}

	// return the object
	return &blockedServicesSchedule, nil
}

// BlockedServicesUpdate - Update blocked services
func (c *ADG) BlockedServicesUpdate(blockedServicesSchedule models.BlockedServicesSchedule) error {
	// convert provided object to JSON
	rb, err := json.Marshal(blockedServicesSchedule)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/blocked_services/update", c.HostURL), strings.NewReader(string(rb)))
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
