package adguard

import (
	"testing"

	"github.com/gmichels/adguard-client-go/models"
	"github.com/stretchr/testify/assert"
)

// Test RewriteList
func TestRewriteList(t *testing.T) {
	adg := createADG()

	// call the method
	result, err := adg.RewriteList()

	// assertions
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// ensure at least 1 rewrite rule is returned
	assert.Len(t, *result, 1)
}

// Test RewriteAdd
func TestRewriteAdd(t *testing.T) {
	adg := createADG()

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

// Test RewriteDelete
func TestRewriteDelete(t *testing.T) {
	adg := createADG()

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

// Test RewriteUpdate
func TestRewriteUpdate(t *testing.T) {
	adg := createADG()

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
