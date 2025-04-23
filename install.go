package adguard

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gmichels/adguard-client-go/models"
)

// InstallGetAddresses - Gets the network interfaces information
func (c *ADG) InstallGetAddresses() (*models.AddressesInfo, error) {
	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/control/install/get_addresses", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var addressesInfo models.AddressesInfo
	err = json.Unmarshal(body, &addressesInfo)
	if err != nil {
		return nil, err
	}

	return &addressesInfo, nil
}

// InstallCheckConfig - Checks configuration
func (c *ADG) InstallCheckConfig(checkConfigRequest models.CheckConfigRequest) (*models.CheckConfigResponse, error) {
	// convert provided object to JSON
	rb, err := json.Marshal(checkConfigRequest)
	if err != nil {
		return nil, err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/install/check_config", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// convert response to object
	var checkConfigResponse models.CheckConfigResponse
	err = json.Unmarshal(body, &checkConfigResponse)
	if err != nil {
		return nil, err
	}

	// return the object
	return &checkConfigResponse, nil
}

// InstallConfigure - Applies the initial configuration
func (c *ADG) InstallConfigure(initialConfiguration models.InitialConfiguration) error {
	// convert provided object to JSON
	rb, err := json.Marshal(initialConfiguration)
	if err != nil {
		return err
	}

	// initialize request
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/control/install/configure", c.HostURL), strings.NewReader(string(rb)))
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
