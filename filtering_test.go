package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test FilteringStatus
func TestFilteringStatus(t *testing.T) {
	adg := createADG()

	// call the method
	result, err := adg.FilteringStatus()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure filtering is enabled
	assert.True(t, result.Enabled)
	// ensure the interval is set to 24
	assert.Equal(t, uint(24), result.Interval)
	// ensure there are 2 filters
	assert.Len(t, result.Filters, 2)
	// ensure there is 1 whitelist filter
	assert.Len(t, result.WhitelistFilters, 1)
	// ensure there are 7 user rules
	assert.Len(t, result.UserRules, 7)
}

// Test FilteringConfig
func TestFilteringConfig(t *testing.T) {
	adg := createADG()

	// create a new filtering configuration
	filterConfig := models.FilterConfig{
		Enabled:  false,
		Interval: 12,
	}

	// call the method
	err := adg.FilteringConfig(filterConfig)

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling FilteringStatus
	result, err := adg.FilteringStatus()
	assert.NoError(t, err)
	assert.False(t, result.Enabled)
}

// Test FilteringAddUrl
func TestFilteringAddUrl(t *testing.T) {
	adg := createADG()

	// create a new filter URL
	filterData := models.AddUrlRequest{
		Name: "Test Filter to Add",
		Url:  "https://raw.githubusercontent.com/gmichels/terraform-provider-adguard/refs/heads/main/assets/list_filter_3.txt",
	}

	// call the method
	err := adg.FilteringAddUrl(filterData)

	// assertions
	assert.NoError(t, err)

	// cleanup: remove the filter URL
	filterDelete := models.RemoveUrlRequest{
		Url: filterData.Url,
	}
	_ = adg.FilteringRemoveUrl(filterDelete)
}

// Test FilteringRemoveUrl
func TestFilteringRemoveUrl(t *testing.T) {
	adg := createADG()

	// add a filter URL to remove
	filterData := models.AddUrlRequest{
		Name: "Test Filter to Remove",
		Url:  "https://raw.githubusercontent.com/gmichels/terraform-provider-adguard/refs/heads/main/assets/list_filter_4.txt",
	}
	_ = adg.FilteringAddUrl(filterData)

	// call the method to remove the filter URL
	filterDelete := models.RemoveUrlRequest{
		Url: filterData.Url,
	}
	err := adg.FilteringRemoveUrl(filterDelete)

	// assertions
	assert.NoError(t, err)
}

// Test FilteringSetRules
func TestFilteringSetRules(t *testing.T) {
	adg := createADG()

	// create new custom rules
	rules := models.SetRulesRequest{
		Rules: []string{"||example.com^", "@@||allowed.com^"},
	}

	// call the method
	err := adg.FilteringSetRules(rules)

	// Assertions
	assert.NoError(t, err)
}

// Test FilteringCheckHost
func TestFilteringCheckHost(t *testing.T) {
	adg := createADG()

	// Call the method
	name := "example.com"
	client := "192.168.1.100"
	qtype := "A"
	result, err := adg.FilteringCheckHost(&name, &client, &qtype)

	// Assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the host is blocked
	assert.Equal(t, "FilteredBlackList", result.Reason)
	// ensure the rule text is correct
	assert.Equal(t, "||example.com^", result.Rules[0].Text)
}

// Test FilteringRefresh
func TestFilteringRefresh(t *testing.T) {
	adg := createADG()

	// create a refresh request
	refreshRequest := models.FilterRefreshRequest{
		Whitelist: false,
	}

	// call the method
	result, err := adg.FilteringRefresh(refreshRequest)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure at least 0 filters were updated
	assert.GreaterOrEqual(t, result.Updated, 0)
}
