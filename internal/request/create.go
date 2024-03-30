package request

type CreateFeatureToggleRequest struct {
	Name    string          `json:"name"`
	Version *VersionRequest `json:"version"`
	Active  bool            `json:"active"`
}

func NewFeatureToggleRequest(name string, version *VersionRequest, active bool) (*CreateFeatureToggleRequest, error) {
	cft := CreateFeatureToggleRequest{
		Name:    name,
		Version: version,
		Active:  active,
	}

	return &cft, nil
}
