package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test Querylog
func TestQuerylog(t *testing.T) {
	adg := testADG()

	// call the method with valid parameters
	limit := 10
	responseStatus := "rewritten"
	result, err := adg.Querylog(nil, nil, &limit, nil, &responseStatus)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure 2 query log entries are returned
	assert.Len(t, result.Data, 2)
	// ensure the first entry has the expected information
	assert.Equal(t, result.Data[0].Question.Name, "example.org")
}

// Test Querylog with nil parameters
func TestQuerylogNilParams(t *testing.T) {
	adg := testADG()

	// call the method with nil parameters
	result, err := adg.Querylog(nil, nil, nil, nil, nil)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure at least 1 query log entry is returned
	assert.GreaterOrEqual(t, len(result.Data), 1)
}

// Test QuerylogClear
func TestQuerylogClear(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.QuerylogClear()

	// assertions
	assert.NoError(t, err)

	// verify the query log is cleared
	result, err := adg.Querylog(nil, nil, nil, nil, nil)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Len(t, result.Data, 0)
}

// Test QuerylogConfig
func TestQuerylogConfig(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.QuerylogConfig()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure query logging is enabled
	assert.True(t, result.Enabled)
	// ensure the interval is 4 hours
	assert.GreaterOrEqual(t, result.Interval, uint64(14400000))
	// ensure the ignored list has 3 entries
	assert.Len(t, result.Ignored, 3)
}

// Test QuerylogConfigUpdate
func TestQuerylogConfigUpdate(t *testing.T) {
	adg := testADG()

	// create a new query log configuration
	queryLogConfig := models.GetQueryLogConfigResponse{
		Enabled:           false,
		Interval:          28800000,
		AnonymizeClientIp: true,
		Ignored:           []string{"example.org", "test.org"},
	}

	// call the method
	err := adg.QuerylogConfigUpdate(queryLogConfig)

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling QuerylogConfig
	result, err := adg.QuerylogConfig()
	assert.NoError(t, err)
	assert.False(t, result.Enabled)
	assert.Equal(t, uint64(28800000), result.Interval)
	assert.True(t, result.AnonymizeClientIp)
	assert.Contains(t, result.Ignored, "example.org")
	assert.Contains(t, result.Ignored, "test.org")
}
