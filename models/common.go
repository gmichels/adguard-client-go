package models

// Error - A generic JSON error response.
type Error struct {
	Message string `json:"message" description:"The error message, an opaque string."`
}

// Enabled - model does not formally exist in the upstream API
type Enabled struct {
	Enabled bool `json:"enabled"`
}
