package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test FilteringStatus
func TestFilteringStatus(t *testing.T) {
	adg := testADG()

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

// Test FilteringStatus - Error initializing request
func TestFilteringStatus_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	result, err := adg.FilteringStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test FilteringStatus - Error performing request
func TestFilteringStatus_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	result, err := adg.FilteringStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test FilteringStatus - Error unmarshaling response
func TestFilteringStatus_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	result, err := adg.FilteringStatus()

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test FilteringConfig
func TestFilteringConfig(t *testing.T) {
	adg := testADG()

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

// Test FilteringConfig - Error initializing request
func TestFilteringConfig_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a new filtering configuration
	filterConfig := models.FilterConfig{
		Enabled:  false,
		Interval: 12,
	}

	// Call the method
	err := adg.FilteringConfig(filterConfig)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test FilteringConfig - Error performing request
func TestFilteringConfig_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a new filtering configuration
	filterConfig := models.FilterConfig{
		Enabled:  false,
		Interval: 12,
	}

	// Call the method
	err := adg.FilteringConfig(filterConfig)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test FilteringAddUrl
func TestFilteringAddUrl(t *testing.T) {
	adg := testADG()

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

// Test FilteringAddUrl - Error initializing request
func TestFilteringAddUrl_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a new filter URL
	filterData := models.AddUrlRequest{
		Name: "Test Filter to Add",
		Url:  "https://example.com/filter.txt",
	}

	// Call the method
	err := adg.FilteringAddUrl(filterData)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test FilteringAddUrl - Error performing request
func TestFilteringAddUrl_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a new filter URL
	filterData := models.AddUrlRequest{
		Name: "Test Filter to Add",
		Url:  "https://example.com/filter.txt",
	}

	// Call the method
	err := adg.FilteringAddUrl(filterData)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test FilteringRemoveUrl
func TestFilteringRemoveUrl(t *testing.T) {
	adg := testADG()

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

// Test FilteringRemoveUrl - Error initializing request
func TestFilteringRemoveUrl_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a filter URL to remove
	filterDelete := models.RemoveUrlRequest{
		Url: "https://example.com/filter.txt",
	}

	// Call the method
	err := adg.FilteringRemoveUrl(filterDelete)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test FilteringRemoveUrl - Error performing request
func TestFilteringRemoveUrl_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a filter URL to remove
	filterDelete := models.RemoveUrlRequest{
		Url: "https://example.com/filter.txt",
	}

	// Call the method
	err := adg.FilteringRemoveUrl(filterDelete)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test FilteringSetUrl
func TestFilteringSetUrl(t *testing.T) {
	adg := testADG()

	// add a filter URL to update
	filterData := models.AddUrlRequest{
		Name: "Test Filter to Update",
		Url:  "https://raw.githubusercontent.com/gmichels/terraform-provider-adguard/refs/heads/main/assets/list_filter_4.txt",
	}
	_ = adg.FilteringAddUrl(filterData)

	// update to a new filter URL configuration
	filterSetUrl := models.FilterSetUrl{
		Data: models.FilterSetUrlData{
			Enabled: true,
			Name:    "Test Filter Updated",
			Url:     "https://raw.githubusercontent.com/gmichels/terraform-provider-adguard/refs/heads/main/assets/list_filter_3.txt",
		},
		Url:       filterData.Url,
		Whitelist: false,
	}

	// Call the method
	err := adg.FilteringSetUrl(filterSetUrl)

	// Assertions
	assert.NoError(t, err)
}

// Test FilteringSetUrl - Error initializing request
func TestFilteringSetUrl_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a new filter URL configuration
	filterSetUrl := models.FilterSetUrl{
		Data: models.FilterSetUrlData{
			Enabled: true,
			Name:    "Test Filter to Set",
			Url:     "https://example.com/filter.txt",
		},
		Url:       "https://example.com/filter.txt",
		Whitelist: false,
	}

	// Call the method
	err := adg.FilteringSetUrl(filterSetUrl)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test FilteringSetUrl - Error performing request
func TestFilteringSetUrl_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a new filter URL configuration
	filterSetUrl := models.FilterSetUrl{
		Data: models.FilterSetUrlData{
			Enabled: true,
			Name:    "Test Filter to Set",
			Url:     "https://example.com/filter.txt",
		},
		Url:       "https://example.com/filter.txt",
		Whitelist: false,
	}

	// Call the method
	err := adg.FilteringSetUrl(filterSetUrl)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test FilteringSetRules
func TestFilteringSetRules(t *testing.T) {
	adg := testADG()

	// create new custom rules
	rules := models.SetRulesRequest{
		Rules: []string{"||example.com^", "@@||allowed.com^"},
	}

	// call the method
	err := adg.FilteringSetRules(rules)

	// Assertions
	assert.NoError(t, err)
}

// Test FilteringSetRules - Error initializing request
func TestFilteringSetRules_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a new set of rules
	rules := models.SetRulesRequest{
		Rules: []string{"||example.com^", "@@||allowed.com^"},
	}

	// Call the method
	err := adg.FilteringSetRules(rules)

	// Assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test FilteringSetRules - Error performing request
func TestFilteringSetRules_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a new set of rules
	rules := models.SetRulesRequest{
		Rules: []string{"||example.com^", "@@||allowed.com^"},
	}

	// Call the method
	err := adg.FilteringSetRules(rules)

	// Assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test FilteringCheckHost
func TestFilteringCheckHost(t *testing.T) {
	adg := testADG()

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

// Test FilteringCheckHost - Error initializing request
func TestFilteringCheckHost_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Call the method
	name := "example.com"
	client := "192.168.1.100"
	qtype := "A"
	result, err := adg.FilteringCheckHost(&name, &client, &qtype)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test FilteringCheckHost - Error performing request
func TestFilteringCheckHost_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Call the method
	name := "example.com"
	client := "192.168.1.100"
	qtype := "A"
	result, err := adg.FilteringCheckHost(&name, &client, &qtype)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test FilteringCheckHost - Error unmarshaling response
func TestFilteringCheckHost_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Call the method
	name := "example.com"
	client := "192.168.1.100"
	qtype := "A"
	result, err := adg.FilteringCheckHost(&name, &client, &qtype)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test FilteringRefresh
func TestFilteringRefresh(t *testing.T) {
	adg := testADG()

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

// Test FilteringRefresh - Error initializing request
func TestFilteringRefresh_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// Create a refresh request
	refreshRequest := models.FilterRefreshRequest{
		Whitelist: false,
	}

	// Call the method
	result, err := adg.FilteringRefresh(refreshRequest)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test FilteringRefresh - Error performing request
func TestFilteringRefresh_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// Create a refresh request
	refreshRequest := models.FilterRefreshRequest{
		Whitelist: false,
	}

	// Call the method
	result, err := adg.FilteringRefresh(refreshRequest)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test FilteringRefresh - Error unmarshaling response
func TestFilteringRefresh_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// Create a refresh request
	refreshRequest := models.FilterRefreshRequest{
		Whitelist: false,
	}

	// Call the method
	result, err := adg.FilteringRefresh(refreshRequest)

	// Assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}
