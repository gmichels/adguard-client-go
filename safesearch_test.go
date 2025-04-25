package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test SafeSearchSettings
func TestSafeSearchSettings(t *testing.T) {
	adg := testADG()

	// create a new SafeSearch configuration
	safeSearchConfig := models.SafeSearchConfig{
		Enabled:    true,
		Bing:       true,
		Duckduckgo: false,
		Ecosia:     true,
		Google:     true,
		Pixabay:    false,
		Yandex:     true,
		Youtube:    true,
	}

	// call the method
	err := adg.SafeSearchSettings(safeSearchConfig)

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling SafeSearchStatus
	result, err := adg.SafeSearchStatus()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Enabled)
	assert.True(t, result.Bing)
	assert.False(t, result.Duckduckgo)
	assert.True(t, result.Ecosia)
	assert.True(t, result.Google)
	assert.False(t, result.Pixabay)
	assert.True(t, result.Yandex)
	assert.True(t, result.Youtube)
}

// Test SafeSearchSettings - Error initializing request
func TestSafeSearchSettings_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a new SafeSearch configuration
	safeSearchConfig := models.SafeSearchConfig{
		Enabled:    true,
		Bing:       true,
		Duckduckgo: false,
		Ecosia:     true,
		Google:     true,
		Pixabay:    false,
		Yandex:     true,
		Youtube:    true,
	}

	// call the method
	err := adg.SafeSearchSettings(safeSearchConfig)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test SafeSearchSettings - Error performing request
func TestSafeSearchSettings_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a new SafeSearch configuration
	safeSearchConfig := models.SafeSearchConfig{
		Enabled:    true,
		Bing:       true,
		Duckduckgo: false,
		Ecosia:     true,
		Google:     true,
		Pixabay:    false,
		Yandex:     true,
		Youtube:    true,
	}

	// call the method
	err := adg.SafeSearchSettings(safeSearchConfig)

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test SafeSearchStatus
func TestSafeSearchStatus(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.SafeSearchStatus()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the status is valid
	assert.Condition(t, func() bool {
		return result.Enabled == true || result.Enabled == false
	})
}

// Test SafeSearchStatus - Error initializing request
func TestSafeSearchStatus_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	result, err := adg.SafeSearchStatus()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test SafeSearchStatus - Error performing request
func TestSafeSearchStatus_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	result, err := adg.SafeSearchStatus()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test SafeSearchStatus - Error unmarshaling response
func TestSafeSearchStatus_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// call the method
	result, err := adg.SafeSearchStatus()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}
