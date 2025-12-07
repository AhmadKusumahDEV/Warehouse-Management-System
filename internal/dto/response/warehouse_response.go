package response

type WarehouseResponse struct {
	WarehouseCode       string `json:"warehouse_code"`
	WarehouseName       string `json:"warehouse_name"`
	LocationDescription string `json:"location_description"`
}
