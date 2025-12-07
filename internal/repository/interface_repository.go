package repository

import (
	"context"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/models"
)

type EmployeeRepository interface {
	FindAll(ctx context.Context) ([]*models.Employee, error)
	FindAllByWarehouse(ctx context.Context, warehouseCode string) ([]*models.Employee, error)
	FindById(ctx context.Context, id string) (*models.Employee, error)
	Save(ctx context.Context, employee *models.Employee) error
	Update(ctx context.Context, employee *models.Employee) error
	Delete(ctx context.Context, id string) error
}

type RoleRepository interface {
	FindAll(ctx context.Context) ([]*models.Role, error)
	FindById(ctx context.Context, id int) (*models.Role, error)
	Save(ctx context.Context, role *models.Role) error
	Delete(ctx context.Context, id int) error
}

type ProductRepository interface {
	FindAll(ctx context.Context) ([]*models.Product, error)
	FindById(ctx context.Context, id int) (*models.Product, error)
	Save(ctx context.Context, product *models.Product) error
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id int) error
}

type WarehouseRepository interface {
	FindAll(ctx context.Context) ([]*models.Warehouse, error)
	FindById(ctx context.Context, id string) (*models.Warehouse, error)
	Save(ctx context.Context, warehouse *models.Warehouse) error
	Update(ctx context.Context, warehouse map[string]any, code string) error
	Delete(ctx context.Context, id string) error
}

type CategoryRepository interface {
	FindAll(ctx context.Context) ([]*models.Category, error)
	Save(ctx context.Context, category *models.Category) error
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int) error
}

type SizeRepository interface {
	FindAll(ctx context.Context) ([]*models.Size, error)
	Save(ctx context.Context, size *models.Size) error
	Update(ctx context.Context, size *models.Size) error
	Delete(ctx context.Context, id int) error
}
