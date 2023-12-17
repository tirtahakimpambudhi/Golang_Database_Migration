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

type StudentRepositoryImpl struct {
}

func NewStudentRepositoryImpl() StudentRepository {
	return &StudentRepositoryImpl{}
}

func (s *StudentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, limit, offset int) (datas []domain.Student, total int) {

	var totalRecords int
	query := fmt.Sprintf("SELECT id,nis,name,jurusan FROM %v LIMIT %v OFFSET %v", config.TBNAME2, limit, offset)
	countQuery := helper.RemoveLimitOffset(query)

	err := tx.QueryRowContext(ctx, countQuery).Scan(&totalRecords)
	helper.PanicIFError(err)

	students := []domain.Student{}
	rows, err := tx.QueryContext(ctx, query)
	defer rows.Close()
	helper.PanicIFError(err)

	for rows.Next() {
		student := domain.Student{}
		rows.Scan(&student.ID, &student.NIS, &student.Name, &student.Jurusan)
		students = append(students, student)
	}

	log.Println(query)
	return students, totalRecords
}

func (s *StudentRepositoryImpl) FindByNISN(ctx context.Context, tx *sql.Tx, NIS string) (data domain.Student, err error) {
	query := fmt.Sprintf("SELECT id,nis,name,jurusan FROM %v WHERE nis = ?", config.TBNAME2)
	student := domain.Student{}
	rows, err := tx.QueryContext(ctx, query, NIS)
	defer rows.Close()
	helper.PanicIFError(err)

	if rows.Next() {
		err := rows.Scan(&student.ID, &student.NIS, &student.Name, &student.Jurusan)
		helper.PanicIFError(err)
		return student, nil
	} else {
		return student, errors.New("students Not Found")
	}
}

func (s *StudentRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, student domain.Student) error {
	query := fmt.Sprintf("INSERT INTO %v (nis,name,jurusan) VALUES (?,?,?)", config.TBNAME2)

	log.Println(query)
	_, err := tx.ExecContext(ctx, query, student.NIS, student.Name, student.Jurusan)
	return err
}

func (s *StudentRepositoryImpl) CreateMany(ctx context.Context, tx *sql.Tx, students []domain.Student) error {
	query := fmt.Sprintf("INSERT INTO %v (nis,name,jurusan) VALUES", config.TBNAME2)
	values := []interface{}{}

	for i, student := range students {
		if i > 0 {
			query += ","
		}
		query += " (?, ?, ?)"
		values = append(values, student.NIS, student.Name, student.Jurusan)
	}

	log.Println(query)
	_, err := tx.ExecContext(ctx, query, values...)
	return err
}

func (s *StudentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, NIS string, student domain.Student) {
	query := fmt.Sprintf("UPDATE %v SET nis = ?,name = ?,jurusan = ? WHERE isbn = ? ", config.TBNAME2)

	log.Println(query)
	_, err := tx.ExecContext(ctx, query, student.NIS, student.Name, student.Jurusan, NIS)
	helper.PanicIFError(err)
}

func (s *StudentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, NIS string) {

	query := fmt.Sprintf("DELETE FROM %v WHERE nis = ?", config.TBNAME2)

	log.Println(query)
	_, err := tx.ExecContext(ctx, query, NIS)
	helper.PanicIFError(err)
}

func (s *StudentRepositoryImpl) DeleteMany(ctx context.Context, tx *sql.Tx, NISs []string) {
	query := fmt.Sprintf("DELETE FROM %v WHERE nis IN (?)", config.TBNAME2)
	query = helper.InQueryPlaceholders(query, len(NISs))

	log.Println(query)
	_, err := tx.ExecContext(ctx, query, helper.SliceToInterface(NISs)...)
	helper.PanicIFError(err)
}
