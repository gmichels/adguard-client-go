package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// Clients - Get information about configured clients
func (c *ADG) Clients() (*models.Clients, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/clients", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var clients models.Clients
	err = json.Unmarshal(body, &clients)
	if err != nil {
		return nil, err
	}

	// return the object
	return &clients, nil
}

// ClientsAdd - Add a new client
func (c *ADG) ClientsAdd(client models.Client) error {
	// convert provided object to JSON
	rb, err := json.Marshal(client)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/clients/add", c.HostURL), strings.NewReader(string(rb)))
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

// ClientsDelete - Remove a client
func (c *ADG) ClientsDelete(clientDelete models.ClientDelete) error {
	// convert provided object to JSON
	rb, err := json.Marshal(clientDelete)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/clients/delete", c.HostURL), strings.NewReader(string(rb)))
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

// ClientsUpdate - Update client information
func (c *ADG) ClientsUpdate(clientUpdate models.ClientUpdate) error {
	// convert provided object to JSON
	rb, err := json.Marshal(clientUpdate)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/clients/update", c.HostURL), strings.NewReader(string(rb)))
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

// ClientsSearch - Get information about clients by their IP addresses, CIDR, MAC addresses, or ClientIDs
func (c *ADG) ClientsSearch(identifiers []string) (*models.ClientsFindResponse, error) {
	// initialize object
	clients := []models.ClientSearchRequestItem{}

	// go through identifiers and append individual objects
	for _, identifier := range identifiers {
		clients = append(clients, models.ClientSearchRequestItem{
			Id: identifier,
		})
	}

	// initialize request object
	clientsSearchRequest := models.ClientsSearchRequest{
		Clients: clients,
	}

	// convert object to JSON
	rb, err := json.Marshal(clientsSearchRequest)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/clients/search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var clientsFindResponse models.ClientsFindResponse
	err = json.Unmarshal(body, &clientsFindResponse)
	if err != nil {
		return nil, err
	}

	// return the object
	return &clientsFindResponse, nil
}
