package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test Clients
func TestClients(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.Clients()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure 1 client is returned
	assert.Len(t, result.Clients, 1)
}

// Test Clients - Error initializing request
func TestClients_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	result, err := adg.Clients()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test Clients - Error performing request
func TestClients_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	result, err := adg.Clients()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test Clients - Error unmarshaling response
func TestClients_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	result, err := adg.Clients()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test ClientsAdd
func TestClientsAdd(t *testing.T) {
	adg := testADG()

	// create a new client
	client := models.Client{Name: "Test Client Add", Ids: []string{"test-client-add"}}

	// call the method
	err := adg.ClientsAdd(client)

	// assertions
	assert.NoError(t, err)

	// cleanup: delete the client after the test
	_ = adg.ClientsDelete(models.ClientDelete{Name: "Test Client Add"})
}

// Test ClientsAdd - Error initializing request
func TestClientsAdd_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a new client
	client := models.Client{Name: "Test Client Add", Ids: []string{"test-client-add"}}

	// Call the method
	err := adg.ClientsAdd(client)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test ClientsAdd - Error performing request
func TestClientsAdd_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a new client
	client := models.Client{Name: "Test Client Add", Ids: []string{"test-client-add"}}

	// Call the method
	err := adg.ClientsAdd(client)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test ClientsDelete
func TestClientsDelete(t *testing.T) {
	adg := testADG()

	// add a client to delete
	client := models.Client{Ids: []string{"test-client-delete"}, Name: "Test Client to Delete"}
	_ = adg.ClientsAdd(client)

	// delete the client
	err := adg.ClientsDelete(models.ClientDelete{Name: "Test Client to Delete"})

	// assertions
	assert.NoError(t, err)
}

// Test ClientsDelete - Error initializing request
func TestClientsDelete_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a client delete request
	clientDelete := models.ClientDelete{Name: "Test Client Delete"}

	// Call the method
	err := adg.ClientsDelete(clientDelete)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test ClientsDelete - Error performing request
func TestClientsDelete_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a client delete request
	clientDelete := models.ClientDelete{Name: "Test Client Delete"}

	// Call the method
	err := adg.ClientsDelete(clientDelete)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test ClientsUpdate
func TestClientsUpdate(t *testing.T) {
	adg := testADG()

	// add a client to update
	client := models.Client{Ids: []string{"test-client-update"}, Name: "Test Client to Update"}
	_ = adg.ClientsAdd(client)

	// update the client
	clientUpdate := models.ClientUpdate{Name: "Test Client to Update", Data: models.Client{Ids: []string{"updated-test-client"}, Name: "Updated Client Name"}}
	err := adg.ClientsUpdate(clientUpdate)

	// assertions
	assert.NoError(t, err)

	// cleanup: delete the client after the test
	_ = adg.ClientsDelete(models.ClientDelete{Name: "Updated Client Name"})
}

// Test ClientsUpdate - Error initializing request
func TestClientsUpdate_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a client update request
	clientUpdate := models.ClientUpdate{
		Name: "Test Client Update",
		Data: models.Client{Ids: []string{"updated-client"}, Name: "Updated Client"},
	}

	// Call the method
	err := adg.ClientsUpdate(clientUpdate)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test ClientsUpdate - Error performing request
func TestClientsUpdate_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a client update request
	clientUpdate := models.ClientUpdate{
		Name: "Test Client Update",
		Data: models.Client{Ids: []string{"updated-client"}, Name: "Updated Client"},
	}

	// Call the method
	err := adg.ClientsUpdate(clientUpdate)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test ClientsSearch
func TestClientsSearch(t *testing.T) {
	adg := testADG()

	// add a client to search
	client := models.Client{Ids: []string{"test-client-search"}, Name: "Test Client to Search"}
	_ = adg.ClientsAdd(client)

	// search for the client
	result, err := adg.ClientsSearch([]string{"test-client-search"})

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, *result, 1)
	assert.Equal(t, "Test Client to Search", (*result)[0]["test-client-search"].Name)

	// cleanup: delete the client after the test
	_ = adg.ClientsDelete(models.ClientDelete{Name: "Test Client to Search"})
}

// Test ClientsSearch - Error initializing request
func TestClientsSearch_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	result, err := adg.ClientsSearch([]string{"test-client-search"})

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test ClientsSearch - Error performing request
func TestClientsSearch_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	result, err := adg.ClientsSearch([]string{"test-client-search"})

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test ClientsSearch - Error unmarshaling response
func TestClientsSearch_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	result, err := adg.ClientsSearch([]string{"test-client-search"})

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}
