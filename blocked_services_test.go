package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test BlockedServicesAll
func TestBlockedServicesAll(t *testing.T) {
	adg := createADG()

	// call the method
	result, err := adg.BlockedServicesAll()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure at least 10 blocked services are returned
	assert.GreaterOrEqual(t, len(result.BlockedServices), 10)
}

// Test BlockedServicesGet
func TestBlockedServicesGet(t *testing.T) {
	adg := createADG()

	// call the method
	result, err := adg.BlockedServicesGet()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure 3 blocked service IDs are returned
	assert.Equal(t, len(result.Ids), 3)
	// ensure the schedule is not empty
	assert.Equal(t, result.Schedule.Monday.Start, uint(0))
	assert.Equal(t, result.Schedule.Monday.End, uint(86340000))
}

// Test BlockedServicesUpdate
func TestBlockedServicesUpdate(t *testing.T) {
	adg := createADG()

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
