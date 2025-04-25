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

// Test Querylog - Error initializing request
func TestQuerylog_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method with valid parameters
	limit := 10
	responseStatus := "rewritten"
	result, err := adg.Querylog(nil, nil, &limit, nil, &responseStatus)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test Querylog - Error performing request
func TestQuerylog_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method with valid parameters
	limit := 10
	responseStatus := "rewritten"
	result, err := adg.Querylog(nil, nil, &limit, nil, &responseStatus)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test Querylog - Error unmarshaling response
func TestQuerylog_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method with valid parameters
	limit := 10
	responseStatus := "rewritten"
	result, err := adg.Querylog(nil, nil, &limit, nil, &responseStatus)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
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

// Test QuerylogClear - Error initializing request
func TestQuerylogClear_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	err := adg.QuerylogClear()

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test QuerylogClear - Error performing request
func TestQuerylogClear_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	err := adg.QuerylogClear()

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
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

// Test QuerylogConfig - Error initializing request
func TestQuerylogConfig_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	result, err := adg.QuerylogConfig()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test QuerylogConfig - Error performing request
func TestQuerylogConfig_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	result, err := adg.QuerylogConfig()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test QuerylogConfig - Error unmarshaling response
func TestQuerylogConfig_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	result, err := adg.QuerylogConfig()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
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

// Test QuerylogConfigUpdate - Error initializing request
func TestQuerylogConfigUpdate_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a new query log configuration
	queryLogConfig := models.GetQueryLogConfigResponse{
		Enabled:           false,
		Interval:          28800000,
		AnonymizeClientIp: true,
		Ignored:           []string{"example.org", "test.org"},
	}

	// Call the method
	err := adg.QuerylogConfigUpdate(queryLogConfig)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test QuerylogConfigUpdate - Error performing request
func TestQuerylogConfigUpdate_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a new query log configuration
	queryLogConfig := models.GetQueryLogConfigResponse{
		Enabled:           false,
		Interval:          28800000,
		AnonymizeClientIp: true,
		Ignored:           []string{"example.org", "test.org"},
	}

	// Call the method
	err := adg.QuerylogConfigUpdate(queryLogConfig)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}
