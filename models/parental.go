package models

// ParentalStatus - model does not formally exist in the upstream API
type ParentalStatus struct {
	Enabled     bool `json:"enabled"`
	Sensitivity int  `json:"sensitivity"`
}
