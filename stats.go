package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetStatsConfig - Returns statistics configuration parameters
func (c *ADG) GetStatsConfig() (*GetStatsConfigResponse, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/stats/config", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to GetStatsConfigResponse object
	var statsConfig GetStatsConfigResponse
	err = json.Unmarshal(body, &statsConfig)
	if err != nil {
		return nil, err
	}

	return &statsConfig, nil
}

// SetStatsConfig - Sets statistics configuration parameters
func (c *ADG) SetStatsConfig(statsConfig GetStatsConfigResponse) (*GetStatsConfigResponse, error) {
	// convert provided statistics config to JSON
	rb, err := json.Marshal(statsConfig)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/stats/config/update", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same statistics config that was passed
	return &statsConfig, nil
}
