package handler

import (
	"github.com/gin-gonic/gin"
)

type EmployeeHandler interface {
	HandlerCreateEmployee(c *gin.Context)
	HandlerGetEmployee(c *gin.Context)
	HandlerDeleteEmployee(c *gin.Context)
	HandlerUpdateEmployee(c *gin.Context)
	HandlerGetAllEmployeeByWarehouse(c *gin.Context)
	HandlerGetAllEmployee(c *gin.Context)
}

type WarehouseHandler interface {
	HandlerGetAllWarehouse(c *gin.Context)
	HandlerCreateWarehouse(c *gin.Context)
	HandlerUpdateWarehouse(c *gin.Context)
	HandlerDeleteWarehouse(c *gin.Context)
}

type CategoryHandler interface {
	HandlerGetAllCategory(c *gin.Context)
	HandlerCreateCategory(c *gin.Context)
	HandlerUpdateCategory(c *gin.Context)
	HandlerDeleteCategory(c *gin.Context)
}
