package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// AccessList - List (dis)allowed clients, blocked hosts, etc
func (c *ADG) AccessList() (*models.AccessList, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/access/list", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var accessList models.AccessList
	err = json.Unmarshal(body, &accessList)
	if err != nil {
		return nil, err
	}

	return &accessList, nil
}

// AccessSet - Set (dis)allowed clients, blocked hosts, etc
func (c *ADG) AccessSet(accessList models.AccessList) error {
	// convert provided object to JSON
	rb, err := JSONMarshal(accessList)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/access/set", c.HostURL), strings.NewReader(string(rb)))
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
