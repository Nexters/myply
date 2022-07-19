package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/google/wire"
)

var Set = wire.NewSet(NewMongoDB)

// MongoInstance contains the Mongo client and database objects
type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database // connection
}

func NewMongoDB(config *configs.Config) (*MongoInstance, error) {
	mongoURI := config.MongoURI + "/" + config.MongoDBName
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.MongoTTL)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(config.MongoDBName)

	if err != nil {
		return nil, err
	}

	return &MongoInstance{
		Client: client,
		Db:     db,
	}, nil
}
