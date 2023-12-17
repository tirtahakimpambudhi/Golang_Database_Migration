package Test

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	db2 "go_database_migration/MongoDB/db"
	"go_database_migration/MongoDB/repository"
	"go_database_migration/model/domain"
	"testing"
	"time"
)

type args struct {
	ctx   context.Context
	coll  *mongo.Collection
	books domain.BooksCollections
}
type test struct {
	name    string
	args    args
	wantErr bool
}

func TestBooksRepositoryImplMongoDB_Create(t *testing.T) {
	db, _ := db2.NewMongoDB()
	now := time.Now()
	book := domain.BooksCollections{
		ISBN:              "12345678",
		Title:             "Books Testing Create",
		Author:            "Testing_Author",
		Publisher:         "Studio Testing",
		Publication_Years: "2024",
		Description:       "No Comment",
		Status_Borrow:     false,
		CreatedAt:         &now,
		UpdateAt:          &now,
	}
	arg := args{
		ctx:   context.Background(),
		coll:  db.Collection("books"),
		books: book,
	}
	tests := []test{
		{
			name:    "Create Book",
			args:    arg,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &repository.BooksRepositoryImplMongoDB{}
			if err := b.Create(tt.args.ctx, tt.args.coll, tt.args.books); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBooksRepositoryImplMongoDB_CreateMany(t *testing.T) {
	db, _ := db2.NewMongoDB()
	now := time.Now()
	book := domain.BooksCollections{
		ISBN:              "12345678910",
		Title:             "Books Testing Create",
		Author:            "Testing_Author",
		Publisher:         "Studio Testing",
		Publication_Years: "2024",
		Description:       "No Comment",
		Status_Borrow:     false,
		CreatedAt:         &now,
		UpdateAt:          &now,
	}
	book1 := domain.BooksCollections{
		ISBN:              "12345678910-11",
		Title:             "Books Testing Create",
		Author:            "Testing_Author",
		Publisher:         "Studio Testing",
		Publication_Years: "2024",
		Description:       "No Comment",
		Status_Borrow:     false,
		CreatedAt:         &now,
		UpdateAt:          &now,
	}
	books := []domain.BooksCollections{
		book,
		book1,
	}
	type args struct {
		ctx   context.Context
		coll  *mongo.Collection
		books []domain.BooksCollections
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Create Many Book", args: args{ctx: context.Background(), coll: db.Collection("books"), books: books}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &repository.BooksRepositoryImplMongoDB{}
			if err := b.CreateMany(tt.args.ctx, tt.args.coll, tt.args.books); (err != nil) != tt.wantErr {
				t.Errorf("CreateMany() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBooksRepositoryImplMongoDB_Delete(t *testing.T) {
	type args struct {
		ctx     context.Context
		coll    *mongo.Collection
		ISBN    string
		wantErr bool
	}
	ctx := context.Background()
	db, _ := db2.NewMongoDB()
	ISBN := "12345678"
	tests := []struct {
		name string
		args args
	}{
		{name: "Delete Books", args: args{ctx: ctx, coll: db.Collection("books"), ISBN: ISBN, wantErr: true}},
		{name: "Delete Books", args: args{ctx: ctx, coll: db.Collection("books"), ISBN: ISBN, wantErr: true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &repository.BooksRepositoryImplMongoDB{}
			if _, err := b.FindByISBN(tt.args.ctx, tt.args.coll, tt.args.ISBN); (err != nil) != tt.args.wantErr {
				t.Errorf("CreateMany() error = %v, wantErr %v", err, tt.args.wantErr)
			}

			b.Delete(tt.args.ctx, tt.args.coll, tt.args.ISBN)
			defer func() {
				err := recover()
				if err != nil {
					t.Log(err)
					return
				}
			}()
		})
	}
}

func TestBooksRepositoryImplMongoDB_DeleteMany(t *testing.T) {
	type args struct {
		ctx   context.Context
		coll  *mongo.Collection
		ISBNs []string
	}
	ctx := context.Background()
	db, _ := db2.NewMongoDB()
	ISBN := []string{"12345678", "123456789"}
	tests := []struct {
		name string
		args args
	}{
		{name: "Delete Many", args: args{ctx: ctx, coll: db.Collection("books"), ISBNs: ISBN}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &repository.BooksRepositoryImplMongoDB{}
			b.DeleteMany(tt.args.ctx, tt.args.coll, tt.args.ISBNs)
		})
	}
}

func TestBooksRepositoryImplMongoDB_FindAll(t *testing.T) {
	type args struct {
		ctx    context.Context
		coll   *mongo.Collection
		limit  int
		offset int
	}
	ctx := context.Background()
	db, _ := db2.NewMongoDB()
	limit := 3
	offset := 0
	tests := []struct {
		name      string
		args      args
		wantTotal int
	}{
		{name: "Find ALL Books", args: args{ctx: ctx, coll: db.Collection("books"), limit: limit, offset: offset}, wantTotal: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &repository.BooksRepositoryImplMongoDB{}
			gotDatas, gotTotal := b.FindAll(tt.args.ctx, tt.args.coll, tt.args.limit, tt.args.offset)
			if gotTotal != tt.wantTotal {
				t.Errorf("FindAll() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			t.Log(len(gotDatas))
		})
	}
}

func TestBooksRepositoryImplMongoDB_FindByISBN(t *testing.T) {
	type args struct {
		ctx  context.Context
		coll *mongo.Collection
		ISBN string
	}
	ctx := context.Background()
	db, _ := db2.NewMongoDB()
	ISBN := "12830-18321-2345"
	tests := []struct {
		name string
		args args

		wantErr bool
	}{
		{name: "Find By ISBN", args: args{ctx: ctx, coll: db.Collection("books"), ISBN: ISBN}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &repository.BooksRepositoryImplMongoDB{}
			gotData, err := b.FindByISBN(tt.args.ctx, tt.args.coll, tt.args.ISBN)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByISBN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(gotData)
			//if !reflect.DeepEqual(gotData, tt.wantData) {
			//	t.Errorf("FindByISBN() gotData = %v, want %v", gotData, tt.wantData)
			//}
		})
	}
}

func TestBooksRepositoryImplMongoDB_Update(t *testing.T) {
	type args struct {
		ctx   context.Context
		coll  *mongo.Collection
		ISBN  string
		books domain.BooksCollections
	}
	ctx := context.Background()
	db, _ := db2.NewMongoDB()
	ISBN := "1234567891011"
	now := time.Now()
	book := domain.BooksCollections{
		ISBN:              "1234-1234-1345",
		Title:             "Books Testing Create",
		Author:            "Testing_Author",
		Publisher:         "Studio Testing",
		Publication_Years: "2024",
		Description:       "No Comment",
		Status_Borrow:     false,
		UpdateAt:          &now,
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Update Books", args: args{ctx: ctx, coll: db.Collection("books"), ISBN: ISBN, books: book}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &repository.BooksRepositoryImplMongoDB{}
			b.Update(tt.args.ctx, tt.args.coll, tt.args.ISBN, tt.args.books)
		})
	}
}
