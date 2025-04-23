package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test Status
func TestStatus(t *testing.T) {
	adg := createADG()

	// call the method
	result, err := adg.Status()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure protection is enabled
	assert.Equal(t, true, result.ProtectionEnabled)
	// ensure 6 DNS addresses are returned
	assert.Equal(t, 6, len(result.DnsAddresses))
}

// Test DnsInfo
func TestDnsInfo(t *testing.T) {
	adg := createADG()

	// call the method
	result, err := adg.DnsInfo()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure DNS protection is enabled
	assert.Equal(t, true, result.ProtectionEnabled)
	// ensure at least one upstream DNS is configured
	assert.GreaterOrEqual(t, len(result.UpstreamDns), 1)
}

// Test DnsConfig
func TestDnsConfig(t *testing.T) {
	adg := createADG()

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

// Test Protection
func TestProtection(t *testing.T) {
	adg := createADG()

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
	assert.Equal(t, true, result.ProtectionEnabled)
}

// Test CacheClear
func TestCacheClear(t *testing.T) {
	adg := createADG()

	// call the method
	err := adg.CacheClear()

	// assertions
	assert.NoError(t, err)
}

// Test TestUpstreamDns
func TestTestUpstreamDns(t *testing.T) {
	adg := createADG()

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

// Test VersionJson
func TestVersionJson(t *testing.T) {
	adg := createADG()

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
	assert.Equal(t, true, result.Disabled)
}

// Test Update
func TestUpdate(t *testing.T) {
	adg := createADG()

	// call the method
	err := adg.Update()

	// assertions
	assert.Error(t, err)
	// ensure a specific error message
	assert.Equal(t, "status: 400, body: /update request isn't allowed now\n", err.Error())
}

// Test Login
func TestLogin(t *testing.T) {
	adg := createADG()

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

// Test Logout
func TestLogout(t *testing.T) {
	adg := createADG()

	// call the method
	err := adg.Logout()

	// assertions
	assert.NoError(t, err)
}

// Test Profile
func TestProfile(t *testing.T) {
	adg := createADG()

	// call the method
	result, err := adg.Profile()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure the profile name is correct
	assert.Equal(t, "admin", result.Name)
}

// Test ProfileUpdate
func TestProfileUpdate(t *testing.T) {
	adg := createADG()

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
