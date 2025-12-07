package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/request"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/response"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/models"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/repository"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/utils"
	uuid "github.com/gofrs/uuid"
)

type WarehouseSErvicesImpl struct {
	repo repository.WarehouseRepository
}

func NewWarehouseServices(repo repository.WarehouseRepository) WarehouseServices {
	return &WarehouseSErvicesImpl{repo: repo}
}

// CreateWarehouse implements WarehouseServices.
func (w *WarehouseSErvicesImpl) CreateWarehouse(ctx context.Context, warehouse *request.CreateWarehouse) error {
	userID, err := uuid.NewV6()

	if err != nil {
		log.Println("error when create userID", err)
		return err
	}

	w.repo.Save(ctx, &models.Warehouse{
		WarehouseName:       warehouse.WarehouseName,
		WarehouseCode:       userID,
		LocationDescription: warehouse.LocationDescription,
	})

	return nil
}

// DeleteWarehouse implements WarehouseServices.
func (w *WarehouseSErvicesImpl) DeleteWarehouse(ctx context.Context, id string) error {
	return w.repo.Delete(ctx, id)
}

// GetAllWarehouse implements WarehouseServices.
func (w *WarehouseSErvicesImpl) GetAllWarehouse(ctx context.Context) ([]*response.WarehouseResponse, error) {
	models, err := w.repo.FindAll(ctx)

	if err != nil {
		log.Println("error on services layer in method GetAllWarehouse when get data from repository", err)
		return nil, err
	}

	return utils.WarehouseReponses(models), nil
}

// GetWarehouseById implements WarehouseServices.
func (w *WarehouseSErvicesImpl) GetWarehouseById(ctx context.Context, id string) (*response.WarehouseResponse, error) {
	models, err := w.repo.FindById(ctx, id)
	if err != nil {
		log.Println("error on services in method GetwarehouseById when get data to repo", err)
		return nil, err
	}

	return utils.WarehouseReponse(models), nil
}

// UpdateWarehouse implements WarehouseServices.
func (w *WarehouseSErvicesImpl) UpdateWarehouse(ctx context.Context, warehouse *request.UpdateWarehouse) error {
	updates := make(map[string]any)

	if warehouse.WarehouseName != nil && *warehouse.WarehouseName != "" {
		// Validasi business logic (contoh: min 3 characters)
		if len(*warehouse.WarehouseName) < 3 {
			return errors.New("warehouse name must be at least 3 characters")
		}
		updates["warehouse_name"] = *warehouse.WarehouseName
	}

	if warehouse.LocationDescription != nil && *warehouse.LocationDescription != "" {
		// Allow empty string untuk clear description
		updates["location_description"] = *warehouse.LocationDescription
	}

	if len(updates) == 0 {
		return errors.New("no fields to update")
	}

	err := w.repo.Update(ctx, updates, warehouse.WarehouseCode)
	if err != nil {
		return fmt.Errorf("failed to update warehouse: %w", err)
	}

	return nil
}
