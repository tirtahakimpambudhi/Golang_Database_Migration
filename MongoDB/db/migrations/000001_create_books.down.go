package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_database_migration/config"
	"log"
)

func main() {

	mongoURL := "mongodb://localhost:27017"
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	client.Database(config.DBNAME).Collection(config.TBNAME).Drop(ctx)

	fmt.Println("Migration down successful")
}
