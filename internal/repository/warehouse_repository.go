package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/models"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/utils"
)

type warehouseRepositoryImpl struct {
	db *sql.DB
}

// Constructor untuk membuat instance repository baru
func NewWarehouseRepository(db *sql.DB) WarehouseRepository {
	return &warehouseRepositoryImpl{
		db: db,
	}
}

// Implementasi method FindAll
func (r *warehouseRepositoryImpl) FindAll(ctx context.Context) ([]*models.Warehouse, error) {
	query := `
		SELECT 
			warehouse_name, warehouse_code, location_description 
		FROM 
			warehouse limit 30`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []*models.Warehouse
	for rows.Next() {
		wh := &models.Warehouse{}
		if err := rows.Scan(
			&wh.WarehouseName,
			&wh.WarehouseCode,
			&wh.LocationDescription,
		); err != nil {
			return nil, err
		}
		warehouses = append(warehouses, wh)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return warehouses, nil
}

// Implementasi method FindById
func (r *warehouseRepositoryImpl) FindById(ctx context.Context, id string) (*models.Warehouse, error) {
	query := `
		SELECT 
			warehouse_name, warehouse_code, location_description 
		FROM 
			warehouse 
		WHERE 
			warehouse_code = $1`

	row := r.db.QueryRowContext(ctx, query, id)
	wh := &models.Warehouse{}

	if err := row.Scan(
		&wh.ID,
		&wh.WarehouseName,
		&wh.WarehouseCode,
		&wh.LocationDescription,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("warehouse not found")
		}
		log.Println("error on method FindById in repository layer", err)
		return nil, err
	}

	return wh, nil
}

// Implementasi method Insert
func (r *warehouseRepositoryImpl) Save(ctx context.Context, warehouse *models.Warehouse) error {
	query := `
		INSERT INTO warehouse 
			(warehouse_name, warehouse_code, location_description) 
		VALUES 
			($1, $2, $3) 
		RETURNING 
			id`

	// Menggunakan QueryRowContext untuk mendapatkan ID yang dikembalikan
	err := r.db.QueryRowContext(ctx, query,
		warehouse.WarehouseName,
		warehouse.WarehouseCode,
		warehouse.LocationDescription,
	).Scan(&warehouse.ID)

	if err != nil {
		log.Println("error on method Save in repository layer", err)
		return err
	}

	return nil
}

// Implementasi method Update
func (r *warehouseRepositoryImpl) Update(ctx context.Context, warehouse map[string]any, code string) error {
	allowedColumns := map[string]bool{
		"warehouse_name":       true,
		"location_description": true,
	}

	// Build query using helper
	qb := utils.NewQueryBuilder()

	for column, value := range warehouse {
		// Validate column
		if !allowedColumns[column] {
			return fmt.Errorf("column '%s' is not allowed to be updated", column)
		}

		// Skip nil values
		if value == nil {
			continue
		}

		// Skip empty strings (optional, tergantung requirement)
		if strVal, ok := value.(string); ok && strVal == "" {
			continue
		}

		qb.AddField(column, value)
	}

	// Check if there are updates
	if !qb.HasUpdates() {
		return errors.New("no valid fields to update")
	}

	// Build final query
	query := fmt.Sprintf(
		"UPDATE warehouse SET %s WHERE warehouse_code = $%d",
		qb.BuildSetClause(),
		qb.GetNextPosition(),
	)

	// Append warehouse_code to args
	args := append(qb.GetArgs(), code)

	// Execute
	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update warehouse: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows updated, warehouse not found")
	}

	return nil
}

// Implementasi method Delete
func (r *warehouseRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM warehouse WHERE warehouse_code = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Cek apakah ada baris yang ter-delete
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows deleted, warehouse not found")
	}

	return nil
}
