package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test Stats
func TestStats(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.Stats()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the number of DNS queries is valid
	assert.GreaterOrEqual(t, result.NumDnsQueries, 0)
	// ensure the number of blocked filtering requests is valid
	assert.GreaterOrEqual(t, result.NumBlockedFiltering, 0)
}

// Test Stats - Error initializing request
func TestStats_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	result, err := adg.Stats()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test Stats - Error performing request
func TestStats_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	result, err := adg.Stats()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test Stats - Error unmarshaling response
func TestStats_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	result, err := adg.Stats()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test StatsReset
func TestStatsReset(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.StatsReset()

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling Stats
	result, err := adg.Stats()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure all statistics are reset to zero
	assert.Equal(t, 0, result.NumDnsQueries)
	assert.Equal(t, 0, result.NumBlockedFiltering)
}

// Test StatsReset - Error initializing request
func TestStatsReset_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	err := adg.StatsReset()

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test StatsReset - Error performing request
func TestStatsReset_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	err := adg.StatsReset()

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test StatsConfig
func TestStatsConfig(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.StatsConfig()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the configuration is valid
	assert.Condition(t, func() bool {
		return result.Enabled == true || result.Enabled == false
	})
	// ensure the interval is valid
	assert.GreaterOrEqual(t, result.Interval, uint64(28800000))
	// ensure the ignored list has 3 entries
	assert.Len(t, result.Ignored, 3)
}

// Test StatsConfig - Error initializing request
func TestStatsConfig_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	result, err := adg.StatsConfig()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test StatsConfig - Error performing request
func TestStatsConfig_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	result, err := adg.StatsConfig()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test StatsConfig - Error unmarshaling response
func TestStatsConfig_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	result, err := adg.StatsConfig()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test StatsConfigUpdate
func TestStatsConfigUpdate(t *testing.T) {
	adg := testADG()

	// create a new statistics configuration
	statsConfig := models.GetStatsConfigResponse{
		Enabled:  false,
		Interval: 3600000, // 1 hour in milliseconds
		Ignored:  []string{"example.com", "test.com"},
	}

	// call the method
	err := adg.StatsConfigUpdate(statsConfig)

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling StatsConfig
	result, err := adg.StatsConfig()
	assert.NoError(t, err)
	assert.False(t, result.Enabled)
	assert.Equal(t, uint64(3600000), result.Interval)
	assert.Contains(t, result.Ignored, "example.com")
	assert.Contains(t, result.Ignored, "test.com")
}

// Test StatsConfigUpdate - Error initializing request
func TestStatsConfigUpdate_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a new statistics configuration
	statsConfig := models.GetStatsConfigResponse{
		Enabled:  false,
		Interval: 3600000, // 1 hour in milliseconds
		Ignored:  []string{"example.com", "test.com"},
	}

	// Call the method
	err := adg.StatsConfigUpdate(statsConfig)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test StatsConfigUpdate - Error performing request
func TestStatsConfigUpdate_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a new statistics configuration
	statsConfig := models.GetStatsConfigResponse{
		Enabled:  false,
		Interval: 3600000, // 1 hour in milliseconds
		Ignored:  []string{"example.com", "test.com"},
	}

	// Call the method
	err := adg.StatsConfigUpdate(statsConfig)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}
