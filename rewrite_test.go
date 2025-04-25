package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test RewriteList
func TestRewriteList(t *testing.T) {
	adg := testADG()

	// call the method
	result, err := adg.RewriteList()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure at least 1 rewrite rule is returned
	assert.Len(t, *result, 1)
}

// Test RewriteList - Error initializing request
func TestRewriteList_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// call the method
	result, err := adg.RewriteList()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test RewriteList - Error performing request
func TestRewriteList_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// call the method
	result, err := adg.RewriteList()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test RewriteList - Error unmarshaling response
func TestRewriteList_InvalidJSONError(t *testing.T) {
	adg, server := testADGWithInvalidJSON(t)
	defer server.Close()

	// call the method
	result, err := adg.RewriteList()

	// assertions
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected end of JSON input")
}

// Test RewriteAdd
func TestRewriteAdd(t *testing.T) {
	adg := testADG()

	// create a new rewrite rule
	rewriteEntry := models.RewriteEntry{
		Domain: "example.xyz",
		Answer: "4.3.2.1",
	}

	// call the method
	err := adg.RewriteAdd(rewriteEntry)

	// assertions
	assert.NoError(t, err)
}

// Test RewriteAdd - Error initializing request
func TestRewriteAdd_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a new rewrite rule
	rewriteEntry := models.RewriteEntry{
		Domain: "example.xyz",
		Answer: "4.3.2.1",
	}

	// call the method
	err := adg.RewriteAdd(rewriteEntry)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test RewriteAdd - Error performing request
func TestRewriteAdd_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a new rewrite rule
	rewriteEntry := models.RewriteEntry{
		Domain: "example.xyz",
		Answer: "4.3.2.1",
	}

	// call the method
	err := adg.RewriteAdd(rewriteEntry)

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test RewriteDelete
func TestRewriteDelete(t *testing.T) {
	adg := testADG()

	// add a rewrite rule to delete
	rewriteEntry := models.RewriteEntry{
		Domain: "example.abc",
		Answer: "4.2.3.1",
	}
	_ = adg.RewriteAdd(rewriteEntry)

	// call the method to delete the rewrite rule
	err := adg.RewriteDelete(rewriteEntry)

	// assertions
	assert.NoError(t, err)
}

// Test RewriteDelete - Error performing request
func TestRewriteDelete_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a rewrite rule to delete
	rewriteEntry := models.RewriteEntry{
		Domain: "example.abc",
		Answer: "4.2.3.1",
	}

	// call the method
	err := adg.RewriteDelete(rewriteEntry)

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}

// Test RewriteUpdate
func TestRewriteUpdate(t *testing.T) {
	adg := testADG()

	// add a rewrite rule to update
	rewriteEntry := models.RewriteEntry{
		Domain: "example.io",
		Answer: "2.1.4.3",
	}
	_ = adg.RewriteAdd(rewriteEntry)

	// update the rewrite rule
	rewriteUpdate := models.RewriteUpdate{
		Target: rewriteEntry,
		Update: models.RewriteEntry{
			Domain: "example2.io",
			Answer: "5.6.7.8",
		},
	}
	err := adg.RewriteUpdate(rewriteUpdate)

	// assertions
	assert.NoError(t, err)
}

// Test RewriteDelete - Error initializing request
func TestRewriteDelete_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a rewrite rule to delete
	rewriteEntry := models.RewriteEntry{
		Domain: "example.abc",
		Answer: "4.2.3.1",
	}

	// call the method
	err := adg.RewriteDelete(rewriteEntry)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test RewriteUpdate - Error initializing request
func TestRewriteUpdate_NewRequestError(t *testing.T) {
	adg := testADGWithNewRequestError()

	// create a rewrite rule to update
	rewriteUpdate := models.RewriteUpdate{
		Target: models.RewriteEntry{
			Domain: "example.io",
			Answer: "2.1.4.3",
		},
		Update: models.RewriteEntry{
			Domain: "example2.io",
			Answer: "5.6.7.8",
		},
	}

	// call the method
	err := adg.RewriteUpdate(rewriteUpdate)

	// assertions
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid URL")
}

// Test RewriteUpdate - Error performing request
func TestRewriteUpdate_DoRequestError(t *testing.T) {
	adg := testADGWithDoRequestError()

	// create a rewrite rule to update
	rewriteUpdate := models.RewriteUpdate{
		Target: models.RewriteEntry{
			Domain: "example.io",
			Answer: "2.1.4.3",
		},
		Update: models.RewriteEntry{
			Domain: "example2.io",
			Answer: "5.6.7.8",
		},
	}

	// call the method
	err := adg.RewriteUpdate(rewriteUpdate)

	// assertions
	assert.Error(t, err)
	assert.Equal(t, "status: 403, body: Forbidden", err.Error())
}
