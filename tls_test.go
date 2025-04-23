package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test TlsStatus
func TestTlsStatus(t *testing.T) {
	adg := createADG()

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

// Test TlsConfigure
func TestTlsConfigure(t *testing.T) {
	adg := createADG()

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

// Test TlsValidate
func TestTlsValidate(t *testing.T) {
	adg := createADG()

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
