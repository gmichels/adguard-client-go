package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gmichels/adguard-client-go/models"
)

// SafeBrowsingEnable - Enable safebrowsing
func (c *ADG) SafeBrowsingEnable() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/safebrowsing/enable", c.HostURL), nil)
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

// SafeBrowsingDisable - Disable safebrowsing
func (c *ADG) SafeBrowsingDisable() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/safebrowsing/disable", c.HostURL), nil)
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

// SafeBrowsingStatus - Get safebrowsing status
func (c *ADG) SafeBrowsingStatus() (*models.Enabled, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/safebrowsing/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var enabled models.Enabled
	err = json.Unmarshal(body, &enabled)
	if err != nil {
		return nil, err
	}

	// return the object
	return &enabled, nil
}
