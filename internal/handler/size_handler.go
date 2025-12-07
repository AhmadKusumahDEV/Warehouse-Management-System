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

type SizeHandlerImpl struct {
	srv service.SizeServices
}

func NewSizeHandlerImpl(srv service.SizeServices) SizeHandler {
	return &SizeHandlerImpl{srv: srv}
}

// HandlerCreateSize implements SizeHandler.
func (s *SizeHandlerImpl) HandlerCreateSize(c *gin.Context) {
	var size request.CreateSize

	err := c.ShouldBindJSON(&size)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "format JSON tidak valid atau id warehouse tidak valid",
			Data:    nil,
		})
		return
	}

	err = s.srv.SaveSize(c.Request.Context(), &size)

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

// HandlerDeleteSize implements SizeHandler.
func (s *SizeHandlerImpl) HandlerDeleteSize(c *gin.Context) {
	id := c.Param("id")

	val, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "format id tidak valid",
			Data:    nil,
		})
		return
	}

	err = s.srv.DeleteSize(c.Request.Context(), val)

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

// HandlerGetAllSize implements SizeHandler.
func (s *SizeHandlerImpl) HandlerGetAllSize(c *gin.Context) {
	resp, err := s.srv.GetAllSize(c.Request.Context())

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

// HandlerUpdateSize implements SizeHandler.
func (s *SizeHandlerImpl) HandlerUpdateSize(c *gin.Context) {
	id := c.Param("id")
	var req request.UpdatedSize

	conv, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "format id tidak valid",
			Data:    nil,
		})
		return
	}

	err = c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "format JSON tidak valid",
			Data:    nil,
		})
		return
	}

	err = s.srv.UpdateSize(c.Request.Context(), &req, conv)

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
