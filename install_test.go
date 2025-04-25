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

// Test InstallGetAddresses - Error initializing request
func TestInstallGetAddresses_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	result, err := adg.InstallGetAddresses()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test InstallGetAddresses - Error performing request
func TestInstallGetAddresses_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError(true)

	// call the method
	result, err := adg.InstallGetAddresses()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "connect: connection refused")
}

// Test InstallGetAddresses - Error unmarshaling response
func TestInstallGetAddresses_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// call the method
	result, err := adg.InstallGetAddresses()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
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

// Test InstallCheckConfig - Error initializing request
func TestInstallCheckConfig_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a configuration request
	checkConfigRequest := models.CheckConfigRequest{
		Dns: models.CheckConfigRequestInfo{
			Ip:      "192.168.1.1",
			Port:    53,
			Autofix: true,
		},
		Web: models.CheckConfigRequestInfo{
			Ip:      "192.168.1.1",
			Port:    80,
			Autofix: true,
		},
		SetStaticIp: true,
	}

	// call the method
	result, err := adg.InstallCheckConfig(checkConfigRequest)

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test InstallCheckConfig - Error performing request
func TestInstallCheckConfig_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError(true)

	// create a configuration request
	checkConfigRequest := models.CheckConfigRequest{
		Dns: models.CheckConfigRequestInfo{
			Ip:      "192.168.1.1",
			Port:    53,
			Autofix: true,
		},
		Web: models.CheckConfigRequestInfo{
			Ip:      "192.168.1.1",
			Port:    80,
			Autofix: true,
		},
		SetStaticIp: true,
	}

	// call the method
	result, err := adg.InstallCheckConfig(checkConfigRequest)

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "connect: connection refused")
}

// Test InstallCheckConfig - Error unmarshaling response
func TestInstallCheckConfig_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// create a configuration request
	checkConfigRequest := models.CheckConfigRequest{
		Dns: models.CheckConfigRequestInfo{
			Ip:      "192.168.1.1",
			Port:    53,
			Autofix: true,
		},
		Web: models.CheckConfigRequestInfo{
			Ip:      "192.168.1.1",
			Port:    80,
			Autofix: true,
		},
		SetStaticIp: true,
	}

	// call the method
	result, err := adg.InstallCheckConfig(checkConfigRequest)

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test InstallConfigure
func TestInstallConfigure(t *testing.T) {
	adg := testADG(true)

	// create an initial configuration
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

// Test InstallConfigure - Error initializing request
func TestInstallConfigure_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create an initial configuration
	initialConfiguration := models.InitialConfiguration{
		Dns: models.AddressInfo{
			Ip:   "192.168.1.1",
			Port: 53,
		},
		Web: models.AddressInfo{
			Ip:   "192.168.1.1",
			Port: 80,
		},
		Username: "admin",
		Password: "password",
	}

	// call the method
	err := adg.InstallConfigure(initialConfiguration)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test InstallConfigure - Error performing request
func TestInstallConfigure_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError(true)

	// create an initial configuration
	initialConfiguration := models.InitialConfiguration{
		Dns: models.AddressInfo{
			Ip:   "192.168.1.1",
			Port: 53,
		},
		Web: models.AddressInfo{
			Ip:   "192.168.1.1",
			Port: 80,
		},
		Username: "admin",
		Password: "password",
	}

	// call the method
	err := adg.InstallConfigure(initialConfiguration)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "connect: connection refused")
}
