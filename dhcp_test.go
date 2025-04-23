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

// Test DhcpResetLeases
func TestDhcpResetLeases(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.DhcpResetLeases()

	// assertions
	assert.NoError(t, err)
}
