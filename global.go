package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// Status - Get DNS server current status and general settings
func (c *ADG) Status() (*models.ServerStatus, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/status", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var serverStatus models.ServerStatus
	err = json.Unmarshal(body, &serverStatus)
	if err != nil {
		return nil, err
	}

	// return the object
	return &serverStatus, nil
}

// DnsInfo - Get genedal DNS parameters
func (c *ADG) DnsInfo() (*models.DNSInfo, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/dns_info", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var dnsInfo models.DNSInfo
	err = json.Unmarshal(body, &dnsInfo)
	if err != nil {
		return nil, err
	}

	// return the object
	return &dnsInfo, nil
}

// DnsConfig - Set general DNS parameters
func (c *ADG) DnsConfig(dnsConfig models.DNSConfig) error {
	// convert provided object to JSON
	rb, err := json.Marshal(dnsConfig)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/dns_config", c.HostURL), strings.NewReader(string(rb)))
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

// Protection - Set protection state and duration
func (c *ADG) Protection(setProtectionRequest models.SetProtectionRequest) error {
	// convert provided object to JSON
	rb, err := json.Marshal(setProtectionRequest)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/protection", c.HostURL), strings.NewReader(string(rb)))
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

// CacheClear - Clear DNS cache
func (c *ADG) CacheClear() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/cache_clear", c.HostURL), nil)
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

// TestUpstreamDns - Test upstream configuration
func (c *ADG) TestUpstreamDns(upstreamsConfig models.UpstreamsConfig) (*models.UpstreamsConfigResponse, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(upstreamsConfig)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/test_upstream_dns", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var upstreamsConfigResponse models.UpstreamsConfigResponse
	err = json.Unmarshal(body, &upstreamsConfigResponse)
	if err != nil {
		return nil, err
	}

	// return the object
	return &upstreamsConfigResponse, nil
}

// VersionJson - Gets information about the latest available version of AdGuard
func (c *ADG) VersionJson(getVersionRequest models.GetVersionRequest) (*models.VersionInfo, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(getVersionRequest)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/version.json", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var versionInfo models.VersionInfo
	err = json.Unmarshal(body, &versionInfo)
	if err != nil {
		return nil, err
	}

	// return the object
	return &versionInfo, nil
}

// Update - Begin auto-upgrade procedure
func (c *ADG) Update() error {
	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/update", c.HostURL), nil)
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

// Login - Perform administrator login
func (c *ADG) Login(login models.Login) error {
	// convert provided object to JSON
	rb, err := json.Marshal(login)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/login", c.HostURL), strings.NewReader(string(rb)))
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

// Logout - Perform administrator logout
func (c *ADG) Logout() error {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/logout", c.HostURL), nil)
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

// ProfileUpdate - Updates current user info
func (c *ADG) ProfileUpdate(profileInfo models.ProfileInfo) error {
	// convert provided object to JSON
	rb, err := json.Marshal(profileInfo)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/control/profile/update", c.HostURL), strings.NewReader(string(rb)))
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

// Profile
func (c *ADG) Profile() (*models.ProfileInfo, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/profile", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var profileInfo models.ProfileInfo
	err = json.Unmarshal(body, &profileInfo)
	if err != nil {
		return nil, err
	}

	// return the object
	return &profileInfo, nil
}
