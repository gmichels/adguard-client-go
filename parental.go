package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetParentalStatus - Returns whether parental control is enabled or not
func (c *ADG) GetParentalStatus() (*bool, error) {
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

	// convert response to an Enabled object
	var enabled Enabled
	err = json.Unmarshal(body, &enabled)
	if err != nil {
		return nil, err
	}

	return &enabled.Enabled, nil
}

// SetParentalStatus - Enable or disable parental controls
func (c *ADG) SetParentalStatus(status bool) error {

	// define which endpoint we need to use based on the action
	endpoint := "disable"
	if status {
		endpoint = "enable"
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/parental/%s", c.HostURL, endpoint), strings.NewReader(string([]byte(`{}`))))
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
