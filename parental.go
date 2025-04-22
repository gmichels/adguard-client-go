package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gmichels/adguard-client-go/models"
)

// ParentalEnable - Enable parental filtering
func (c *ADG) ParentalEnable() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/parental/enable", c.HostURL), nil)
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

// ParentalDisable - Disable parental filtering
func (c *ADG) ParentalDisable() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/parental/disable", c.HostURL), nil)
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

// ParentalStatus - Get parental filtering status
func (c *ADG) ParentalStatus() (*models.ParentalStatus, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/parental/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var parentalStatus models.ParentalStatus
	err = json.Unmarshal(body, &parentalStatus)
	if err != nil {
		return nil, err
	}

	// return the object
	return &parentalStatus, nil
}
