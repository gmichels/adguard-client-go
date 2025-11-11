package adguard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test ParentalEnable
func TestParentalEnable(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.ParentalEnable()

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling ParentalStatus
	result, err := adg.ParentalStatus()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure parental filtering is enabled
	assert.True(t, result.Enabled)
}

// Test ParentalEnable - Error initializing request
func TestParentalEnable_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	err := adg.ParentalEnable()

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test ParentalEnable - Error performing request
func TestParentalEnable_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	err := adg.ParentalEnable()

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test ParentalDisable
func TestParentalDisable(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.ParentalDisable()

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling ParentalStatus
	result, err := adg.ParentalStatus()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure parental filtering is disabled
	assert.False(t, result.Enabled)
}

// Test ParentalDisable - Error initializing request
func TestParentalDisable_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	err := adg.ParentalDisable()

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test ParentalDisable - Error performing request
func TestParentalDisable_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	err := adg.ParentalDisable()

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test ParentalStatus
func TestParentalStatus(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.ParentalStatus()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the sensitivity level is within a valid range
	assert.GreaterOrEqual(t, result.Sensitivity, 0)
	assert.LessOrEqual(t, result.Sensitivity, 10)
}

// Test ParentalStatus - Error initializing request
func TestParentalStatus_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	result, err := adg.ParentalStatus()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test ParentalStatus - Error performing request
func TestParentalStatus_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	result, err := adg.ParentalStatus()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test ParentalStatus - Error unmarshaling response
func TestParentalStatus_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// call the method
	result, err := adg.ParentalStatus()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}
