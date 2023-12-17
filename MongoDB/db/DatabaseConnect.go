package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_database_migration/config"
	"time"
)

func NewMongoDB() (*mongo.Database, error) {
	clientOptions := options.Client().SetMaxPoolSize(uint64(100)).SetMinPoolSize(uint64(10)).SetMaxConnIdleTime(5 * time.Minute)
	ctx := context.Background()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(config.DBNAME), nil
}
