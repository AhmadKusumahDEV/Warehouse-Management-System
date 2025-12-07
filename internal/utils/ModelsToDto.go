package utils

import (
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/response"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/models"
)

func EmployeeResponse(e *models.Employee) *response.EmployeeResponse {
	return &response.EmployeeResponse{
		UserID:        e.UserID,
		Name:          e.EmployeeName,
		Role:          int(e.IDRole),
		Employee_code: e.EmployeeCode,
		WarehouseCode: e.WarehouseCode,
	}
}

func EmployeeReponses(e []*models.Employee) []*response.EmployeeResponse {
	var res []*response.EmployeeResponse
	for _, v := range e {
		res = append(res, EmployeeResponse(v))
	}
	return res
}

func WarehouseReponse(w *models.Warehouse) *response.WarehouseResponse {
	return &response.WarehouseResponse{
		WarehouseCode:       w.WarehouseCode.String(),
		WarehouseName:       w.WarehouseName,
		LocationDescription: w.LocationDescription,
	}
}

func WarehouseReponses(w []*models.Warehouse) []*response.WarehouseResponse {
	var res []*response.WarehouseResponse
	for _, v := range w {
		res = append(res, WarehouseReponse(v))
	}
	return res
}

func CategeryResponse(c *models.Category) *response.CategoryResponses {
	return &response.CategoryResponses{
		ID:   int(c.ID),
		Name: c.Name,
	}
}

func CategeryReponses(c []*models.Category) []*response.CategoryResponses {
	var res []*response.CategoryResponses
	for _, v := range c {
		res = append(res, CategeryResponse(v))
	}
	return res
}

func SizeResponse(c *models.Size) *response.SizeResponse {
	return &response.SizeResponse{
		ID:   int(c.ID),
		Name: c.Name,
	}
}

func SizeReponses(c []*models.Size) []*response.SizeResponse {
	var res []*response.SizeResponse
	for _, v := range c {
		res = append(res, SizeResponse(v))
	}
	return res
}
