package adguard

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test NewClient - Valid parameters
func TestInitNewClient_ValidParameters(t *testing.T) {
	host := "localhost"
	username := "admin"
	password := "password"
	scheme := "http"
	timeout := 5

	client, err := NewClient(&host, &username, &password, &scheme, &timeout)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "http://localhost", client.HostURL)
	assert.Equal(t, username, client.Auth.Username)
	assert.Equal(t, password, client.Auth.Password)
	assert.Equal(t, time.Duration(timeout)*time.Second, client.HTTPClient.Timeout)
}

// Test NewClient - Missing required parameters
func TestInitClient_MissingParameters(t *testing.T) {
	// case 1: Missing `host`
	host := ""
	username := "admin"
	password := "password"
	scheme := "http"
	timeout := 5

	client, err := NewClient(&host, &username, &password, &scheme, &timeout)

	// assertions
	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required parameter `host`")

	// case 2: Missing `username`
	host = "localhost"
	username = ""
	client, err = NewClient(&host, &username, &password, &scheme, &timeout)

	// assertions
	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required parameter `username`")

	// case 3: Missing `password`
	username = "admin"
	password = ""
	client, err = NewClient(&host, &username, &password, &scheme, &timeout)

	// assertions
	assert.Nil(t, client)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "required parameter `password`")

	// case 4: Missing `scheme`
	password = "password"
	scheme = ""
	client, err = NewClient(&host, &username, &password, &scheme, &timeout)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Contains(t, client.HostURL, "https://")

	// case 5: Missing `timeout`
	scheme = "http"
	client, err = NewClient(&host, &username, &password, &scheme, nil)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, client.HTTPClient.Timeout, time.Duration(10)*time.Second)
}

// Test doRequest - Successful request
func TestInitDoRequest_Success(t *testing.T) {
	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"success"}`))
	}))
	defer server.Close()

	// create a client
	host := "localhost"
	username := "admin"
	password := "password"
	scheme := "http"
	timeout := 5
	client, _ := NewClient(&host, &username, &password, &scheme, &timeout)

	// Update the client's HostURL to the mock server
	client.HostURL = server.URL

	// create a request
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/test", client.HostURL), nil)

	// call doRequest
	body, err := client.doRequest(req)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, body)
	assert.Contains(t, string(body), `"message":"success"`)
}

// Test doRequest - Request error
func TestInitDoRequest_RequestError(t *testing.T) {
	// create a client
	host := "localhost"
	username := "admin"
	password := "password"
	scheme := "http"
	timeout := 5
	client, _ := NewClient(&host, &username, &password, &scheme, &timeout)

	// create a request with an invalid URL
	req, _ := http.NewRequest("GET", "http://invalid-url", nil)

	// call doRequest
	body, err := client.doRequest(req)

	// assertions
	assert.Nil(t, body)
	assert.Error(t, err)
}

// Test doRequest - Non-200 response
func TestInitDoRequest_Non200Response(t *testing.T) {
	// create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"error":"forbidden"}`))
	}))
	defer server.Close()

	// create a client
	host := "localhost"
	username := "admin"
	password := "password"
	scheme := "http"
	timeout := 5
	client, _ := NewClient(&host, &username, &password, &scheme, &timeout)

	// Update the client's HostURL to the mock server
	client.HostURL = server.URL

	// create a request
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/test", client.HostURL), nil)

	// call doRequest
	body, err := client.doRequest(req)

	// assertions
	assert.Nil(t, body)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "status: 403")
	assert.Contains(t, err.Error(), `{"error":"forbidden"}`)
}
