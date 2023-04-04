package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// GetDnsInfo - Returns the current DNS parameters
func (c *ADG) GetDnsInfo() (*DNSInfo, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dns_info", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to DNSInfo object
	var dnsInfo DNSInfo
	err = json.Unmarshal(body, &dnsInfo)
	if err != nil {
		return nil, err
	}

	return &dnsInfo, nil
}

// SetDnsConfig - Sets DNS parameters
func (c *ADG) SetDnsConfig(dnsConfig DNSConfig) (*DNSConfig, error) {
	// convert provided DNS Config to JSON
	rb, err := json.Marshal(dnsConfig)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dns_config", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same DNS config that was passed
	return &dnsConfig, nil
}
