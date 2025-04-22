package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// TlsStatus - Returns TLS configuration and its status
func (c *ADG) TlsStatus() (*models.TlsConfig, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/tls/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var tlsConfig models.TlsConfig
	err = json.Unmarshal(body, &tlsConfig)
	if err != nil {
		return nil, err
	}

	// return the object
	return &tlsConfig, nil
}

// TlsConfigure - Updates current TLS configuration
func (c *ADG) TlsConfigure(tlsConfig models.TlsConfig) (*models.TlsConfig, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(tlsConfig)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/tls/configure", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var responseTlsConfig models.TlsConfig
	err = json.Unmarshal(body, &responseTlsConfig)
	if err != nil {
		return nil, err
	}

	// return the object
	return &responseTlsConfig, nil
}

// TlsValidate - Checks if the current TLS configuration is valid
func (c *ADG) TlsValidate(tlsConfig models.TlsConfig) (*models.TlsConfig, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(tlsConfig)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/tls/validate", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var responseTlsConfig models.TlsConfig
	err = json.Unmarshal(body, &responseTlsConfig)
	if err != nil {
		return nil, err
	}

	// return the object
	return &responseTlsConfig, nil
}
