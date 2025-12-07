package request

type CreateWarehouse struct {
	WarehouseName       string `json:"warehouse_name" binding:"required, min=3, max=40"`
	LocationDescription string `json:"location_description" binding:"required, min=23, max=60"`
}

type UpdateWarehouse struct {
	WarehouseName       *string `json:"warehouse_name"`
	WarehouseCode       string  `json:"warehouse_code" binding:"required"`
	LocationDescription *string `json:"location_description"`
}
