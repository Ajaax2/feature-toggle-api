package request

type UpdateFeatureToggleRequest struct {
	ID     string `json:"id"`
	Active bool   `json:"active"`
}
