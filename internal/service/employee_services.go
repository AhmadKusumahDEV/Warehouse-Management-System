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
	"github.com/go-playground/validator/v10"
	uuid "github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

type EmployeeServicesImpl struct {
	EmployeeRepository repository.EmployeeRepository
	validate           *validator.Validate
}

func (s *EmployeeServicesImpl) CreateEmployee(ctx context.Context, req *request.CreateEmployee) error {
	err := s.validate.Struct(req)

	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}

	// 3. Logika Bisnis: Generate UUID untuk user_id
	userID, err := uuid.NewV6()

	if err != nil {
		return errors.New("failed to generate user_id")
	}

	employecode, err := uuid.NewV6()

	if err != nil {
		return errors.New("failed to generate employee_code")
	}

	// 4. Siapkan Model untuk Repository
	employeeModel := &models.Employee{
		UserID:        userID.String(),
		EmployeeName:  req.EmployeeName,
		Password:      string(hashedPassword), // Simpan password yang sudah di-hash
		EmployeeCode:  employecode.String(),
		IDRole:        uint(req.IDRole),
		WarehouseCode: req.WarehouseCode,
	}

	// 5. Panggil Repository
	// (Repository Insert akan mengisi employeeModel.ID)
	err = s.EmployeeRepository.Save(ctx, employeeModel)
	if err != nil {
		return err
	}

	// 6. Kembalikan model yang sudah lengkap
	return nil
}

func (s *EmployeeServicesImpl) UpdateEmployee(ctx context.Context, employee_code string, req *request.UpdatedEmployee) error {
	// 1. Validasi input ID
	if employee_code == "" {
		return errors.New("employee ID is required")
	}

	// 2. READ - Ambil data existing employee
	existingEmployee, err := s.EmployeeRepository.FindById(ctx, employee_code)
	if err != nil {
		return fmt.Errorf("failed to find employee: %w", err)
	}

	// 3. MODIFY - Update field yang dikirim (selective update)
	// Perbaikan: Cek pointer dengan benar untuk optional fields
	if req.EmployeeName != "" {
		existingEmployee.EmployeeName = req.EmployeeName
	}

	// Perbaikan: Cek value, bukan address
	if req.IDRole != 0 {
		existingEmployee.IDRole = uint(req.IDRole)
	}

	// 4. Handle password update dengan validasi
	if req.Password != nil && *req.Password != "" {
		// Validasi password minimal length (contoh: min 8 karakter)
		if len(*req.Password) < 8 {
			return errors.New("password must be at least 8 characters")
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*req.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		existingEmployee.Password = string(hashedPassword)
	}

	// 5. Set updated timestamp jika ada field UpdatedAt
	// existingEmployee.UpdatedAt = time.Now()

	// 6. WRITE - Simpan perubahan ke database
	if err := s.EmployeeRepository.Update(ctx, existingEmployee); err != nil {
		return fmt.Errorf("failed to update employee: %w", err)
	}

	return nil
}

func (s *EmployeeServicesImpl) GetAllEmployee(ctx context.Context) ([]*response.EmployeeResponse, error) {
	models, err := s.EmployeeRepository.FindAll(ctx)

	if err != nil {
		log.Println("error on services layer in method GetAllEmployee when get data from repository", err)
		return nil, err
	}

	resp := utils.EmployeeReponses(models)

	return resp, nil
}

func (s *EmployeeServicesImpl) GetAllEmployeeByWarehouse(ctx context.Context, warehouseCode string) ([]*response.EmployeeResponse, error) {
	models, err := s.EmployeeRepository.FindAllByWarehouse(ctx, warehouseCode)

	if err != nil {
		log.Println("error on services layer in method GetAllEmployeeByWarehouse when get data from repository", err)
		return nil, err
	}

	resp := utils.EmployeeReponses(models)

	return resp, nil
}

func (s *EmployeeServicesImpl) GetEmployeeById(ctx context.Context, id string) (*response.EmployeeResponse, error) {
	models, err := s.EmployeeRepository.FindById(ctx, id)

	if err != nil {
		log.Println("error on services layer in method GetEmployeeById when get data from repository", err)
		return nil, err
	}

	resp := utils.EmployeeResponse(models)

	return resp, nil
}

func (s *EmployeeServicesImpl) DeleteEmployee(ctx context.Context, id string) error {
	// Service tidak perlu logic tambahan, langsung panggil repo
	return s.EmployeeRepository.Delete(ctx, id)
}

func NewEmployeeServices(employeeRepository repository.EmployeeRepository, validate *validator.Validate) EmployeeServices {
	return &EmployeeServicesImpl{
		EmployeeRepository: employeeRepository,
		validate:           validate,
	}
}
