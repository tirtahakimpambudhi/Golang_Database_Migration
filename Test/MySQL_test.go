package Test

import (
	"context"
	"github.com/stretchr/testify/assert"
	database "go_database_migration/MySQL/db"
	"go_database_migration/MySQL/repository"
	"go_database_migration/helper"
	"go_database_migration/model/domain"
	"testing"
)

func TestGetAll(t *testing.T) {
	MySQLDB := database.NewDB()
	repository := repository.NewBooksRepository()

	tx, err := MySQLDB.Begin()
	ctx := context.Background()
	helper.PanicIFError(err)

	data, total := repository.FindAll(ctx, tx, 10, 0)
	helper.CommitOrRollback(tx)

	assert.Equal(t, len(data), total)

	defer MySQLDB.Close()
}

func TestGetByISBN(t *testing.T) {
	MySQLDB := database.NewDB()
	repository := repository.NewBooksRepository()

	tx, err := MySQLDB.Begin()
	ctx := context.Background()
	helper.PanicIFError(err)

	_, err = repository.FindByISBN(ctx, tx, "123456789")
	//err not nil because the data is empty
	assert.NotEqual(t, nil, err)
	helper.CommitOrRollback(tx)

	defer MySQLDB.Close()
}

func TestCreateBook(t *testing.T) {
	book := domain.Books{
		ID:                1,
		Author:            "Testing",
		Title:             "Testing Books",
		ISBN:              "123456789",
		Publisher:         "Testing Studio",
		Publication_Years: "2024",
		Status_Borrow:     false,
		Description:       "This Books Testing",
	}
	MySQLDB := database.NewDB()
	repository := repository.NewBooksRepository()

	tx, err := MySQLDB.Begin()
	ctx := context.Background()
	helper.PanicIFError(err)

	_, err = repository.FindByISBN(ctx, tx, book.ISBN)
	if err == nil {
		t.Fatalf("Book with ISBN '%v' exist ", book.ISBN)
	}
	err = repository.Create(ctx, tx, book)
	assert.Equal(t, nil, err)
	helper.CommitOrRollback(tx)

	defer MySQLDB.Close()
}

func TestCreateManyBooks(t *testing.T) {
	books := []domain.Books{
		domain.Books{
			Author:            "Testing",
			Title:             "Testing Books",
			ISBN:              "1234567810",
			Publisher:         "Testing Studio",
			Publication_Years: "2024",
			Status_Borrow:     false,
			Description:       "This Books Testing",
		},
		domain.Books{
			Author:            "Testing",
			Title:             "Testing Books Versi 2",
			ISBN:              "11121314151617",
			Publisher:         "Testing Studio",
			Publication_Years: "2024",
			Status_Borrow:     false,
			Description:       "This Books Testing",
		},
	}
	MySQLDB := database.NewDB()
	repository := repository.NewBooksRepository()

	tx, err := MySQLDB.Begin()
	ctx := context.Background()
	helper.PanicIFError(err)

	for _, book := range books {
		_, err = repository.FindByISBN(ctx, tx, book.ISBN)
		if err == nil {
			t.Fatalf("Book with ISBN '%v' exist ", book.ISBN)
			return
		}
	}
	err = repository.CreateMany(ctx, tx, books)
	assert.Equal(t, nil, err)
	helper.CommitOrRollback(tx)

	defer MySQLDB.Close()
}

func TestDeleteManyBooks(t *testing.T) {
	var ISBNs []string
	books := []domain.Books{
		domain.Books{
			Author:            "Testing",
			Title:             "Testing Books",
			ISBN:              "1234567810",
			Publisher:         "Testing Studio",
			Publication_Years: "2024",
			Status_Borrow:     false,
			Description:       "This Books Testing",
		},
		domain.Books{
			Author:            "Testing",
			Title:             "Testing Books Versi 2",
			ISBN:              "11121314151617",
			Publisher:         "Testing Studio",
			Publication_Years: "2024",
			Status_Borrow:     false,
			Description:       "This Books Testing",
		},
	}
	MySQLDB := database.NewDB()
	repository := repository.NewBooksRepository()

	tx, err := MySQLDB.Begin()
	ctx := context.Background()
	helper.PanicIFError(err)

	for _, book := range books {
		_, err = repository.FindByISBN(ctx, tx, book.ISBN)
		if err != nil {
			t.Fatalf("Book with ISBN '%v' not exist ", book.ISBN)
			return
		}
		ISBNs = append(ISBNs, book.ISBN)
	}
	repository.DeleteMany(ctx, tx, ISBNs)
	assert.Equal(t, nil, err)
	helper.CommitOrRollback(tx)

	defer MySQLDB.Close()
}

func TestDeleteBook(t *testing.T) {
	book := domain.Books{
		ID:                1,
		Author:            "Testing",
		Title:             "Testing Books",
		ISBN:              "123456789",
		Publisher:         "Testing Studio",
		Publication_Years: "2024",
		Status_Borrow:     false,
		Description:       "This Books Testing",
	}
	MySQLDB := database.NewDB()
	repository := repository.NewBooksRepository()

	tx, err := MySQLDB.Begin()
	ctx := context.Background()
	helper.PanicIFError(err)

	_, err = repository.FindByISBN(ctx, tx, book.ISBN)
	if err != nil {
		t.Fatalf("Book with ISBN '%v' not exist ", book.ISBN)
	}
	repository.Delete(ctx, tx, book.ISBN)
	assert.Equal(t, nil, err)
	helper.CommitOrRollback(tx)

	defer MySQLDB.Close()
}

func TestUpdateBook(t *testing.T) {
	ISBN := "123456789"
	book := domain.Books{
		ID:                1,
		Author:            "Testing",
		Title:             "Testing Books",
		ISBN:              "123456788",
		Publisher:         "Testing Studio",
		Publication_Years: "2024",
		Status_Borrow:     false,
		Description:       "This Books Testing",
	}
	MySQLDB := database.NewDB()
	repository := repository.NewBooksRepository()

	tx, err := MySQLDB.Begin()
	ctx := context.Background()
	helper.PanicIFError(err)

	_, err = repository.FindByISBN(ctx, tx, ISBN)
	if err != nil {
		t.Fatalf("Book with ISBN '%v' not exist ", ISBN)
	}
	repository.Update(ctx, tx, ISBN, book)
	assert.Equal(t, nil, err)
	helper.CommitOrRollback(tx)

	defer MySQLDB.Close()
}
