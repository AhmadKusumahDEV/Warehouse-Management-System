package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/models"
)

type CategoryRepositoryImpl struct {
	db *sql.DB
}

// Delete implements CategoryRepository.
func (c *CategoryRepositoryImpl) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM category WHERE id = $1`

	_, err := c.db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("error on delete Category in repository layer", err)
		return err
	}

	return nil
}

// FindAll implements CategoryRepository.
func (c *CategoryRepositoryImpl) FindAll(ctx context.Context) ([]*models.Category, error) {
	query := `SELECT id , name FROM category`

	rows, err := c.db.QueryContext(ctx, query)

	if err != nil {
		log.Println("error on FindAll Category in repository layer", err)
		return nil, err
	}

	defer rows.Close()

	var categorys []*models.Category
	for rows.Next() {
		category := &models.Category{}
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			log.Println("error on FindAll Category in repository layer", err)
			return nil, err
		}

		categorys = append(categorys, category)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return categorys, nil
}

// Save implements CategoryRepository.
func (c *CategoryRepositoryImpl) Save(ctx context.Context, category *models.Category) error {
	query := `INSERT INTO category (name) VALUES ($1)`

	_, err := c.db.ExecContext(ctx, query, category.Name)

	if err != nil {
		log.Println("error on Save Category in repository layer", err)
		return err
	}

	return nil
}

// Update implements CategoryRepository.
func (c *CategoryRepositoryImpl) Update(ctx context.Context, category *models.Category) error {
	query := `UPDATE category SET name = $1 WHERE id = $2`

	_, err := c.db.ExecContext(ctx, query, category.Name, category.ID)

	if err != nil {
		log.Println("error on Update Category in repository layer", err)
		return err
	}

	return nil
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}
