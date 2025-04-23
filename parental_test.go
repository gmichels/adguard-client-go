package adguard

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test ParentalEnable
func TestParentalEnable(t *testing.T) {
	adg := createADG()

	// call the method
	err := adg.ParentalEnable()

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling ParentalStatus
	result, err := adg.ParentalStatus()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure parental filtering is enabled
	assert.Equal(t, true, result.Enabled)
}

// Test ParentalDisable
func TestParentalDisable(t *testing.T) {
	adg := createADG()

	// call the method
	err := adg.ParentalDisable()

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling ParentalStatus
	result, err := adg.ParentalStatus()
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure parental filtering is disabled
	assert.Equal(t, false, result.Enabled)
}

// Test ParentalStatus
func TestParentalStatus(t *testing.T) {
	adg := createADG()

	// call the method
	result, err := adg.ParentalStatus()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the sensitivity level is within a valid range
	assert.GreaterOrEqual(t, result.Sensitivity, 0)
	assert.LessOrEqual(t, result.Sensitivity, 10)
}
