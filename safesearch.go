package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// SafeSearchSettings - Update safesearch settings
func (c *ADG) SafeSearchSettings(safeSearchConfig models.SafeSearchConfig) error {
	// convert provided object to JSON
	rb, err := json.Marshal(safeSearchConfig)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/safesearch/settings", c.HostURL), strings.NewReader(string(rb)))
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

// SafeSearchStatus - Get safesearch status
func (c *ADG) SafeSearchStatus() (*models.SafeSearchConfig, error) {
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

	// convert response to object
	var safeSearchConfig models.SafeSearchConfig
	err = json.Unmarshal(body, &safeSearchConfig)
	if err != nil {
		return nil, err
	}

	// return the object
	return &safeSearchConfig, nil
}
