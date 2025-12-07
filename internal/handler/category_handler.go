package handler

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/request"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/dto/response"
	"github.com/AhmadKusumahDEV/Warehouse-Management-System/internal/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandlerImpl struct {
	srv service.CategoryServices
}

func NewCategoryHandler(srv service.CategoryServices) CategoryHandler {
	return &CategoryHandlerImpl{srv: srv}
}

// HandlerCreateCategory godoc
// @Summary      Buat Category Baru
// @Description  Membuat category baru dengan data yang diberikan
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Param        Category  body      request.CreateCategory  true  "Data Category Baru"
// @Success      201       {object}  models.Category
// @Failure      400       {object}  response.ApiResponse
// @Failure      408       {object}  response.ApiResponse
// @Failure      500       {object}  response.ApiResponse
// @Failure      504       {object}  response.ApiResponse
// @Router       /category [post]
func (cg *CategoryHandlerImpl) HandlerCreateCategory(c *gin.Context) {
	var category request.CreateCategory

	err := c.ShouldBindJSON(&category)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "format JSON tidak valid atau id warehouse tidak valid",
			Data:    nil,
		})
		return
	}

	err = cg.srv.CreateCategory(c.Request.Context(), &category)

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

func (cg *CategoryHandlerImpl) HandlerDeleteCategory(c *gin.Context) {
	id := c.Param("id")

	val, _ := strconv.Atoi(id)
	err := cg.srv.DeleteCategory(c.Request.Context(), val)

	if err != nil {
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

	c.JSON(http.StatusOK, response.ApiResponse{ // 201
		Status:  http.StatusAccepted,
		Message: "success",
		Data:    nil,
	})
}

func (cg *CategoryHandlerImpl) HandlerGetAllCategory(c *gin.Context) {
	resp, err := cg.srv.GetAllCategory(c.Request.Context())

	if err != nil {
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
		Data:    resp,
	})
}

func (cg *CategoryHandlerImpl) HandlerUpdateCategory(c *gin.Context) {
	id := c.Param("id")
	val, _ := strconv.Atoi(id)

	var category request.UpdatedCategory
	err := c.ShouldBindJSON(&category)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "format JSON tidak valid",
			Data:    nil,
		})
		return
	}

	err = cg.srv.UpdateCategory(c.Request.Context(), &category, val)

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

		c.JSON(http.StatusInternalServerError, response.ApiResponse{ // 500
			Status:  http.StatusInternalServerError,
			Message: "Terjadi kesalahan pada server",
			Data:    nil, // Jangan tampilkan err asli ke user jika 500
		})
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{ // 201
		Status:  http.StatusAccepted,
		Message: "success",
		Data:    nil,
	})
}
