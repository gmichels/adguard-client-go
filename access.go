package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// ListAccess - List (dis)allowed clients, blocked hosts, etc
func (c *ADG) ListAccess() (*AccessList, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/access/list", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to AccessList object
	var accessList AccessList
	err = json.Unmarshal(body, &accessList)
	if err != nil {
		return nil, err
	}

	return &accessList, nil
}

// SetAccess - Set (dis)allowed clients, blocked hosts, etc
func (c *ADG) SetAccess(accessList AccessList) (*AccessList, error) {
	// convert provided access list to JSON
	rb, err := json.Marshal(accessList)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/access/set", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same access list that was passed
	return &accessList, nil
}
