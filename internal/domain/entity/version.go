package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type Version struct {
	Value     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVersion(value string) (*Version, error) {
	version := Version{
		Value: value,
	}

	version.prepare()
	err := version.Validate()

	if err != nil {
		return nil, err
	}

	return &version, nil

}

func (version *Version) prepare() {
	version.CreatedAt = time.Now()
	version.UpdatedAt = time.Now()
}

func (version *Version) Validate() error {
	_, err := govalidator.ValidateStruct(version)

	if err != nil {
		return err
	}

	return nil
}
