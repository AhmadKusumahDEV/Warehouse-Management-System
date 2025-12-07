package service

import (
	"context"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/request"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/response"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/models"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/repository"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/utils"
)

type SizeServicesImpl struct {
	repo repository.SizeRepository
}

func NewSizeServices(repo repository.SizeRepository) SizeServices {
	return &SizeServicesImpl{repo: repo}
}

// DeleteSize implements SizeServices.
func (s *SizeServicesImpl) DeleteSize(ctx context.Context, id int) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// GetAllSize implements SizeServices.
func (s *SizeServicesImpl) GetAllSize(ctx context.Context) ([]*response.SizeResponse, error) {
	models, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return utils.SizeReponses(models), nil
}

// SaveSize implements SizeServices.
func (s *SizeServicesImpl) SaveSize(ctx context.Context, size *request.CreateSize) error {
	var sizes models.Size

	sizes.Name = size.Name
	err := s.repo.Save(ctx, &sizes)
	if err != nil {
		return err
	}

	return nil
}

// UpdateSize implements SizeServices.
func (s *SizeServicesImpl) UpdateSize(ctx context.Context, size *request.UpdatedSize, id int) error {
	err := s.repo.Update(ctx, &models.Size{
		ID:   uint(id),
		Name: size.Name,
	})

	if err != nil {
		return err
	}

	return nil
}
