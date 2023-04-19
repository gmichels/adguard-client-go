package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetQueryLogConfig - Returns query log configuration parameters
func (c *ADG) GetQueryLogConfig() (*GetQueryLogConfigResponse, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/querylog/config", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to GetQueryLogConfigResponse object
	var queryLogConfig GetQueryLogConfigResponse
	err = json.Unmarshal(body, &queryLogConfig)
	if err != nil {
		return nil, err
	}

	return &queryLogConfig, nil
}

// SetQueryLogConfig - Sets query log configuration parameters
func (c *ADG) SetQueryLogConfig(queryLogConfig GetQueryLogConfigResponse) (*GetQueryLogConfigResponse, error) {
	// convert provided query log config to JSON
	rb, err := json.Marshal(queryLogConfig)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/querylog/config/update", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same query log config that was passed
	return &queryLogConfig, nil
}
