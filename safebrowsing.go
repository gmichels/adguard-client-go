package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetSafeBrowsingStatus - Returns whether safe-browsing is enabled or not
func (c *ADG) GetSafeBrowsingStatus() (*bool, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/safebrowsing/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to SafeBrowsingStatus object
	var safeBrowsingStatus SafeBrowsingStatus
	err = json.Unmarshal(body, &safeBrowsingStatus)
	if err != nil {
		return nil, err
	}

	return &safeBrowsingStatus.Enabled, nil
}

// SetSafeBrowsingStatus - Enable or disable safe-browsing
func (c *ADG) SetSafeBrowsingStatus(status bool) error {

	// define which endpoint we need to use based on the action
	endpoint := "disable"
	if status {
		endpoint = "enable"
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/safebrowsing/%s", c.HostURL, endpoint), nil)
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

	// return nothing
	return nil
}
