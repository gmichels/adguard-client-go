package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test BlockedServicesAll
func TestBlockedServicesAll(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.BlockedServicesAll()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure at least 10 blocked services are returned
	assert.GreaterOrEqual(t, len(result.BlockedServices), 10)
}

// Test BlockedServicesAll - Error initializing request
func TestBlockedServicesAll_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	result, err := adg.BlockedServicesAll()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test BlockedServicesAll - Error performing request
func TestBlockedServicesAll_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	result, err := adg.BlockedServicesAll()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test BlockedServicesAll - Error unmarshaling response
func TestBlockedServicesAll_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// call the method
	result, err := adg.BlockedServicesAll()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test BlockedServicesGet
func TestBlockedServicesGet(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.BlockedServicesGet()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure 3 blocked service IDs are returned
	assert.Len(t, result.Ids, 3)
	// ensure the schedule is not empty
	assert.Equal(t, result.Schedule.Monday.Start, uint(0))
	assert.Equal(t, result.Schedule.Monday.End, uint(86340000))
}

// Test BlockedServicesGet - Error initializing request
func TestBlockedServicesGet_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	result, err := adg.BlockedServicesGet()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test BlockedServicesGet - Error performing request
func TestBlockedServicesGet_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	result, err := adg.BlockedServicesGet()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test BlockedServicesGet - Error unmarshaling response
func TestBlockedServicesGet_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// call the method
	result, err := adg.BlockedServicesGet()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test BlockedServicesUpdate
func TestBlockedServicesUpdate(t *testing.T) {
	adg := testADG()

	// create a new blocked services schedule
	blockedServicesSchedule := models.BlockedServicesSchedule{
		Ids: []string{"facebook", "youtube"},
	}

	// call the method
	err := adg.BlockedServicesUpdate(blockedServicesSchedule)

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling BlockedServicesGet
	result, err := adg.BlockedServicesGet()
	assert.NoError(t, err)
	assert.Contains(t, result.Ids, "facebook")
	assert.Contains(t, result.Ids, "youtube")

	// cleanup: reset the blocked services schedule
	cleanupBlockedServicesSchedule := models.BlockedServicesSchedule{
		Ids: []string{},
	}
	_ = adg.BlockedServicesUpdate(cleanupBlockedServicesSchedule)
}

// Test BlockedServicesUpdate - Error initializing request
func TestBlockedServicesUpdate_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a new blocked services schedule
	blockedServicesSchedule := models.BlockedServicesSchedule{
		Ids: []string{"facebook", "youtube"},
	}

	// call the method
	err := adg.BlockedServicesUpdate(blockedServicesSchedule)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test BlockedServicesUpdate - Error performing request
func TestBlockedServicesUpdate_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a new blocked services schedule
	blockedServicesSchedule := models.BlockedServicesSchedule{
		Ids: []string{"facebook", "youtube"},
	}

	// call the method
	err := adg.BlockedServicesUpdate(blockedServicesSchedule)

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test BlockedServicesUpdate - Marshal error
func TestBlockedServicesUpdate_MarshalError(t *testing.T) {
	adg := testADG()
	defer forceMarshalError(t)()

	err := adg.BlockedServicesUpdate(models.BlockedServicesSchedule{})
	assert.Error(t, err)
}

// Test BlockedServicesAll contains groups and group ids
func TestBlockedServicesAll_Groups(t *testing.T) {
	adg := testADG()

	result, err := adg.BlockedServicesAll()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	// groups should be present
	assert.GreaterOrEqual(t, len(result.Groups), 1)
	// each blocked service may have a group_id (optional)
	for _, svc := range result.BlockedServices {
		// group id may be empty; ensure field exists by checking zero value
		_ = svc.GroupId
	}
}
