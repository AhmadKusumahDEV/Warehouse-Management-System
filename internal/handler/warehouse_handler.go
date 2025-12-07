package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/request"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/response"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type WarehouseHandlerImpl struct {
	WarehouseService service.WarehouseServices
}

func NewWarehouseHandler(warehouseService service.WarehouseServices) WarehouseHandler {
	return &WarehouseHandlerImpl{WarehouseService: warehouseService}
}

// CreateWarehouse godoc
// @Summary      Buat Warehouse Baru
// @Description  Membuat warehouse baru
// @Tags         warehouses
// @Accept       json
// @Produce      json
// @Param        warehouse  body      request.CreateWarehouse true  "Data Warehouse Baru"
// @Success      201        {object}  models.Warehouse
// @Failure      400        {object}  request.ErrorResponse
// @Failure      408        {object}  request.ErrorResponse
// @Failure      500        {object}  request.ErrorResponse
// @Failure      504        {object}  request.ErrorResponse
// @Router       /warehouses [post]
func (w *WarehouseHandlerImpl) HandlerCreateWarehouse(c *gin.Context) {
	var warehouse request.CreateWarehouse

	err := c.ShouldBindJSON(&warehouse)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "format JSON tidak valid atau id warehouse tidak valid",
			Data:    nil,
		})
		return
	}

	err = w.WarehouseService.CreateWarehouse(c.Request.Context(), &warehouse)

	if err != nil {

		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, response.ApiResponse{ // 504
				Status:  http.StatusGatewayTimeout,
				Message: "Request Timeout",
				Data:    nil,
			})
			return
		}

		if errors.Is(err, context.Canceled) {
			c.JSON(408, response.ApiResponse{
				Status:  408,
				Message: "Request dibatalkan oleh client",
				Data:    nil,
			})
			return
		}

		c.JSON(500, response.ApiResponse{
			Status:  500,
			Message: "error when create warehouse",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusAccepted, response.ApiResponse{
		Status:  201,
		Message: "success",
		Data:    nil,
	})
}

// DeleteWarehouse godoc
// @Summary      Hapus Warehouse
// @Description  Menghapus data warehouse berdasarkan ID
// @Tags         warehouses
// @Produce      json
// @Param        id   path      int  true  "Warehouse ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  response.ApiResponse  "ID tidak valid"
// @Failure      404  {object}  response.ApiResponse  "Warehoouse tidak ditemukan"
// @Failure      408  {object}  response.ApiResponse  "Request dibatalkan oleh client"
// @Failure      504  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /warehouses/{id} [delete]
func (w *WarehouseHandlerImpl) HandlerDeleteWarehouse(c *gin.Context) {
	id := c.Param("id")
	_, err := uuid.FromString(id)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "id tidak valid",
			Data:    nil,
		})
		return
	}

	err = w.WarehouseService.DeleteWarehouse(c.Request.Context(), id)

	if err != nil {

		// Cek A: Apakah error karena Timeout?
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, response.ApiResponse{ // 504
				Status:  http.StatusGatewayTimeout,
				Message: "Request Timeout",
				Data:    nil,
			})
			return
		}

		// Cek B: Apakah error karena Client Cancel (Tutup koneksi)?
		if errors.Is(err, context.Canceled) {
			// 499 (Client Closed Request)
			c.JSON(408, response.ApiResponse{
				Status:  408,
				Message: "Request dibatalkan oleh client",
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, response.ApiResponse{ // 500
			Status:  http.StatusInternalServerError,
			Message: "Terjadi kesalahan pada server",
			Data:    nil, // Jangan tampilkan err asli ke user jika 500
		})
		return
	}

	c.JSON(200, response.ApiResponse{
		Status:  200,
		Message: "success",
		Data:    nil,
	})
}

// GetAllWarehouse godoc
// @Summary      Get Semua Warehouse
// @Description  Mengambil daftar semua warehouse
// @Tags         warehouses
// @Produce      json
// @Success      200 {array}   models.Warehouse
// @Failure      500 {object}  response.ApiResponse
// @Failure      504 {object}  response.ApiResponse
// @Router       /warehouses [get]
func (w *WarehouseHandlerImpl) HandlerGetAllWarehouse(c *gin.Context) {
	warehouse, err := w.WarehouseService.GetAllWarehouse(c.Request.Context())

	if err != nil {

		// A. Timeout / Cancel
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, response.ApiResponse{
				Status:  http.StatusGatewayTimeout,
				Message: "Request Timeout",
				Data:    nil,
			})
			return
		}

		// C. Internal Error (500)
		c.JSON(http.StatusInternalServerError, response.ApiResponse{
			Status:  http.StatusInternalServerError,
			Message: "Terjadi kesalahan internal server",
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    warehouse, // Kirim object/struct data, bukan string
	})
}

// UpdateWarehouse godoc
// @Summary      Update Warehouse (Parsial)
// @Description  Memperbarui data warehouse (bisa sebagian) berdasarkan ID
// @Tags         warehouses
// @Accept       json
// @Produce      json
// @Param        id        path      int                      true  "Warehouse ID"
// @Param        warehouse body      request.UpdateWarehouse  true  "Data update warehouse"
// @Success      201       {object}  map[string]string
// @Failure      400       {object}  response.ApiResponse  "ID atau data JSON tidak valid"
// @Failure      404       {object}  response.ApiResponse  "Employee tidak ditemukan"
// @Failure      408       {object}  response.ApiResponse
// @Failure      500       {object}  response.ApiResponse
// @Failure      504       {object}  response.ApiResponse
// @Router       /warehouses/{id} [patch]
func (w *WarehouseHandlerImpl) HandlerUpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	_, err := uuid.FromString(id)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "id tidak valid",
			Data:    nil,
		})
		return
	}

	var warehouse request.UpdateWarehouse

	err = c.ShouldBindJSON(&warehouse)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "format JSON tidak valid atau id warehouse tidak valid",
			Data:    "",
		})
		return
	}

	err = w.WarehouseService.UpdateWarehouse(c.Request.Context(), &warehouse)

	if err != nil {
		// Cek A: Apakah error karena Timeout?
		if errors.Is(err, context.DeadlineExceeded) {
			c.JSON(http.StatusGatewayTimeout, response.ApiResponse{ // 504
				Status:  http.StatusGatewayTimeout,
				Message: "Request Timeout",
				Data:    nil,
			})
			return
		}

		// Cek B: Apakah error karena Client Cancel (Tutup koneksi)?
		if errors.Is(err, context.Canceled) {
			// 499 (Client Closed Request)
			c.JSON(408, response.ApiResponse{
				Status:  408,
				Message: "Request dibatalkan oleh client",
				Data:    nil,
			})
			return
		}

		c.JSON(500, response.ApiResponse{
			Status:  500,
			Message: "error internal server",
			Data:    "",
		})
		return
	}

	c.JSON(201, response.ApiResponse{
		Status:  201,
		Message: "success",
		Data:    nil,
	})
}
