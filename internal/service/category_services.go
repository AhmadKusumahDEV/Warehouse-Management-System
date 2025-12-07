package service

import (
	"context"
	"log"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/request"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/response"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/models"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/repository"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/utils"
)

type CategoryServicesImpl struct {
	repo repository.CategoryRepository
}

// CreateCategory implements CategoryServices.
func (c *CategoryServicesImpl) CreateCategory(ctx context.Context, category *request.CreateCategory) error {
	model := &models.Category{Name: category.Name}
	err := c.repo.Save(ctx, model)

	if err != nil {
		log.Println("error on layer services in CreateCategory when save category", err)
		return err
	}

	return nil
}

// DeleteCategory implements CategoryServices.
func (c *CategoryServicesImpl) DeleteCategory(ctx context.Context, id int) error {
	err := c.repo.Delete(ctx, id)

	if err != nil {
		log.Println("error on layer services in DeleteCategory when delete category", err)
		return err
	}

	return nil
}

// GetAllCategory implements CategoryServices.
func (c *CategoryServicesImpl) GetAllCategory(ctx context.Context) ([]*response.CategoryResponses, error) {
	models, err := c.repo.FindAll(ctx)

	if err != nil {
		log.Println("error on layer services in GetAllCategory when get all category", err)
		return nil, err
	}

	return utils.CategeryReponses(models), nil
}

// UpdateCategory implements CategoryServices.
func (c *CategoryServicesImpl) UpdateCategory(ctx context.Context, category *request.UpdatedCategory, id int) error {
	model := &models.Category{Name: category.Name, ID: uint(id)}
	err := c.repo.Update(ctx, model)

	if err != nil {
		log.Println("error on layer services in UpdateCategory when update category", err)
		return err
	}

	return nil
}

func NewCategoryServices(repo repository.CategoryRepository) CategoryServices {
	return &CategoryServicesImpl{repo: repo}
}
