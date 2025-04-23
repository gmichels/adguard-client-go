package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test Clients
func TestClients(t *testing.T) {
	adg := createADG()

	// call the method
	result, err := adg.Clients()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure 1 client is returned
	assert.Len(t, result.Clients, 1)
}

// Test ClientsAdd
func TestClientsAdd(t *testing.T) {
	adg := createADG()

	// create a new client
	client := models.Client{Name: "Test Client Add", Ids: []string{"test-client-add"}}

	// call the method
	err := adg.ClientsAdd(client)

	// assertions
	assert.NoError(t, err)

	// cleanup: delete the client after the test
	_ = adg.ClientsDelete(models.ClientDelete{Name: "Test Client Add"})
}

// Test ClientsDelete
func TestClientsDelete(t *testing.T) {
	adg := createADG()

	// add a client to delete
	client := models.Client{Ids: []string{"test-client-delete"}, Name: "Test Client to Delete"}
	_ = adg.ClientsAdd(client)

	// delete the client
	err := adg.ClientsDelete(models.ClientDelete{Name: "Test Client to Delete"})

	// assertions
	assert.NoError(t, err)
}

// Test ClientsUpdate
func TestClientsUpdate(t *testing.T) {
	adg := createADG()

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

// Test ClientsSearch
func TestClientsSearch(t *testing.T) {
	adg := createADG()

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
