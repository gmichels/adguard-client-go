package adguard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test AppleDohMobileconfig
func TestAppleDohMobileconfig(t *testing.T) {
	adg := testADG()

	// call the method with valid parameters
	host := "example.com"
	clientId := "test-client-id"
	result, err := adg.AppleDohMobileconfig(&host, &clientId)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the response is in XML format
	assert.Contains(t, result, "<?xml")
	// ensure the host is included in the response
	assert.Contains(t, result, "example.com")
	// ensure the client ID is included in the response
	assert.Contains(t, result, "test-client-id")
}

// Test AppleDohMobileconfig with nil parameters
func TestAppleDohMobileconfigNilParams(t *testing.T) {
	adg := testADG()

	// call the method with nil parameters
	result, err := adg.AppleDohMobileconfig(nil, nil)

	// assertions
	assert.Error(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, err.Error(), "status: 500, body: {\"message\":\"no host in query parameters and no server_name\"}")
}

// Test AppleDohMobileconfig - Error initializing request
func TestAppleDohMobileconfig_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method with valid parameters
	host := "example.com"
	clientId := "test-client-id"
	result, err := adg.AppleDohMobileconfig(&host, &clientId)

	// assertions
	assert.Empty(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test AppleDohMobileconfig - Error performing request
func TestAppleDohMobileconfig_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError(true)

	// call the method with valid parameters
	host := "example.com"
	clientId := "test-client-id"
	result, err := adg.AppleDohMobileconfig(&host, &clientId)

	// assertions
	assert.Empty(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "connect: connection refused")
}

// Test AppleDotMobileconfig
func TestAppleDotMobileconfig(t *testing.T) {
	adg := testADG()

	// call the method with valid parameters
	host := "example.com"
	clientId := "test-client-id"
	result, err := adg.AppleDotMobileconfig(&host, &clientId)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the response is in XML format
	assert.Contains(t, result, "<?xml")
	// ensure the host is included in the response
	assert.Contains(t, result, "example.com")
	// ensure the client ID is included in the response
	assert.Contains(t, result, "test-client-id")
}

// Test AppleDotMobileconfig with nil parameters
func TestAppleDotMobileconfigNilParams(t *testing.T) {
	adg := testADG()

	// call the method with nil parameters
	result, err := adg.AppleDotMobileconfig(nil, nil)

	// assertions
	assert.Error(t, err)
	assert.NotNil(t, result)
	assert.Contains(t, err.Error(), "status: 500, body: {\"message\":\"no host in query parameters and no server_name\"}")
}

// Test AppleDotMobileconfig - Error initializing request
func TestAppleDotMobileconfig_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method with valid parameters
	host := "example.com"
	clientId := "test-client-id"
	result, err := adg.AppleDotMobileconfig(&host, &clientId)

	// assertions
	assert.Empty(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test AppleDotMobileconfig - Error performing request
func TestAppleDotMobileconfig_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError(true)

	// call the method with valid parameters
	host := "example.com"
	clientId := "test-client-id"
	result, err := adg.AppleDotMobileconfig(&host, &clientId)

	// assertions
	assert.Empty(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "connect: connection refused")
}
