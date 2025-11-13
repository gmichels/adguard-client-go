package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test Status
func TestStatus(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.Status()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure protection is enabled
	assert.True(t, result.ProtectionEnabled)
	// ensure 6 DNS addresses are returned
	assert.Len(t, result.DnsAddresses, 6)
}

// Test Status - Error initializing request
func TestStatus_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	result, err := adg.Status()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test Status - Error performing request
func TestStatus_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	result, err := adg.Status()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test Status - Error unmarshaling response
func TestStatus_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// call the method
	result, err := adg.Status()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test DnsInfo
func TestDnsInfo(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.DnsInfo()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure DNS protection is enabled
	assert.True(t, result.ProtectionEnabled)
	// ensure at least one upstream DNS is configured
	assert.GreaterOrEqual(t, len(result.UpstreamDns), 1)
}

// Test DnsInfo CacheEnabled parsing
func TestDnsInfo_CacheEnabled(t *testing.T) {
	adg := testADG()

	result, err := adg.DnsInfo()

	assert.NoError(t, err)
	assert.NotNil(t, result)
	// Ensure CacheEnabled field is present (default may be false)
	_ = result.CacheEnabled
}

// Test DnsInfo - Error initializing request
func TestDnsInfo_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	result, err := adg.DnsInfo()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DnsInfo - Error performing request
func TestDnsInfo_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	result, err := adg.DnsInfo()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test DnsInfo - Error unmarshaling response
func TestDnsInfo_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// call the method
	result, err := adg.DnsInfo()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test DnsConfig
func TestDnsConfig(t *testing.T) {
	adg := testADG()

	// ceate a new DNS configuration
	dnsConfig := models.DNSConfig{
		UpstreamDns:     []string{"https://dns.google/dns-query"},
		BlockingMode:    "refused",
		UpstreamTimeout: 30,
	}

	// call the method
	err := adg.DnsConfig(dnsConfig)

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling DnsInfo
	result, err := adg.DnsInfo()
	assert.NoError(t, err)
	assert.Contains(t, result.UpstreamDns, "https://dns.google/dns-query")
	assert.Equal(t, "refused", result.BlockingMode)
	assert.Equal(t, uint(30), result.UpstreamTimeout)
}

// Test DnsConfig - Error initializing request
func TestDnsConfig_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a new DNS configuration
	dnsConfig := models.DNSConfig{
		UpstreamDns:     []string{"https://dns.google/dns-query"},
		BlockingMode:    "refused",
		UpstreamTimeout: 30,
	}

	// call the method
	err := adg.DnsConfig(dnsConfig)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test DnsConfig - Error performing request
func TestDnsConfig_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a new DNS configuration
	dnsConfig := models.DNSConfig{
		UpstreamDns:     []string{"https://dns.google/dns-query"},
		BlockingMode:    "refused",
		UpstreamTimeout: 30,
	}

	// call the method
	err := adg.DnsConfig(dnsConfig)

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test Protection
func TestProtection(t *testing.T) {
	adg := testADG()

	// create a protection request
	protectionRequest := models.SetProtectionRequest{
		Enabled: true,
	}

	// call the method
	err := adg.Protection(protectionRequest)

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling Status
	result, err := adg.Status()
	assert.NoError(t, err)
	assert.True(t, result.ProtectionEnabled)
}

// Test Protection - Error initializing request
func TestProtection_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a protection request
	protectionRequest := models.SetProtectionRequest{
		Enabled: true,
	}

	// call the method
	err := adg.Protection(protectionRequest)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test Protection - Error performing request
func TestProtection_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a protection request
	protectionRequest := models.SetProtectionRequest{
		Enabled: true,
	}

	// call the method
	err := adg.Protection(protectionRequest)

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test DnsConfig and Protection - Marshal errors
func TestDnsConfigAndProtection_MarshalErrors(t *testing.T) {
	adg := testADG()
	defer forceMarshalError(t)()

	assert.Error(t, adg.DnsConfig(models.DNSConfig{}))
	assert.Error(t, adg.Protection(models.SetProtectionRequest{}))
}

// Test DnsConfig and Protection - Marshal errors
func TestDnsConfigAndProtection_MarshalErrors(t *testing.T) {
	adg := testADG()
	defer forceMarshalError(t)()

	assert.Error(t, adg.DnsConfig(models.DNSConfig{}))
	assert.Error(t, adg.Protection(models.SetProtectionRequest{}))
}

// Test CacheClear
func TestCacheClear(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.CacheClear()

	// assertions
	assert.NoError(t, err)
}

// Test CacheClear - Error initializing request
func TestCacheClear_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	err := adg.CacheClear()

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test CacheClear - Error performing request
func TestCacheClear_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	err := adg.CacheClear()

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test TestUpstreamDns
func TestTestUpstreamDns(t *testing.T) {
	adg := testADG()

	// create an upstream configuration
	upstreamsConfig := models.UpstreamsConfig{
		UpstreamDns: []string{"https://8.8.8.8/dns-query"},
	}

	// call the method
	result, err := adg.TestUpstreamDns(upstreamsConfig)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

// Test TestUpstreamDns - Error initializing request
func TestTestUpstreamDns_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create an upstream configuration
	upstreamsConfig := models.UpstreamsConfig{
		UpstreamDns: []string{"https://8.8.8.8/dns-query"},
	}

	// call the method
	result, err := adg.TestUpstreamDns(upstreamsConfig)

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test TestUpstreamDns - Error performing request
func TestTestUpstreamDns_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create an upstream configuration
	upstreamsConfig := models.UpstreamsConfig{
		UpstreamDns: []string{"https://8.8.8.8/dns-query"},
	}

	// call the method
	result, err := adg.TestUpstreamDns(upstreamsConfig)

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test TestUpstreamDns - Error unmarshaling response
func TestTestUpstreamDns_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// create an upstream configuration
	upstreamsConfig := models.UpstreamsConfig{
		UpstreamDns: []string{"https://8.8.8.8/dns-query"},
	}

	// call the method
	result, err := adg.TestUpstreamDns(upstreamsConfig)

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test VersionJson
func TestVersionJson(t *testing.T) {
	adg := testADG()

	// create a version request
	versionRequest := models.GetVersionRequest{
		RecheckNow: true,
	}

	// call the method
	result, err := adg.VersionJson(versionRequest)

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure version checking is disabled
	assert.True(t, result.Disabled)
}

// Test VersionJson - Error initializing request
func TestVersionJson_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a version request
	versionRequest := models.GetVersionRequest{
		RecheckNow: true,
	}

	// call the method
	result, err := adg.VersionJson(versionRequest)

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test VersionJson - Error performing request
func TestVersionJson_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a version request
	versionRequest := models.GetVersionRequest{
		RecheckNow: true,
	}

	// call the method
	result, err := adg.VersionJson(versionRequest)

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test VersionJson - Error unmarshaling response
func TestVersionJson_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// create a version request
	versionRequest := models.GetVersionRequest{
		RecheckNow: true,
	}

	// call the method
	result, err := adg.VersionJson(versionRequest)

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test Update
func TestUpdate(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.Update()

	// assertions
	assert.Error(t, err)
	// ensure a specific error message
	assert.Equal(t, "status: 400, body: /update request isn't allowed now\n", err.Error())
}

// Test Update - Error initializing request
func TestUpdate_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	err := adg.Update()

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test Update - Error performing request
func TestUpdate_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	err := adg.Update()

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test Login
func TestLogin(t *testing.T) {
	adg := testADG()

	// create a login request
	login := models.Login{
		Name:     "admin",
		Password: "SecretP@ssw0rd",
	}

	// call the method
	err := adg.Login(login)

	// assertions
	assert.NoError(t, err)
}

// Test Login - Error initializing request
func TestLogin_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a login request
	login := models.Login{
		Name:     "admin",
		Password: "SecretP@ssw0rd",
	}

	// call the method
	err := adg.Login(login)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test Login - Error performing request
func TestLogin_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a login request
	login := models.Login{
		Name:     "admin",
		Password: "wrongpassword",
	}

	// call the method
	err := adg.Login(login)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "status: 403")
}

// Test Logout
func TestLogout(t *testing.T) {
	adg := testADG()

	// call the method
	err := adg.Logout()

	// assertions
	assert.NoError(t, err)
}

// Test Logout - Error initializing request
func TestLogout_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	err := adg.Logout()

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test Logout - Error performing request
func TestLogout_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	err := adg.Logout()

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test Profile
func TestProfile(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.Profile()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the profile name is correct
	assert.Equal(t, "admin", result.Name)
}

// Test Profile - Error initializing request
func TestProfile_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	result, err := adg.Profile()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test Profile - Error performing request
func TestProfile_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	result, err := adg.Profile()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test Profile - Error unmarshaling response
func TestProfile_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// call the method
	result, err := adg.Profile()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test ProfileUpdate
func TestProfileUpdate(t *testing.T) {
	adg := testADG()

	// create a profile update request
	profileInfo := models.ProfileInfo{
		Name:     "admin",
		Language: "en",
		Theme:    "dark",
	}

	// call the method
	err := adg.ProfileUpdate(profileInfo)

	// assertions
	assert.NoError(t, err)

	// verify the changes by calling Profile
	result, err := adg.Profile()
	assert.NoError(t, err)
	assert.Equal(t, "dark", result.Theme)
}

// Test ProfileUpdate - Error initializing request
func TestProfileUpdate_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a profile update request
	profileInfo := models.ProfileInfo{
		Name:     "admin",
		Language: "en",
		Theme:    "dark",
	}

	// call the method
	err := adg.ProfileUpdate(profileInfo)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test ProfileUpdate - Error performing request
func TestProfileUpdate_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a profile update request
	profileInfo := models.ProfileInfo{
		Name:     "admin",
		Language: "en",
		Theme:    "dark",
	}

	// call the method
	err := adg.ProfileUpdate(profileInfo)

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 401, body: ", err.Error())
}

// Test TestUpstreamDns, VersionJson, Login, ProfileUpdate - Marshal errors
func TestTestUpstreamDns_Version_Login_Profile_MarshalErrors(t *testing.T) {
	adg := testADG()
	defer forceMarshalError(t)()

	_, err := adg.TestUpstreamDns(models.UpstreamsConfig{})
	assert.Error(t, err)
	_, err = adg.VersionJson(models.GetVersionRequest{})
	assert.Error(t, err)
	assert.Error(t, adg.Login(models.Login{}))
	assert.Error(t, adg.ProfileUpdate(models.ProfileInfo{}))
}

// Test TestUpstreamDns, VersionJson, Login, ProfileUpdate - Marshal errors
func TestTestUpstreamDns_Version_Login_Profile_MarshalErrors(t *testing.T) {
	adg := testADG()
	defer forceMarshalError(t)()

	_, err := adg.TestUpstreamDns(models.UpstreamsConfig{})
	assert.Error(t, err)
	_, err = adg.VersionJson(models.GetVersionRequest{})
	assert.Error(t, err)
	assert.Error(t, adg.Login(models.Login{}))
	assert.Error(t, adg.ProfileUpdate(models.ProfileInfo{}))
}
