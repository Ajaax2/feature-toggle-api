package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type FeatureToggle struct {
	ID        string
	Version   *Version
	Name      string
	Active    bool
	CreatedAt time.Time
	UpdateAt  time.Time
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewFeatureToggle(version *Version, name string, active bool) (*FeatureToggle, error) {
	ft := FeatureToggle{
		Version: version,
		Name:    name,
		Active:  active,
	}

	ft.prepare()

	err := ft.Validate()

	if err != nil {
		return nil, err
	}

	return &ft, nil

}

func (ft *FeatureToggle) prepare() {
	ft.ID = uuid.New().String()
	ft.CreatedAt = time.Now()
	ft.UpdateAt = time.Now()
}

func (ft *FeatureToggle) Validate() error {
	_, err := govalidator.ValidateStruct(ft)

	if err != nil {
		return err
	}

	return nil
}
