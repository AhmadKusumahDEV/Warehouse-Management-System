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

type EmployeeHandlerImpl struct {
	EmployeeService service.EmployeeServices
}

func NewEmployeeHandler(employeeService service.EmployeeServices) EmployeeHandler {
	return &EmployeeHandlerImpl{
		EmployeeService: employeeService,
	}
}

// HandlerCreateEmployee godoc
// @Summary      Buat Employee Baru
// @Description  Membuat employee baru dengan data yang diberikan
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        employee  body      request.CreateEmployee  true  "Data Employee Baru"
// @Success      201       {object}  models.Employee
// @Failure      400       {object}  response.ApiResponse
// @Failure      408       {object}  response.ApiResponse
// @Failure      500       {object}  response.ApiResponse
// @Failure      504       {object}  response.ApiResponse
// @Router       /employees [post]
func (e *EmployeeHandlerImpl) HandlerCreateEmployee(c *gin.Context) {
	var employee request.CreateEmployee

	err := c.ShouldBindJSON(&employee)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "format JSON tidak valid atau id warehouse tidak valid",
			Data:    nil,
		})
		return
	}

	err = e.EmployeeService.CreateEmployee(c.Request.Context(), &employee)

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
			Data:    nil,
		})
		return
	}

	c.JSON(201, response.ApiResponse{
		Status:  201,
		Message: "success",
		Data:    nil,
	})
}

// HandlerDeleteEmployee godoc
// @Summary      Hapus Employee
// @Description  Menghapus data employee berdasarkan ID
// @Tags         employees
// @Produce      json
// @Param        id   path      int  true  "Employee ID"
// @Success      200  {object}  map[string]string
// @Failure      400  {object}  response.ApiResponse  "ID tidak valid"
// @Failure      404  {object}  response.ApiResponse  "Employee tidak ditemukan"
// @Failure      408  {object}  response.ApiResponse  "Request dibatalkan oleh client"
// @Failure      504  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /employees/{id} [delete]
func (e *EmployeeHandlerImpl) HandlerDeleteEmployee(c *gin.Context) {
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

	err = e.EmployeeService.DeleteEmployee(c.Request.Context(), id)

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

// HandlerGetAllEmployee godoc
// @Summary      Get Semua Employee
// @Description  Mengambil daftar semua employee
// @Tags         employees
// @Produce      json
// @Success      200 {array}   models.Employee
// @Failure      500 {object}  response.ApiResponse
// @Failure      504 {object}  response.ApiResponse
// @Router       /employees [get]
func (e *EmployeeHandlerImpl) HandlerGetAllEmployee(c *gin.Context) {
	employee, err := e.EmployeeService.GetAllEmployee(c.Request.Context())

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

	// 4. Sukses (200 OK)
	c.JSON(http.StatusOK, response.ApiResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    employee, // Kirim object/struct data, bukan string
	})
}

// HandlerGetAllEmployeeByWarehouse godoc
// @Summary      Get Employee Berdasarkan Warehouse
// @Description  Mengambil daftar employee yang difilter berdasarkan warehouse_code
// @Tags         employees
// @Produce      json
// @Param        warehouse_code   query     string  true  "Kode Warehouse"
// @Success      200              {array}   models.Employee
// @Failure      400              {object}  response.ApiResponse  "ID tidak valid"
// @Failure      500              {object}  response.ApiResponse
// @Failure      504              {object}  response.ApiResponse
// @Router       /employees/by-warehouse [get]
// @Router       /employees [get]
func (e *EmployeeHandlerImpl) HandlerGetAllEmployeeByWarehouse(c *gin.Context) {
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

	employee, err := e.EmployeeService.GetAllEmployeeByWarehouse(c.Request.Context(), id)

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
		Data:    employee,
	})
}

// HandlerGetEmployee godoc
// @Summary      Get Employee Berdasarkan ID
// @Description  Mengambil satu data employee berdasarkan ID
// @Tags         employees
// @Produce      json
// @Param        id   path      int  true  "Employee ID"
// @Success      200  {object}  models.Employee
// @Failure      400  {object}  response.ApiResponse  "ID tidak valid"
// @Failure      404  {object}  response.ApiResponse  "Employee tidak ditemukan"
// @Failure      500  {object}  response.ApiResponse
// @Router       /employees/{id} [get]
func (e *EmployeeHandlerImpl) HandlerGetEmployee(c *gin.Context) {
	employee, err := e.EmployeeService.GetAllEmployee(c.Request.Context())

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
		Data:    employee,
	})
}

// HandlerUpdateEmployee godoc
// @Summary      Update Employee (Parsial)
// @Description  Memperbarui data employee (bisa sebagian) berdasarkan ID
// @Tags         employees
// @Accept       json
// @Produce      json
// @Param        id        path      int                        true  "Employee ID"
// @Param        employee  body      request.UpdatedEmployee  true  "Data update employee"
// @Success      201       {object}  map[string]string
// @Failure      400       {object}  response.ApiResponse  "ID atau data JSON tidak valid"
// @Failure      404       {object}  response.ApiResponse  "Employee tidak ditemukan"
// @Failure      408       {object}  response.ApiResponse
// @Failure      500       {object}  response.ApiResponse
// @Failure      504       {object}  response.ApiResponse
// @Router       /employees/{id} [patch]
func (e *EmployeeHandlerImpl) HandlerUpdateEmployee(c *gin.Context) {
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

	var employee request.UpdatedEmployee

	err = c.ShouldBindJSON(&employee)

	if err != nil {
		c.JSON(400, response.ApiResponse{
			Status:  400,
			Message: "format JSON tidak valid atau id warehouse tidak valid",
			Data:    "",
		})
		return
	}

	err = e.EmployeeService.UpdateEmployee(c.Request.Context(), id, &employee)

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
