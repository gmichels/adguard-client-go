package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetClients - Get information about configured clients
func (c *ADG) GetClients() (*Clients, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clients", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to a Clients object
	var clients Clients
	err = json.Unmarshal(body, &clients)
	if err != nil {
		return nil, err
	}

	return &clients, nil
}

// SearchClient - Get information about clients by their IP addresses, CIDR, MAC addresses, or ClientIDs
func (c *ADG) SearchClient(identifier string) (*ClientFindResponse, error) {
	// create search request object
	clients := []ClientSearchRequestItem{
		{
			Id: identifier,
		},
	}

	// convert provided clients to JSON
	rb, err := json.Marshal(clients)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/clients/search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to a ClientFindResponse object
	var clientFindResponse ClientFindResponse
	err = json.Unmarshal(body, &clientFindResponse)
	if err != nil {
		return nil, err
	}

	// return the client find response
	return &clientFindResponse, nil
}

// AddClient - Add a new client
func (c *ADG) AddClient(client Client) (*Client, error) {
	// convert provided client to JSON
	rb, err := json.Marshal(client)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/clients/add", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same client that was passed as upstream doesn't return anything
	return &client, nil
}

// UpdateClient - Update client information
func (c *ADG) UpdateClient(clientUpdate ClientUpdate) (*Client, error) {
	// convert provided update client to JSON
	rb, err := json.Marshal(clientUpdate)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/clients/update", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same client that was passed as upstream doesn't return anything
	return &clientUpdate.Data, nil
}

// DeleteClient - Remove a client
func (c *ADG) DeleteClient(clientDelete ClientDelete) error {
	// convert provided delete client to JSON
	rb, err := json.Marshal(clientDelete)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/clients/delete", c.HostURL), strings.NewReader(string(rb)))
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
