package adguard

import (
	"encoding/json"
	// "errors"
	"fmt"
	"net/http"
	// "strings"
)

// GetAllClients - Returns all clients
// func (c *ADG) GetAllClients() (*[]Client, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clients", c.HostURL), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var clients Clients
// 	// clients := {}Clients
// 	err = json.Unmarshal(body, &clients)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &clients.Clients, nil
// }

// GetAllClients - Returns a client based on an identifier
func (c *ADG) GetClient(identifier string) (*Client, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/clients", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var allClients AllClients
	// rest := {}Clients
	err = json.Unmarshal(body, &allClients)
	if err != nil {
		return nil, err
	}

	for _, clientInfo := range allClients.Clients {
		if clientInfo.Name == identifier {
			return &clientInfo, nil
		}
	}

	return nil, nil
}
