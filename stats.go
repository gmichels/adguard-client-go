package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// Stats - Get DNS server statistics
func (c *ADG) Stats() (*models.Stats, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/stats", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var stats models.Stats
	err = json.Unmarshal(body, &stats)
	if err != nil {
		return nil, err
	}

	// return the object
	return &stats, nil
}

// StatsReset - Reset all statistics to zeroes
func (c *ADG) StatsReset() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/stats_reset", c.HostURL), nil)
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

// StatsConfig - Get statistics parameters
func (c *ADG) StatsConfig() (*models.GetStatsConfigResponse, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/stats/config", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var statsConfig models.GetStatsConfigResponse
	err = json.Unmarshal(body, &statsConfig)
	if err != nil {
		return nil, err
	}

	// return the object
	return &statsConfig, nil
}

// StatsConfigUpdate - Sets statistics parameters
func (c *ADG) StatsConfigUpdate(statsConfig models.GetStatsConfigResponse) error {
	// convert provided object to JSON
	rb, err := JSONMarshal(statsConfig)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/control/stats/config/update", c.HostURL), strings.NewReader(string(rb)))
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
