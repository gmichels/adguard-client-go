package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test DhcpStatus
func TestDhcpStatus(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.DhcpStatus()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure DHCP is disabled by default
	assert.False(t, result.Enabled)
	// ensure the default interface is lo0
	assert.Equal(t, "lo0", result.InterfaceName)
	// ensure there is a DHCP range
	assert.Equal(t, "192.168.200.50", result.V4.RangeEnd)
	// Ensure 3 leases are present
	assert.Len(t, result.Leases, 3)
}

// Test DhcpStatus - Error initializing request
func TestDhcpStatus_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	result, err := adg.DhcpStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DhcpStatus - Error performing request
func TestDhcpStatus_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	result, err := adg.DhcpStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test DhcpStatus - Error unmarshaling response
func TestDhcpStatus_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	result, err := adg.DhcpStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test DhcpInterfaces
func TestDhcpInterfaces(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.DhcpInterfaces()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure at least one interface is returned
	assert.GreaterOrEqual(t, len(*result), 1)
}

// Test DhcpInterfaces - Error initializing request
func TestDhcpInterfaces_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	result, err := adg.DhcpInterfaces()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DhcpInterfaces - Error performing request
func TestDhcpInterfaces_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	result, err := adg.DhcpInterfaces()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test DhcpInterfaces - Error unmarshaling response
func TestDhcpInterfaces_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	result, err := adg.DhcpInterfaces()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test DhcpSetConfig
func TestDhcpSetConfig(t *testing.T) {
	adg := testADG()

	// create a new DHCP configuration
	dhcpConfig := models.DhcpConfig{
		Enabled:       true,
		InterfaceName: "eth0",
		V4: models.DhcpConfigV4{
			GatewayIp:     "192.168.1.1",
			SubnetMask:    "255.255.255.0",
			RangeStart:    "192.168.1.100",
			RangeEnd:      "192.168.1.200",
			LeaseDuration: 3600,
		},
	}

	// call the method
	err := adg.DhcpSetConfig(dhcpConfig)

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling DhcpStatus
	result, err := adg.DhcpStatus()
	assert.NoError(t, err)
	assert.True(t, result.Enabled)
	assert.Equal(t, "eth0", result.InterfaceName)
}

// Test DhcpSetConfig - Error initializing request
func TestDhcpSetConfig_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a new DHCP configuration
	dhcpConfig := models.DhcpConfig{
		Enabled:       true,
		InterfaceName: "eth0",
		V4: models.DhcpConfigV4{
			GatewayIp:     "192.168.1.1",
			SubnetMask:    "255.255.255.0",
			RangeStart:    "192.168.1.100",
			RangeEnd:      "192.168.1.200",
			LeaseDuration: 3600,
		},
	}

	// Call the method
	err := adg.DhcpSetConfig(dhcpConfig)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DhcpSetConfig - Error performing request
func TestDhcpSetConfig_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a new DHCP configuration
	dhcpConfig := models.DhcpConfig{
		Enabled:       true,
		InterfaceName: "eth0",
		V4: models.DhcpConfigV4{
			GatewayIp:     "192.168.1.1",
			SubnetMask:    "255.255.255.0",
			RangeStart:    "192.168.1.100",
			RangeEnd:      "192.168.1.200",
			LeaseDuration: 3600,
		},
	}

	// Call the method
	err := adg.DhcpSetConfig(dhcpConfig)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test DhcpFindActiveDhcp
func TestDhcpFindActiveDhcp(t *testing.T) {
	adg := testADG()

	// create a request to find active DHCP servers
	dhcpFindActiveReq := models.DhcpFindActiveReq{
		Interface: "eth0",
	}

	// call the method
	result, err := adg.DhcpFindActiveDhcp(dhcpFindActiveReq)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure no other DHCP server is found
	assert.Equal(t, "no", result.V4.OtherServer.Found)
}

// Test DhcpFindActiveDhcp - Error initializing request
func TestDhcpFindActiveDhcp_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a request to find active DHCP servers
	dhcpFindActiveReq := models.DhcpFindActiveReq{
		Interface: "eth0",
	}

	// Call the method
	result, err := adg.DhcpFindActiveDhcp(dhcpFindActiveReq)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DhcpFindActiveDhcp - Error performing request
func TestDhcpFindActiveDhcp_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a request to find active DHCP servers
	dhcpFindActiveReq := models.DhcpFindActiveReq{
		Interface: "eth0",
	}

	// Call the method
	result, err := adg.DhcpFindActiveDhcp(dhcpFindActiveReq)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test DhcpFindActiveDhcp - Error unmarshaling response
func TestDhcpFindActiveDhcp_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Create a request to find active DHCP servers
	dhcpFindActiveReq := models.DhcpFindActiveReq{
		Interface: "eth0",
	}

	// Call the method
	result, err := adg.DhcpFindActiveDhcp(dhcpFindActiveReq)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test DhcpAddStaticLease
func TestDhcpAddStaticLease(t *testing.T) {
	adg := testADG()

	// create a new static lease
	staticLease := models.DhcpStaticLease{
		Mac:      "00:11:22:33:44:55",
		Ip:       "192.168.1.150",
		Hostname: "test-static-lease",
	}

	// call the method
	err := adg.DhcpAddStaticLease(staticLease)

	// assertions
	assert.NoError(t, err)

	// cleanup: remove the static lease
	_ = adg.DhcpRemoveStaticLease(staticLease)
}

// Test DhcpAddStaticLease - Error initializing request
func TestDhcpAddStaticLease_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a new static lease
	staticLease := models.DhcpStaticLease{
		Mac:      "00:11:22:33:44:55",
		Ip:       "192.168.1.150",
		Hostname: "test-static-lease",
	}

	// Call the method
	err := adg.DhcpAddStaticLease(staticLease)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DhcpAddStaticLease - Error performing request
func TestDhcpAddStaticLease_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a new static lease
	staticLease := models.DhcpStaticLease{
		Mac:      "00:11:22:33:44:55",
		Ip:       "192.168.1.150",
		Hostname: "test-static-lease",
	}

	// Call the method
	err := adg.DhcpAddStaticLease(staticLease)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test DhcpRemoveStaticLease
func TestDhcpRemoveStaticLease(t *testing.T) {
	adg := testADG()

	// add a static lease to remove
	staticLease := models.DhcpStaticLease{
		Mac:      "00:11:22:33:44:55",
		Ip:       "192.168.1.150",
		Hostname: "test-static-lease",
	}
	_ = adg.DhcpAddStaticLease(staticLease)

	// call the method to remove the lease
	err := adg.DhcpRemoveStaticLease(staticLease)

	// assertions
	assert.NoError(t, err)
}

// Test DhcpRemoveStaticLease - Error initializing request
func TestDhcpRemoveStaticLease_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a static lease to remove
	staticLease := models.DhcpStaticLease{
		Mac:      "00:11:22:33:44:55",
		Ip:       "192.168.1.150",
		Hostname: "test-static-lease",
	}

	// Call the method
	err := adg.DhcpRemoveStaticLease(staticLease)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DhcpRemoveStaticLease - Error performing request
func TestDhcpRemoveStaticLease_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a static lease to remove
	staticLease := models.DhcpStaticLease{
		Mac:      "00:11:22:33:44:55",
		Ip:       "192.168.1.150",
		Hostname: "test-static-lease",
	}

	// Call the method
	err := adg.DhcpRemoveStaticLease(staticLease)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test DhcpUpdateStaticLease
func TestDhcpUpdateStaticLease(t *testing.T) {
	adg := testADG()

	// add a static lease to update
	staticLease := models.DhcpStaticLease{
		Mac:      "00:11:22:33:44:55",
		Ip:       "192.168.1.150",
		Hostname: "test-static-lease",
	}
	_ = adg.DhcpAddStaticLease(staticLease)

	// update the static lease
	updatedLease := models.DhcpStaticLease{
		Mac:      "00:11:22:33:44:55",
		Ip:       "192.168.1.151",
		Hostname: "updated-static-lease",
	}
	err := adg.DhcpUpdateStaticLease(updatedLease)

	// assertions
	assert.NoError(t, err)

	// cleanup: remove the updated lease
	_ = adg.DhcpRemoveStaticLease(updatedLease)
}

// Test DhcpUpdateStaticLease - Error initializing request
func TestDhcpUpdateStaticLease_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a static lease to update
	updatedLease := models.DhcpStaticLease{
		Mac:      "00:11:22:33:44:55",
		Ip:       "192.168.1.151",
		Hostname: "updated-static-lease",
	}

	// Call the method
	err := adg.DhcpUpdateStaticLease(updatedLease)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DhcpUpdateStaticLease - Error performing request
func TestDhcpUpdateStaticLease_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a static lease to update
	updatedLease := models.DhcpStaticLease{
		Mac:      "00:11:22:33:44:55",
		Ip:       "192.168.1.151",
		Hostname: "updated-static-lease",
	}

	// Call the method
	err := adg.DhcpUpdateStaticLease(updatedLease)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test DhcpReset
func TestDhcpReset(t *testing.T) {
	adg := testADG()

	// aall the method
	err := adg.DhcpReset()

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling DhcpStatus
	result, err := adg.DhcpStatus()
	assert.NoError(t, err)
	// ensure DHCP is disabled after reset
	assert.False(t, result.Enabled)
}

// Test DhcpReset - Error initializing request
func TestDhcpReset_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	err := adg.DhcpReset()

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DhcpReset - Error performing request
func TestDhcpReset_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	err := adg.DhcpReset()

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test DhcpResetLeases
func TestDhcpResetLeases(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.DhcpResetLeases()

	// assertions
	assert.NoError(t, err)
}

// Test DhcpResetLeases - Error initializing request
func TestDhcpResetLeases_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	err := adg.DhcpResetLeases()

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DhcpResetLeases - Error performing request
func TestDhcpResetLeases_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	err := adg.DhcpResetLeases()

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}
