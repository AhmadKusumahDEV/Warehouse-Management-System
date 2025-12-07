package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/models"
)

type EmployeeRepositoryImpl struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &EmployeeRepositoryImpl{
		db: db,
	}
}

// FindAllByWarehouse implements EmployeeRepository.
func (r *EmployeeRepositoryImpl) FindAll(ctx context.Context) ([]*models.Employee, error) {
	query := `
		SELECT 
			 user_id, employee_name, employee_code, id_role, warehouse_code 
		FROM 
			employee limit 50`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []*models.Employee
	for rows.Next() {
		emp := &models.Employee{}
		if err := rows.Scan(
			&emp.UserID,
			&emp.EmployeeName,
			&emp.EmployeeCode,
			&emp.IDRole,
			&emp.WarehouseCode,
		); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

// Implementasi method FindAllByWarehouse
func (r *EmployeeRepositoryImpl) FindAllByWarehouse(ctx context.Context, warehouseCode string) ([]*models.Employee, error) {
	query := `
		SELECT 
			user_id, employee_name, employee_code, id_role, warehouse_code 
		FROM 
			employee 
		WHERE 
			warehouse_code = $1`

	rows, err := r.db.QueryContext(ctx, query, warehouseCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []*models.Employee
	for rows.Next() {
		emp := &models.Employee{}
		if err := rows.Scan(
			&emp.UserID,
			&emp.EmployeeName,
			&emp.EmployeeCode,
			&emp.IDRole,
			&emp.WarehouseCode,
		); err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

// Implementasi method FindById
func (r *EmployeeRepositoryImpl) FindById(ctx context.Context, employee_code string) (*models.Employee, error) {
	query := `
		SELECT 
			user_id, employee_name, password , employee_code, id_role, warehouse_code 
		FROM 
			employee 
		WHERE 
			employee_code = $1`

	row := r.db.QueryRowContext(ctx, query, employee_code)
	emp := &models.Employee{}

	if err := row.Scan(
		&emp.UserID,
		&emp.EmployeeName,
		&emp.Password,
		&emp.EmployeeCode,
		&emp.IDRole,
		&emp.WarehouseCode,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// Jika tidak ditemukan, kembalikan nil, nil (atau error spesifik jika perlu)
			return nil, errors.New("employee not found")
		}
		return nil, err
	}

	return emp, nil
}

// Implementasi method Insert
func (r *EmployeeRepositoryImpl) Save(ctx context.Context, employee *models.Employee) error {
	query := `
		INSERT INTO employee 
			(user_id, employee_name, password, employee_code, id_role, warehouse_code) 
		VALUES 
			($1, $2, $3, $4, $5, $6) 
		RETURNING 
			id`

	// Menggunakan QueryRowContext karena kita butuh ID yang dikembalikan (RETURNING id)
	err := r.db.QueryRowContext(ctx, query,
		employee.UserID,
		employee.EmployeeName,
		employee.Password,
		employee.EmployeeCode,
		employee.IDRole,
		employee.WarehouseCode,
	).Scan(&employee.ID) // Scan ID baru ke dalam struct employee

	if err != nil {
		return err
	}

	return nil
}

// Implementasi method Update
func (r *EmployeeRepositoryImpl) Update(ctx context.Context, employee *models.Employee) error {
	query := `
		UPDATE employee 
		SET 
			employee_name = $1, 
			password = $2, 
			id_role = $3, 
			warehouse_code = $4 
		WHERE 
			employee_code = $5`

	result, err := r.db.ExecContext(ctx, query,
		employee.EmployeeName,
		employee.Password,
		employee.IDRole,
		employee.WarehouseCode,
		employee.EmployeeCode,
	)

	if err != nil {
		return err
	}

	// Cek apakah ada baris yang ter-update
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no rows updated, employee not found")
	}

	return nil
}

// Implementasi method Delete
func (r *EmployeeRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM employee WHERE employee_code = $1`

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
		return errors.New("no rows deleted, employee not found")
	}

	return nil
}
