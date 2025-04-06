package adguard

import (
	"fmt"
	"net/http"
	"strings"

	// custom json module to allow for omitting zero value structs
	"github.com/clarketm/json"
)

// DhcpStatus - Gets the current DHCP settings and status
func (c *ADG) DhcpStatus() (*DhcpStatus, error) {
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

// DhcpStatus - Gets the available DHCP interfaces
func (c *ADG) DhcpInterfaces() (*NetInterfaces, error) {
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

// DhcpSetConfig - Updates the current DHCP server configuration
func (c *ADG) DhcpSetConfig(dhcpConfig DhcpConfig) (*DhcpConfig, error) {
	// convert provided object to JSON
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

// DhcpFindActiveDhcp - Searches for an active DHCP server on the network
func (c *ADG) DhcpFindActiveDhcp(dhcpFindActiveReq DhcpFindActiveReq) (*DhcpSearchResult, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(dhcpFindActiveReq)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dhcp/find_active_dhcp", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to a DhcpSearchResult object
	var dhcpSearchResult DhcpSearchResult
	err = json.Unmarshal(body, &dhcpSearchResult)
	if err != nil {
		return nil, err
	}

	// return the DHCP search result
	return &dhcpSearchResult, nil
}

// DhcpAddStaticLease - Adds a static lease
func (c *ADG) DhcpAddStaticLease(dhcpStaticLease DhcpStaticLease) (*DhcpStaticLease, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(dhcpStaticLease)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dhcp/add_static_lease", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same object that was passed
	return &dhcpStaticLease, nil
}

// DhcpRemoveStaticLease - Removes a static lease
func (c *ADG) DhcpRemoveStaticLease(dhcpStaticLease DhcpStaticLease) (*DhcpStaticLease, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(dhcpStaticLease)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dhcp/remove_static_lease", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same object that was passed
	return &dhcpStaticLease, nil
}

// DhcpUpdateStaticLease - Updates a static lease
func (c *ADG) DhcpUpdateStaticLease(dhcpStaticLease DhcpStaticLease) (*DhcpStaticLease, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(dhcpStaticLease)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dhcp/update_static_lease", c.HostURL), strings.NewReader(string(rb)))
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

	// return the same object that was passed
	return &dhcpStaticLease, nil
}

// DhcpReset - Reset DHCP configuration
func (c *ADG) DhcpReset() error {
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

// DhcpResetLeases - Reset DHCP leases
func (c *ADG) DhcpResetLeases() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dhcp/reset_leases", c.HostURL), strings.NewReader(string([]byte(`{}`))))
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
