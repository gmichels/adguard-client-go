package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// RewriteList - Get list of Rewrite rules
func (c *ADG) RewriteList() (*models.RewriteList, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/rewrite/list", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var rewriteList models.RewriteList
	err = json.Unmarshal(body, &rewriteList)
	if err != nil {
		return nil, err
	}

	// return the object
	return &rewriteList, nil
}

// RewriteAdd - Add a new Rewrite rule
func (c *ADG) RewriteAdd(rewriteEntry models.RewriteEntry) error {
	// convert provided object to JSON
	rb, err := JSONMarshal(rewriteEntry)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/rewrite/add", c.HostURL), strings.NewReader(string(rb)))
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

// RewriteDelete - Remove a Rewrite rule
func (c *ADG) RewriteDelete(rewriteEntry models.RewriteEntry) error {
	// convert provided object to JSON
	rb, err := JSONMarshal(rewriteEntry)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/rewrite/delete", c.HostURL), strings.NewReader(string(rb)))
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

// RewriteUpdate - Update a Rewrite rule
func (c *ADG) RewriteUpdate(rewriteUpdate models.RewriteUpdate) error {
	// convert provided object to JSON
	rb, err := JSONMarshal(rewriteUpdate)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/control/rewrite/update", c.HostURL), strings.NewReader(string(rb)))
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
