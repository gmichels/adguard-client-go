package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test InstallGetAddresses
func TestInstallGetAddresses(t *testing.T) {
	adg := testADG(true)

	// call the method
	result, err := adg.InstallGetAddresses()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the interfaces are not empty
	assert.GreaterOrEqual(t, len(result.Interfaces), 2)
	// ensure the DNS and web ports are valid
	assert.GreaterOrEqual(t, result.DnsPort, uint16(1))
	assert.GreaterOrEqual(t, result.WebPort, uint16(1))
	// ensure the version is not empty
	assert.NotEmpty(t, result.Version)
}

// Test InstallCheckConfig
func TestInstallCheckConfig(t *testing.T) {
	adg := testADG(true)

	// create a configuration request
	checkConfigRequest := models.CheckConfigRequest{
		Dns: models.CheckConfigRequestInfo{
			Ip:      "0.0.0.0",
			Port:    53,
			Autofix: true,
		},
		Web: models.CheckConfigRequestInfo{
			Ip:      "0.0.0.0",
			Port:    8080,
			Autofix: true,
		},
		SetStaticIp: false,
	}

	// call the method
	result, err := adg.InstallCheckConfig(checkConfigRequest)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

// Test InstallConfigure
func TestInstallConfigure(t *testing.T) {
	adg := testADG(true)

	// Create an initial configuration
	initialConfiguration := models.InitialConfiguration{
		Dns: models.AddressInfo{
			Ip:   "0.0.0.0",
			Port: 53,
		},
		Web: models.AddressInfo{
			Ip:   "0.0.0.0",
			Port: 8080,
		},
		Username: "admin",
		Password: "SecretP@ssw0rd",
	}

	// call the method
	err := adg.InstallConfigure(initialConfiguration)

	// assertions
	assert.NoError(t, err)
}
