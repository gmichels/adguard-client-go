package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetTlsConfig - Returns TLS configuration and status
func (c *ADG) GetTlsConfig() (*TlsConfig, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/tls/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to AccessList object
	var tlsConfig TlsConfig
	err = json.Unmarshal(body, &tlsConfig)
	if err != nil {
		return nil, err
	}

	return &tlsConfig, nil
}

// SetAccess - Sets the access list
func (c *ADG) SetTlsConfig(tlsConfig TlsConfig) (*TlsConfig, error) {
	// convert provided TLS config to JSON
	rb, err := json.Marshal(tlsConfig)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tls/configure", c.HostURL), strings.NewReader(string(rb)))
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
	return &tlsConfig, nil
}
