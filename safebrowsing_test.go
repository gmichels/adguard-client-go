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
