package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// Querylog - Get DNS server query log.
func (c *ADG) Querylog(olderThan *string, offset *int, limit *int, search *string, responseStatus *string) (*models.QueryLog, error) {
	// create query parameters dynamically
	queryParams := ""
	if olderThan != nil && *olderThan != "" {
		queryParams += fmt.Sprintf("?older_than=%s", *olderThan)
	}
	if offset != nil && *offset != 0 {
		queryParams += fmt.Sprintf("&offset=%d", *offset)
	}
	if limit != nil && *limit != 0 {
		queryParams += fmt.Sprintf("&limit=%d", *limit)
	}
	if search != nil && *search != "" {
		queryParams += fmt.Sprintf("&search=%s", *search)
	}
	if responseStatus != nil && *responseStatus != "" {
		queryParams += fmt.Sprintf("&response_status=%s", *responseStatus)
	}

	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/querylog", c.HostURL, queryParams), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var querylog models.QueryLog
	err = json.Unmarshal(body, &querylog)
	if err != nil {
		return nil, err
	}

	return &querylog, nil
}

// QuerylogClear - Clear the query log
func (c *ADG) QuerylogClear() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/querylog/clear", c.HostURL), nil)
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

// QuerylogConfig - Get query log parameters
func (c *ADG) QuerylogConfig() (*models.GetQueryLogConfigResponse, error) {
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

	// convert response to object
	var queryLogConfig models.GetQueryLogConfigResponse
	err = json.Unmarshal(body, &queryLogConfig)
	if err != nil {
		return nil, err
	}

	// return the object
	return &queryLogConfig, nil
}

// QuerylogConfigUpdate - Set query log parameters
func (c *ADG) QuerylogConfigUpdate(getQueryLogConfigResponse models.GetQueryLogConfigResponse) error {
	// convert provided object to JSON
	rb, err := json.Marshal(getQueryLogConfigResponse)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/querylog/config/update", c.HostURL), strings.NewReader(string(rb)))
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
