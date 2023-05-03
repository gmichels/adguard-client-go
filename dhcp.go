package adguard

import (
	"fmt"
	"net/http"
	"strings"

	// custom json module to allow for omitting zero value structs
	"github.com/clarketm/json"
)

// GetDhcpStatus - Returns the current DHCP status
func (c *ADG) GetDhcpStatus() (*DhcpStatus, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dhcp/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to DhcpStatus object
	var dhcpStatus DhcpStatus
	err = json.Unmarshal(body, &dhcpStatus)
	if err != nil {
		return nil, err
	}

	return &dhcpStatus, nil
}

// GetDhcpStatus - Returns the available DHCP interfaces
func (c *ADG) GetDhcpInterfaces() (*NetInterfaces, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/dhcp/interfaces", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to NetInterfaces object
	var netInterfaces NetInterfaces
	err = json.Unmarshal(body, &netInterfaces)
	if err != nil {
		return nil, err
	}

	return &netInterfaces, nil
}

// SetDhcpConfig - Sets DHCP server configuration
func (c *ADG) SetDhcpConfig(dhcpConfig DhcpConfig) (*DhcpConfig, error) {
	// convert provided DHCP config to JSON
	rb, err := json.Marshal(dhcpConfig)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dhcp/set_config", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same DHCP config that was passed
	return &dhcpConfig, nil
}

// ManageDhcpStaticLease - Add ot remove a DHCP static lease
func (c *ADG) ManageDhcpStaticLease(add bool, dhcpStaticLease DhcpStaticLease) (*DhcpStaticLease, error) {
	// convert provided DHCP static lease to JSON
	rb, err := json.Marshal(dhcpStaticLease)
	if err != nil {
		return nil, err
	}

	// default to remove operation
	operation := "remove"
	if add {
		// flip to add operation
		operation = "add"
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dhcp/%s_static_lease", c.HostURL, operation), strings.NewReader(string(rb)))
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

	// return the same DHCP static lease that was passed
	return &dhcpStaticLease, nil
}

// ResetDhcpConfig - reset all DHCP configuration to defaults
func (c *ADG) ResetDhcpConfig() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dhcp/reset", c.HostURL), strings.NewReader(string([]byte(`{}`))))
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

	// return nothing
	return nil
}
