package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test SafeSearchSettings
func TestSafeSearchSettings(t *testing.T) {
	adg := createADG()

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
	assert.Equal(t, true, result.Enabled)
	assert.Equal(t, true, result.Bing)
	assert.Equal(t, false, result.Duckduckgo)
	assert.Equal(t, true, result.Ecosia)
	assert.Equal(t, true, result.Google)
	assert.Equal(t, false, result.Pixabay)
	assert.Equal(t, true, result.Yandex)
	assert.Equal(t, true, result.Youtube)
}

// Test SafeSearchStatus
func TestSafeSearchStatus(t *testing.T) {
	adg := createADG()

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
