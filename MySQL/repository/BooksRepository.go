package repository

import (
	"context"
	"database/sql"
	"go_database_migration/model/domain"
)

type BooksRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, limit, offset int) (datas []domain.Books, total int)
	FindByISBN(ctx context.Context, tx *sql.Tx, ISBN string) (data domain.Books, err error)
	Create(ctx context.Context, tx *sql.Tx, books domain.Books) error
	CreateMany(ctx context.Context, tx *sql.Tx, books []domain.Books) error
	Update(ctx context.Context, tx *sql.Tx, ISBN string, books domain.Books)
	Delete(ctx context.Context, tx *sql.Tx, ISBN string)
	DeleteMany(ctx context.Context, tx *sql.Tx, ISBNs []string)
}
