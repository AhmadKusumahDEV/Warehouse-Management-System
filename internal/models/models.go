package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// 1. Category
type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

func (Category) TableName() string {
	return "category"
}

// 2. Role
type Role struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	RoleName string `gorm:"not null" json:"role_name"`
}

func (Role) TableName() string {
	return "role"
}

// 3. Size
type Size struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

func (Size) TableName() string {
	return "size"
}

// 4. Status
type Status struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

func (Status) TableName() string {
	return "status"
}

// 5. Warehouse
type Warehouse struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	WarehouseName       string    `gorm:"not null;unique" json:"warehouse_name"`
	WarehouseCode       uuid.UUID `gorm:"not null;unique" json:"warehouse_code"`
	LocationDescription string    `json:"location_description"`

	// Relasi (Has Many)
	Inventories []Inventory `gorm:"foreignKey:CodeWarehouse;references:WarehouseCode" json:"inventories,omitempty"`
}

func (Warehouse) TableName() string {
	return "warehouse"
}

// 6. Employee
type Employee struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	UserID        string `gorm:"unique" json:"user_id"`
	EmployeeName  string `json:"employee_name"`
	Password      string `json:"-"`
	EmployeeCode  string `gorm:"unique" json:"employee_code"`
	IDRole        uint   `gorm:"not null" json:"id_role"`
	WarehouseCode string `gorm:"not null" json:"warehouse_code"`

	// Relasi (Belongs To)
	Role      Role      `gorm:"foreignKey:IDRole" json:"role"`
	Warehouse Warehouse `gorm:"foreignKey:WarehouseCode;references:WarehouseCode" json:"warehouse"`
}

func (Employee) TableName() string {
	return "employee"
}

// 7. Product
type Product struct {
	ID                 uint   `gorm:"primaryKey" json:"id"`
	ProductName        string `gorm:"not null" json:"product_name"`
	Price              int    `gorm:"not null" json:"price"`
	DescriptionProduct string `json:"description_product"` // Diasumsikan bisa null
	ProductCode        string `gorm:"not null;unique" json:"product_code"`
	IDCategory         uint   `gorm:"not null" json:"id_category"`

	// Relasi (Belongs To)
	Category Category `gorm:"foreignKey:IDCategory" json:"category"`

	// Relasi (Has Many)
	ProductDetails []ProductDetail `gorm:"foreignKey:CodeProduct;references:ProductCode" json:"product_details,omitempty"`
}

func (Product) TableName() string {
	return "product"
}

// 8. ProductDetail
type ProductDetail struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	CodeProduct string `gorm:"not null;uniqueIndex:idx_product_size" json:"code_product"`
	IDSize      uint   `gorm:"not null;uniqueIndex:idx_product_size" json:"id_size"`
	Barcode     string `gorm:"not null;unique" json:"barcode"`

	// Relasi (Belongs To)
	Product Product `gorm:"foreignKey:CodeProduct;references:ProductCode" json:"product"`
	Size    Size    `gorm:"foreignKey:IDSize" json:"size"`
}

func (ProductDetail) TableName() string {
	return "product_detail"
}

// 9. Inventory
type Inventory struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Quantity      int    `gorm:"not null;default:0" json:"quantity"`
	CodeProduct   string `gorm:"not null;uniqueIndex:idx_inv_product_size_wh" json:"code_product"`
	IDSize        uint   `gorm:"not null;uniqueIndex:idx_inv_product_size_wh" json:"id_size"`
	CodeWarehouse string `gorm:"not null;uniqueIndex:idx_inv_product_size_wh" json:"code_warehouse"`

	// Relasi (Belongs To)
	Product   Product   `gorm:"foreignKey:CodeProduct;references:ProductCode" json:"product"`
	Warehouse Warehouse `gorm:"foreignKey:CodeWarehouse;references:WarehouseCode" json:"warehouse"`
	Size      Size      `gorm:"foreignKey:IDSize" json:"size"`
}

func (Inventory) TableName() string {
	return "inventory"
}

// 10. Transaction
type Transaction struct {
	ID                    uint      `gorm:"primaryKey" json:"id"`
	CodeTransaksi         string    `gorm:"not null;unique" json:"code_transaksi"`
	OriginEntityName      string    `gorm:"not null" json:"origin_entity_name"`
	DestinationEntityName string    `json:"destination_entity_name"` // Diasumsikan bisa null
	EmployeeCode          string    `gorm:"not null" json:"employee_code"`
	IDStatus              uint      `gorm:"not null" json:"id_status"`
	CreatedAt             time.Time `gorm:"not null" json:"created_at"`
	TipeTransaksi         string    `json:"tipe_transaksi"` // Diasumsikan bisa null

	// Relasi (Belongs To)
	Employee Employee `gorm:"foreignKey:EmployeeCode;references:EmployeeCode" json:"employee"`
	Status   Status   `gorm:"foreignKey:IDStatus" json:"status"`

	// Relasi (Has Many)
	Details []DetailTransaction `gorm:"foreignKey:IDTransaction" json:"details"`
}

func (Transaction) TableName() string {
	return "transactions"
}

// 11. DetailTransaction
type DetailTransaction struct {
	ID              uint `gorm:"primaryKey" json:"id"`
	IDTransaction   uint `gorm:"not null;uniqueIndex:idx_trx_product" json:"id_transaction"`
	IDDetailProduct uint `gorm:"not null;uniqueIndex:idx_trx_product" json:"id_detail_product"`
	Quantity        int  `gorm:"not null" json:"quantity"`
	ScannerQuantity int  `gorm:"not null;default:0" json:"scanner_quantity"`

	// Relasi (Belongs To)
	ProductDetail ProductDetail `gorm:"foreignKey:IDDetailProduct" json:"product_detail"`
	Transaction   Transaction   `gorm:"foreignKey:IDTransaction" json:"-"` // Sembunyikan dari JSON untuk menghindari circular dependency
}

func (DetailTransaction) TableName() string {
	return "detail_transactions"
}
