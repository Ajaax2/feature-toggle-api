package request

type CreateFeatureToggleRequest struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Active  bool   `json:"active"`
}

func NewFeatureToggleRequest(name string, version string, active bool) (*CreateFeatureToggleRequest, error) {
	cft := CreateFeatureToggleRequest{
		Name:    name,
		Version: version,
		Active:  active,
	}

	return &cft, nil
}
