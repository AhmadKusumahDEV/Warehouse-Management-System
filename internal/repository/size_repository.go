package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/models"
)

type SizeRepositoryImpl struct {
	db *sql.DB
}

// Delete implements SizeRepository.
func (s *SizeRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM size WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("error when delete size on repository layer", err)
		return err
	}

	return nil
}

// FindAll implements SizeRepository.
func (s *SizeRepositoryImpl) FindAll(ctx context.Context) ([]*models.Size, error) {
	query := `SELECT id, name FROM size`

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		log.Println("error when find all size on repository layer", err)
		return nil, err
	}

	defer rows.Close()

	var sizes []*models.Size
	for rows.Next() {
		size := &models.Size{}
		if err := rows.Scan(&size.ID, &size.Name); err != nil {
			log.Println("error when find all size on repository layer", err)
			return nil, err
		}

		sizes = append(sizes, size)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return sizes, nil

}

// Save implements SizeRepository.
func (s *SizeRepositoryImpl) Save(ctx context.Context, size *models.Size) error {
	query := `INSERT INTO size (name) VALUES ($1)`

	_, err := s.db.ExecContext(ctx, query, size.Name)
	if err != nil {
		log.Println("error when save size on repository layer", err)
		return err
	}

	return nil
}

// Update implements SizeRepository.
func (s *SizeRepositoryImpl) Update(ctx context.Context, size *models.Size) error {
	query := `UPDATE size SET name = $1 WHERE id = $2`

	_, err := s.db.ExecContext(ctx, query, size.Name, size.ID)

	if err != nil {
		log.Println("error when update size on repository layer", err)
		return err
	}

	return nil
}

func NewSizeRepository(db *sql.DB) SizeRepository {
	return &SizeRepositoryImpl{
		db: db,
	}
}
