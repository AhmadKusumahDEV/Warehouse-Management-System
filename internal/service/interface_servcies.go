package service

import (
	"context"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/request"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/response"
)

type EmployeeServices interface {
	GetAllEmployee(ctx context.Context) ([]*response.EmployeeResponse, error)
	GetAllEmployeeByWarehouse(ctx context.Context, warehouseCode string) ([]*response.EmployeeResponse, error)
	GetEmployeeById(ctx context.Context, id string) (*response.EmployeeResponse, error)
	CreateEmployee(ctx context.Context, employee *request.CreateEmployee) error
	UpdateEmployee(ctx context.Context, id string, req *request.UpdatedEmployee) error
	DeleteEmployee(ctx context.Context, id string) error
}

type WarehouseServices interface {
	GetAllWarehouse(ctx context.Context) ([]*response.WarehouseResponse, error)
	GetWarehouseById(ctx context.Context, id string) (*response.WarehouseResponse, error)
	CreateWarehouse(ctx context.Context, warehouse *request.CreateWarehouse) error
	UpdateWarehouse(ctx context.Context, warehouse *request.UpdateWarehouse) error
	DeleteWarehouse(ctx context.Context, id string) error
}

type CategoryServices interface {
	GetAllCategory(ctx context.Context) ([]*response.CategoryResponses, error)
	CreateCategory(ctx context.Context, category *request.CreateCategory) error
	UpdateCategory(ctx context.Context, category *request.UpdatedCategory, id int) error
	DeleteCategory(ctx context.Context, id int) error
}

type SizeServices interface {
	GetAllSize(ctx context.Context) ([]*response.SizeResponse, error)
	SaveSize(ctx context.Context, size *request.CreateSize) error
	UpdateSize(ctx context.Context, size *request.UpdatedSize, id int) error
	DeleteSize(ctx context.Context, id int) error
}
