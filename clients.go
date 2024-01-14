package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetAllClients - Returns all clients
func (c *ADG) GetAllClients() (*AllClients, error) {
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

	// convert response to an AllClients object
	var allClients AllClients
	err = json.Unmarshal(body, &allClients)
	if err != nil {
		return nil, err
	}

	return &allClients, nil
}

// GetClient - Returns a client based on an identifier
func (c *ADG) GetClient(identifier string) (*Client, error) {
	// retrieve all clients
	allClients, err := c.GetAllClients()
	if err != nil {
		return nil, err
	}

	// go through the clients in the response until we find the one we want
	for _, clientInfo := range allClients.Clients {
		if clientInfo.Name == identifier {
			return &clientInfo, nil
		}
	}

	// when no matches are found
	return nil, nil
}

// CreateClient - Create a client
func (c *ADG) CreateClient(client Client) (*Client, error) {
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

	// return the same client that was passed
	return &client, nil
}

// UpdateClient - Update a client
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

	// return the client data that was passed
	return &clientUpdate.Data, nil
}

// DeleteClient - Deletes a client
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
