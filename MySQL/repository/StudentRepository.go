package repository

import (
	"context"
	"database/sql"
	"go_database_migration/model/domain"
)

type StudentRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx, limit, offset int) (datas []domain.Student, total int)
	FindByNISN(ctx context.Context, tx *sql.Tx, NIS string) (data domain.Student, err error)
	Create(ctx context.Context, tx *sql.Tx, student domain.Student) error
	CreateMany(ctx context.Context, tx *sql.Tx, students []domain.Student) error
	Update(ctx context.Context, tx *sql.Tx, NIS string, student domain.Student)
	Delete(ctx context.Context, tx *sql.Tx, NIS string)
	DeleteMany(ctx context.Context, tx *sql.Tx, NISs []string)
}
