package mongodb

import (
	"context"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {

	url := viper.GetString("mongodb.url")
	database := viper.GetString("mongodb.database")

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(url))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(database), nil
}
