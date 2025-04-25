package adguard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test SafeBrowsingEnable
func TestSafeBrowsingEnable(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.SafeBrowsingEnable()

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling SafeBrowsingStatus
	result, err := adg.SafeBrowsingStatus()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure safebrowsing is enabled
	assert.True(t, result.Enabled)
}

// Test SafeBrowsingEnable - Error initializing request
func TestSafeBrowsingEnable_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	err := adg.SafeBrowsingEnable()

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test SafeBrowsingEnable - Error performing request
func TestSafeBrowsingEnable_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	err := adg.SafeBrowsingEnable()

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test SafeBrowsingDisable
func TestSafeBrowsingDisable(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.SafeBrowsingDisable()

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling SafeBrowsingStatus
	result, err := adg.SafeBrowsingStatus()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure safebrowsing is disabled
	assert.False(t, result.Enabled)
}

// Test SafeBrowsingDisable - Error initializing request
func TestSafeBrowsingDisable_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	err := adg.SafeBrowsingDisable()

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test SafeBrowsingDisable - Error performing request
func TestSafeBrowsingDisable_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	err := adg.SafeBrowsingDisable()

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test SafeBrowsingStatus
func TestSafeBrowsingStatus(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.SafeBrowsingStatus()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the status is valid (either true or false)
	assert.Condition(t, func() bool {
		return result.Enabled == true || result.Enabled == false
	})
}

// Test SafeBrowsingStatus - Error initializing request
func TestSafeBrowsingStatus_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	result, err := adg.SafeBrowsingStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test SafeBrowsingStatus - Error performing request
func TestSafeBrowsingStatus_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	result, err := adg.SafeBrowsingStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test SafeBrowsingStatus - Error unmarshaling response
func TestSafeBrowsingStatus_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	result, err := adg.SafeBrowsingStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}
