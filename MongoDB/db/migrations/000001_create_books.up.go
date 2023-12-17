package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_database_migration/config"
	"go_database_migration/model/domain"
	"log"
	"time"
)

func main() {

	mongoURL := "mongodb://localhost:27017"
	ctx := context.Background()
	clientOptions := options.Client().ApplyURI(mongoURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()
	now := time.Now()

	db := client.Database(config.DBNAME)
	indexModel := []mongo.IndexModel{
		mongo.IndexModel{
			Keys:    bson.D{{"isbn", 1}},
			Options: options.Index().SetUnique(true).SetSparse(false),
		},
		mongo.IndexModel{
			Keys:    bson.D{{"title", 1}},
			Options: options.Index().SetSparse(false),
		},
		mongo.IndexModel{
			Keys:    bson.D{{"author", 1}},
			Options: options.Index().SetSparse(false),
		},
		mongo.IndexModel{
			Keys:    bson.D{{"publication_years", 1}},
			Options: options.Index().SetSparse(false).SetMax(5),
		},
	}
	db.CreateCollection(ctx, config.TBNAME)

	db.Collection(config.TBNAME).Indexes().CreateMany(ctx, indexModel)
	db.Collection(config.TBNAME).InsertOne(ctx, domain.BooksCollections{
		ISBN:              "12830-18321-2345",
		Title:             "Migrations Books",
		Author:            "Migrations Author",
		Publisher:         "Migration Studio",
		Publication_Years: "2024",
		Description:       "The Book For Test Migration Succes or not",
		Status_Borrow:     false,
		CreatedAt:         &now,
		UpdateAt:          &now,
	})

	fmt.Println("Migration up successful")
}
