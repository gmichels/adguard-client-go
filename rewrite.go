package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetAllRewrites - Returns all DNS rewrite rules
func (c *ADG) GetAllRewrites() (*[]RewriteEntry, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/rewrite/list", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to a list of RewriteEntry objects
	var allRewrites []RewriteEntry
	err = json.Unmarshal(body, &allRewrites)
	if err != nil {
		return nil, err
	}

	return &allRewrites, nil
}

// GetRewrite - Return a DNS rewrite rule based on the domain
func (c *ADG) GetRewrite(domain string) (*RewriteEntry, error) {
	// retrieve all DNS rewrite rules
	allRewrites, err := c.GetAllRewrites()
	if err != nil {
		return nil, err
	}

	// loop over the results until we find the one we want
	for _, rewrite := range *allRewrites {
		if rewrite.Domain == domain {
			return &rewrite, nil
		}
	}

	// when no matches are found
	return nil, nil
}

// CreateRewrite - Create a DNS rewrite rule
func (c *ADG) CreateRewrite(rewrite RewriteEntry) (*RewriteEntry, error) {
	// convert provided rewrite rule to JSON
	rb, err := json.Marshal(rewrite)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/rewrite/add", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same rewrite rule that was passed
	return &rewrite, nil
}

// UpdateRewrite - Update a DNS rewrite rule
func (c *ADG) UpdateRewrite(rewrite RewriteEntry) (*RewriteEntry, error) {
	// there is no real update endpoint in ADG, need to delete and re-create
	err := c.DeleteRewrite(rewrite.Domain)
	if err != nil {
		return nil, err
	}

	_, err = c.CreateRewrite(rewrite)
	if err != nil {
		return nil, err
	}

	// return the rewrite rule that was passed
	return &rewrite, nil
}

// DeleteRewrite - Delete a DNS rewrite rule based on the domain
func (c *ADG) DeleteRewrite(domain string) error {
	// confirm the DNS rewrite rule entry exists
	rewrite, err := c.GetRewrite(domain)
	if err != nil {
		return err
	} else if rewrite == nil {
		return fmt.Errorf("unable to find a DNS rewrite rule for `%s`", domain)
	}

	// convert DNS rewrite to JSON
	rb, err := json.Marshal(rewrite)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/rewrite/delete", c.HostURL), strings.NewReader(string(rb)))
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

	// no need to return anything
	return nil
}
