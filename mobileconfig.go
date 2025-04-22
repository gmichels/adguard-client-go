package adguard

import (
	"fmt"
	"net/http"
	"strings"
)

// AppleDohMobileconfig - Get DNS over HTTPS .mobileconfig
func (c *ADG) AppleDohMobileconfig(host, clientId *string) (string, error) {
	// create query parameters dynamically
	queryParams := ""
	queryParamsList := []string{}

	if host != nil && *host != "" {
		queryParamsList = append(queryParamsList, fmt.Sprintf("host=%s", *host))
	}
	if clientId != nil && *clientId != "" {
		queryParamsList = append(queryParamsList, fmt.Sprintf("client_id=%s", *clientId))
	}

	if len(queryParamsList) > 0 {
		queryParams = "?" + strings.Join(queryParamsList, "&")
	}

	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apple/doh.mobileconfig%s", c.HostURL, queryParams), nil)
	if err != nil {
		return "", err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	// return the string
	return string(body), nil
}

// AppleDotMobileconfig - Get DNS over TLS .mobileconfig
func (c *ADG) AppleDotMobileconfig(host, clientId *string) (string, error) {
	// create query parameters dynamically
	queryParams := ""
	queryParamsList := []string{}

	if host != nil && *host != "" {
		queryParamsList = append(queryParamsList, fmt.Sprintf("host=%s", *host))
	}
	if clientId != nil && *clientId != "" {
		queryParamsList = append(queryParamsList, fmt.Sprintf("client_id=%s", *clientId))
	}

	if len(queryParamsList) > 0 {
		queryParams = "?" + strings.Join(queryParamsList, "&")
	}

	// initialize request
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/apple/dot.mobileconfig%s", c.HostURL, queryParams), nil)
	if err != nil {
		return "", err
	}

	// perform request
	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	// return the string
	return string(body), nil
}
