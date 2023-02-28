package adguard

import (
	"fmt"
	// "io/ioutil"
	"io"
	"net/http"
	"time"
)

// HostURL
const HostURL string = "https://dns-int.michels.link/control"

// ADG Client -
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
func NewClient(host, username, password *string) (*ADG, error) {
	c := ADG{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default ADG URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	c.Auth = AuthStruct{
		Username: *username,
		Password: *password,
	}

	return &c, nil
}

func (c *ADG) doRequest(req *http.Request) ([]byte, error) {
	req.SetBasicAuth(c.Auth.Username, c.Auth.Password)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
