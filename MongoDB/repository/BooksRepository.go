package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_database_migration/helper"
	"go_database_migration/model/domain"
)

type BooksRepositoryMongoDB interface {
	FindAll(ctx context.Context, coll *mongo.Collection, limit, offset int) (datas []domain.BooksDocument, total int)
	FindByISBN(ctx context.Context, coll *mongo.Collection, ISBN string) (data domain.BooksDocument, err error)
	Create(ctx context.Context, coll *mongo.Collection, books domain.BooksCollections) error
	CreateMany(ctx context.Context, coll *mongo.Collection, books []domain.BooksCollections) error
	Update(ctx context.Context, coll *mongo.Collection, ISBN string, books domain.BooksCollections)
	Delete(ctx context.Context, coll *mongo.Collection, ISBN string)
	DeleteMany(ctx context.Context, coll *mongo.Collection, ISBNs []string)
}

type BooksRepositoryImplMongoDB struct {
}

func (b *BooksRepositoryImplMongoDB) Create(ctx context.Context, coll *mongo.Collection, books domain.BooksCollections) error {
	_, err := coll.InsertOne(ctx, books)
	return err
}

func (b *BooksRepositoryImplMongoDB) CreateMany(ctx context.Context, coll *mongo.Collection, books []domain.BooksCollections) error {
	var bookInterfaces []interface{}

	for _, book := range books {
		bookInterfaces = append(bookInterfaces, book)
	}

	_, err := coll.InsertMany(ctx, bookInterfaces)
	return err
}

func (b *BooksRepositoryImplMongoDB) Update(ctx context.Context, coll *mongo.Collection, ISBN string, books domain.BooksCollections) {
	updates := bson.D{{"$set", bson.M{
		"isbn":              books.ISBN,
		"title":             books.Title,
		"author":            books.Author,
		"publisher":         books.Publisher,
		"publication_years": books.Publication_Years,
		"updateAt":          books.UpdateAt,
		"status_borrow":     books.Status_Borrow,
		"description":       books.Description,
	}}}
	filter := bson.D{{"isbn", ISBN}}
	_, err := coll.UpdateOne(ctx, filter, updates)
	helper.PanicIFError(err)
}

func (b *BooksRepositoryImplMongoDB) FindAll(ctx context.Context, coll *mongo.Collection, limit, offset int) (datas []domain.BooksDocument, total int) {
	var books []domain.BooksDocument
	findOpts := options.Find().SetSkip(int64(offset)).SetLimit(int64(limit))
	cursor, err := coll.Find(ctx, bson.M{}, findOpts)
	helper.PanicIFError(err)
	defer cursor.Close(ctx)

	totaldata, err := coll.EstimatedDocumentCount(ctx)
	helper.PanicIFError(err)

	for cursor.Next(ctx) {
		var book domain.BooksDocument
		err = cursor.Decode(&book)
		helper.PanicIFError(err)
		books = append(books, book)
	}

	return books, int(totaldata)
}

func (b *BooksRepositoryImplMongoDB) FindByISBN(ctx context.Context, coll *mongo.Collection, ISBN string) (data domain.BooksDocument, err error) {
	var book domain.BooksDocument
	filter := bson.D{{"isbn", ISBN}}
	cursor := coll.FindOne(ctx, filter)

	if cursor.Err() != nil {
		return book, cursor.Err()
	}
	cursor.Decode(&book)

	return book, nil
}

func (b *BooksRepositoryImplMongoDB) Delete(ctx context.Context, coll *mongo.Collection, ISBN string) {
	filter := bson.D{{"isbn", ISBN}}
	_, err := coll.DeleteOne(ctx, filter)
	helper.PanicIFError(err)
}

func (b *BooksRepositoryImplMongoDB) DeleteMany(ctx context.Context, coll *mongo.Collection, ISBNs []string) {
	filter := bson.D{{"isbn", bson.M{"$in": ISBNs}}}

	_, err := coll.DeleteMany(ctx, filter)
	helper.PanicIFError(err)
}
