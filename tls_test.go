package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test TlsStatus
func TestTlsStatus(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.TlsStatus()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the TLS configuration is valid
	assert.Condition(t, func() bool {
		return result.Enabled == true || result.Enabled == false
	})
	// ensure the certificate path is not empty
	assert.NotEmpty(t, result.CertificatePath)
}

// Test TlsStatus - Error initializing request
func TestTlsStatus_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	result, err := adg.TlsStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test TlsStatus - Error performing request
func TestTlsStatus_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	result, err := adg.TlsStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test TlsStatus - Error unmarshaling response
func TestTlsStatus_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	result, err := adg.TlsStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test TlsConfigure
func TestTlsConfigure(t *testing.T) {
	adg := testADG()

	// create a new TLS configuration
	tlsConfig := models.TlsConfig{
		Enabled:         false,
		ServerName:      "Test AdGuard Home Modified",
		CertificatePath: "/opt/adguardhome/ssl/ca.crt",
		PrivateKeyPath:  "/opt/adguardhome/ssl/ca.key",
		ServePlainDns:   true,
	}

	// call the method
	result, err := adg.TlsConfigure(tlsConfig)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Enabled)
	assert.Equal(t, "Test AdGuard Home Modified", result.ServerName)
}

// Test TlsConfigure - Error initializing request
func TestTlsConfigure_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a new TLS configuration
	tlsConfig := models.TlsConfig{
		Enabled:         false,
		ServerName:      "Test AdGuard Home Modified",
		CertificatePath: "/opt/adguardhome/ssl/ca.crt",
		PrivateKeyPath:  "/opt/adguardhome/ssl/ca.key",
		ServePlainDns:   true,
	}

	// Call the method
	result, err := adg.TlsConfigure(tlsConfig)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test TlsConfigure - Error performing request
func TestTlsConfigure_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a new TLS configuration
	tlsConfig := models.TlsConfig{
		Enabled:         false,
		ServerName:      "Test AdGuard Home Modified",
		CertificatePath: "/opt/adguardhome/ssl/ca.crt",
		PrivateKeyPath:  "/opt/adguardhome/ssl/ca.key",
		ServePlainDns:   true,
	}

	// Call the method
	result, err := adg.TlsConfigure(tlsConfig)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test TlsConfigure - Error unmarshaling response
func TestTlsConfigure_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Create a new TLS configuration
	tlsConfig := models.TlsConfig{
		Enabled:         false,
		ServerName:      "Test AdGuard Home Modified",
		CertificatePath: "/opt/adguardhome/ssl/ca.crt",
		PrivateKeyPath:  "/opt/adguardhome/ssl/ca.key",
		ServePlainDns:   true,
	}

	// Call the method
	result, err := adg.TlsConfigure(tlsConfig)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test TlsValidate
func TestTlsValidate(t *testing.T) {
	adg := testADG()

	// create a TLS configuration to validate
	tlsConfig := models.TlsConfig{
		Enabled:         true,
		CertificatePath: "/opt/adguardhome/ssl/server.crt",
		PrivateKeyPath:  "/opt/adguardhome/ssl/server.key",
	}

	// call the method
	result, err := adg.TlsValidate(tlsConfig)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Enabled)
	assert.True(t, result.ValidCert)
	assert.False(t, result.ValidChain)
	assert.Equal(t, "/opt/adguardhome/ssl/server.crt", result.CertificatePath)
	assert.Equal(t, "/opt/adguardhome/ssl/server.key", result.PrivateKeyPath)
}

// Test TlsValidate - Error initializing request
func TestTlsValidate_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a TLS configuration to validate
	tlsConfig := models.TlsConfig{
		Enabled:         true,
		CertificatePath: "/opt/adguardhome/ssl/server.crt",
		PrivateKeyPath:  "/opt/adguardhome/ssl/server.key",
	}

	// Call the method
	result, err := adg.TlsValidate(tlsConfig)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test TlsValidate - Error performing request
func TestTlsValidate_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a TLS configuration to validate
	tlsConfig := models.TlsConfig{
		Enabled:         true,
		CertificatePath: "/opt/adguardhome/ssl/server.crt",
		PrivateKeyPath:  "/opt/adguardhome/ssl/server.key",
	}

	// Call the method
	result, err := adg.TlsValidate(tlsConfig)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test TlsValidate - Error unmarshaling response
func TestTlsValidate_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Create a TLS configuration to validate
	tlsConfig := models.TlsConfig{
		Enabled:         true,
		CertificatePath: "/opt/adguardhome/ssl/server.crt",
		PrivateKeyPath:  "/opt/adguardhome/ssl/server.key",
	}

	// Call the method
	result, err := adg.TlsValidate(tlsConfig)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}
