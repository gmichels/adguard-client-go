package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test AccessList
func TestAccessList(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.AccessList()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure no allowed clients are returned
	assert.Len(t, result.AllowedClients, 0)
	// ensure 2 disallowed clients are returned
	assert.Len(t, result.DisallowedClients, 2)
	// ensure 3 blocked hosts are returned
	assert.Len(t, (result.BlockedHosts), 3)
}

// Test AccessSet
func TestAccessSet(t *testing.T) {
	adg := testADG()

	// create a new access list
	accessList := models.AccessList{
		AllowedClients:    []string{"192.168.1.100"},
		DisallowedClients: []string{"192.168.1.200"},
		BlockedHosts:      []string{"example.com"},
	}

	// call the method
	err := adg.AccessSet(accessList)

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling AccessList
	result, err := adg.AccessList()
	assert.NoError(t, err)
	assert.Contains(t, result.AllowedClients, "192.168.1.100")
	assert.Contains(t, result.DisallowedClients, "192.168.1.200")
	assert.Contains(t, result.BlockedHosts, "example.com")

	// cleanup: reset the access list
	cleanupAccessList := models.AccessList{
		AllowedClients:    []string{},
		DisallowedClients: []string{},
		BlockedHosts:      []string{},
	}
	_ = adg.AccessSet(cleanupAccessList)
}
