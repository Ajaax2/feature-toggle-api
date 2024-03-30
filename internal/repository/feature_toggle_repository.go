package repository

import (
	"github.com/Ajaax2/feature-toggle-api/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type FeatureToggleRepository interface {
	CreateFeatureToggle(ft *model.FeatureToggle)
	FindById(id string) (*model.FeatureToggle, error)
	FindAll() ([]*model.FeatureToggle, error)
	FindByVersion(version string) (*model.FeatureToggle, error)
	ExistsByNameAndVersionValue(name string, version string) (*model.FeatureToggle, error)
	DeleteById(id string)
}

type FeatureToggleMongoRepository struct {
	DB *mongo.Database
}

func NewFeatureToggleMongoRepository(db *mongo.Database) *FeatureToggleMongoRepository {
	return &FeatureToggleMongoRepository{DB: db}
}
