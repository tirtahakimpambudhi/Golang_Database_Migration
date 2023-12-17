package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go_database_migration/config"
	"go_database_migration/helper"
	"go_database_migration/model/domain"
	"log"
)

type BooksRepositoryImpl struct {
}

func (repository *BooksRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, limit, offset int) ([]domain.Books, int) {
	var totalRecords int
	query := fmt.Sprintf("SELECT id,isbn,title,author,status_borrow,publisher,publication_years,description,createdAt,updatedAt FROM %v LIMIT %v OFFSET %v", config.TBNAME, limit, offset)
	countQuery := helper.RemoveLimitOffset(query)

	err := tx.QueryRowContext(ctx, countQuery).Scan(&totalRecords)
	helper.PanicIFError(err)

	books := []domain.Books{}
	rows, err := tx.QueryContext(ctx, query)
	defer rows.Close()
	helper.PanicIFError(err)

	for rows.Next() {
		book := domain.Books{}
		rows.Scan(&book.ID, &book.ISBN, &book.Title, &book.Author, &book.Status_Borrow, &book.Publisher, &book.Publication_Years, &book.Description, &book.CreatedAt, &book.UpdateAt)
		books = append(books, book)
	}

	log.Println(query)
	return books, totalRecords
}

func (repository *BooksRepositoryImpl) FindByISBN(ctx context.Context, tx *sql.Tx, ISBN string) (domain.Books, error) {
	query := fmt.Sprintf("SELECT id,isbn,title,author,status_borrow,publisher,publication_years,description,createdAt,updatedAt FROM %v WHERE isbn = ?", config.TBNAME)
	book := domain.Books{}
	rows, err := tx.QueryContext(ctx, query, ISBN)
	defer rows.Close()
	helper.PanicIFError(err)

	if rows.Next() {
		err := rows.Scan(&book.ID, &book.ISBN, &book.Title, &book.Author, &book.Status_Borrow, &book.Publisher, &book.Publication_Years, &book.Description, &book.CreatedAt, &book.UpdateAt)
		helper.PanicIFError(err)
		return book, nil
	} else {
		return book, errors.New("Books Not Found")
	}
}

func (repository *BooksRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, book domain.Books) error {
	query := fmt.Sprintf("INSERT INTO %v (isbn, title, author, status_borrow, publisher, publication_years, description) VALUES (?,?,?,?,?,?,?)", config.TBNAME)

	log.Println(query)
	_, err := tx.ExecContext(ctx, query, book.ISBN, book.Title, book.Author, book.Status_Borrow, book.Publisher, book.Publication_Years, book.Description)
	return err
}
func (repository *BooksRepositoryImpl) CreateMany(ctx context.Context, tx *sql.Tx, books []domain.Books) error {
	query := fmt.Sprintf("INSERT INTO %v (isbn, title, author, status_borrow, publisher, publication_years, description) VALUES", config.TBNAME)
	values := []interface{}{}

	for i, book := range books {
		if i > 0 {
			query += ","
		}
		query += " (?, ?, ?, ?, ?, ?, ?)"
		values = append(values, book.ISBN, book.Title, book.Author, book.Status_Borrow, book.Publisher, book.Publication_Years, book.Description)
	}

	log.Println(query)
	_, err := tx.ExecContext(ctx, query, values...)
	return err
}

func (repository *BooksRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, ISBN string, book domain.Books) {
	query := fmt.Sprintf("UPDATE %v SET isbn = ?,title = ?,author = ? , status_borrow = ? ,publisher = ?,publication_years = ? , description = ? WHERE isbn = ? ", config.TBNAME)

	log.Println(query)
	_, err := tx.ExecContext(ctx, query, book.ISBN, book.Title, book.Author, book.Status_Borrow, book.Publisher, book.Publication_Years, book.Description, ISBN)
	helper.PanicIFError(err)
}

func (repository *BooksRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, ISBN string) {
	query := fmt.Sprintf("DELETE FROM %v WHERE isbn = ?", config.TBNAME)

	log.Println(query)
	_, err := tx.ExecContext(ctx, query, ISBN)
	helper.PanicIFError(err)
}
func (repository *BooksRepositoryImpl) DeleteMany(ctx context.Context, tx *sql.Tx, ISBNs []string) {
	query := fmt.Sprintf("DELETE FROM %v WHERE isbn IN (?)", config.TBNAME)
	query = helper.InQueryPlaceholders(query, len(ISBNs))

	log.Println(query)
	_, err := tx.ExecContext(ctx, query, helper.SliceToInterface(ISBNs)...)
	helper.PanicIFError(err)
}

func NewBooksRepository() BooksRepository {
	return &BooksRepositoryImpl{}
}
