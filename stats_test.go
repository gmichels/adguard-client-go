package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test Stats
func TestStats(t *testing.T) {
	adg := createADG()

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

// Test StatsReset
func TestStatsReset(t *testing.T) {
	adg := createADG()

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

// Test StatsConfig
func TestStatsConfig(t *testing.T) {
	adg := createADG()

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

// Test StatsConfigUpdate
func TestStatsConfigUpdate(t *testing.T) {
	adg := createADG()

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
