package adguard

import (
	"fmt"
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

// Test AccessList - Error initializing request
func TestAccessList_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	result, err := adg.AccessList()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test AccessList - Error performing request
func TestAccessList_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	result, err := adg.AccessList()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test AccessList - Error unmarshaling response
func TestAccessList_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// call the method
	result, err := adg.AccessList()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
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

func TestAccessSet_MarshalError(t *testing.T) {
	adg := testADG()

	// override JSONMarshal to force an error
	orig := JSONMarshal
	JSONMarshal = func(v any) ([]byte, error) {
		return nil, fmt.Errorf("forced marshal error")
	}
	defer func() { JSONMarshal = orig }()

	err := adg.AccessSet(models.AccessList{})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "forced marshal error")
}

// Test AccessSet - Error initializing request
func TestAccessSet_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a new access list
	accessList := models.AccessList{
		AllowedClients:    []string{"192.168.1.100"},
		DisallowedClients: []string{"192.168.1.200"},
		BlockedHosts:      []string{"example.com"},
	}

	// call the method
	err := adg.AccessSet(accessList)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test AccessSet - Error performing request
func TestAccessSet_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a new access list
	accessList := models.AccessList{
		AllowedClients:    []string{"192.168.1.100"},
		DisallowedClients: []string{"192.168.1.200"},
		BlockedHosts:      []string{"example.com"},
	}

	// call the method
	err := adg.AccessSet(accessList)

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}
