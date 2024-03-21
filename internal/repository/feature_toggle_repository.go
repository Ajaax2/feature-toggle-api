package repository

import (
	"github.com/Ajaax2/feature-toggle-api/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type FeatureToggleRepository interface {
	CreateFeatureToggle(ft *entity.FeatureToggle)
	FindById(id string) (*entity.FeatureToggle, error)
	FindAll() ([]*entity.FeatureToggle, error)
	FindByVersion(version string) (*entity.FeatureToggle, error)
	ExistsByNameAndVersionValue(name string, version string) (*entity.FeatureToggle, error)
	DeleteById(id string)
}

type FeatureToggleRepositoryDb struct {
	DB *mongo.Database
}

func NewFeatureToggleRepository(db *mongo.Database) *FeatureToggleRepositoryDb {
	return &FeatureToggleRepositoryDb{DB: db}
}
