package adguard

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// ADG Client
type ADG struct {
	HostURL    string
	HTTPClient *http.Client
	Auth       AuthStruct
}

// AuthStruct
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewClient
func NewClient(host, username, password, scheme *string, timeout *int) (*ADG, error) {
	// sanity checks
	if *host == "" {
		return nil, fmt.Errorf("required parameter `host`")
	}
	if *username == "" {
		return nil, fmt.Errorf("required parameter `username`")
	}
	if *password == "" {
		return nil, fmt.Errorf("required parameter `password`")
	}
	if *scheme == "" {
		*scheme = "https"
	}
	if *timeout == 0 {
		*timeout = 10
	}

	// instantiate client
	c := ADG{
		HTTPClient: &http.Client{Timeout: time.Duration(*timeout) * time.Second},
		HostURL:    fmt.Sprintf("%s://%s/control", *scheme, *host),
	}
	// instantiate auth
	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	return &c, nil
}

func (c *ADG) doRequest(req *http.Request) ([]byte, error) {
	// add auth info to request
	req.SetBasicAuth(c.Auth.Username, c.Auth.Password)

	// perform request
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// parse body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// sanity check
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
