package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// DhcpStatus - Gets the current DHCP settings and status
func (c *ADG) DhcpStatus() (*models.DhcpStatus, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/dhcp/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var dhcpStatus models.DhcpStatus
	err = json.Unmarshal(body, &dhcpStatus)
	if err != nil {
		return nil, err
	}

	// return the object
	return &dhcpStatus, nil
}

// DhcpInterfaces - Gets the available DHCP interfaces
func (c *ADG) DhcpInterfaces() (*models.NetInterfaces, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/dhcp/interfaces", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var netInterfaces models.NetInterfaces
	err = json.Unmarshal(body, &netInterfaces)
	if err != nil {
		return nil, err
	}

	// return the object
	return &netInterfaces, nil
}

// DhcpSetConfig - Updates the current DHCP server configuration
func (c *ADG) DhcpSetConfig(dhcpConfig models.DhcpConfig) error {
	// convert provided object to JSON
	rb, err := json.Marshal(dhcpConfig)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/dhcp/set_config", c.HostURL), strings.NewReader(string(rb)))
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

// DhcpFindActiveDhcp - Searches for an active DHCP server on the network
func (c *ADG) DhcpFindActiveDhcp(dhcpFindActiveReq models.DhcpFindActiveReq) (*models.DhcpSearchResult, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(dhcpFindActiveReq)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/dhcp/find_active_dhcp", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var dhcpSearchResult models.DhcpSearchResult
	err = json.Unmarshal(body, &dhcpSearchResult)
	if err != nil {
		return nil, err
	}

	// return the object
	return &dhcpSearchResult, nil
}

// DhcpAddStaticLease - Adds a static lease
func (c *ADG) DhcpAddStaticLease(dhcpStaticLease models.DhcpStaticLease) error {
	// convert provided object to JSON
	rb, err := json.Marshal(dhcpStaticLease)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/dhcp/add_static_lease", c.HostURL), strings.NewReader(string(rb)))
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

// DhcpRemoveStaticLease - Removes a static lease
func (c *ADG) DhcpRemoveStaticLease(dhcpStaticLease models.DhcpStaticLease) error {
	// convert provided object to JSON
	rb, err := json.Marshal(dhcpStaticLease)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/dhcp/remove_static_lease", c.HostURL), strings.NewReader(string(rb)))
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

// DhcpUpdateStaticLease - Updates a static lease
func (c *ADG) DhcpUpdateStaticLease(dhcpStaticLease models.DhcpStaticLease) error {
	// convert provided object to JSON
	rb, err := json.Marshal(dhcpStaticLease)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/dhcp/update_static_lease", c.HostURL), strings.NewReader(string(rb)))
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

// DhcpReset - Reset DHCP configuration
func (c *ADG) DhcpReset() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/dhcp/reset", c.HostURL), nil)
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

// DhcpResetLeases - Reset DHCP leases
func (c *ADG) DhcpResetLeases() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/dhcp/reset_leases", c.HostURL), nil)
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
