package request

type CreateEmployee struct {
	EmployeeName  string `json:"employee_name" binding:"required,min=3,max=23"`
	Password      string `json:"password" binding:"required,min=8,max=23"`
	IDRole        int    `json:"id_role" binding:"required"`
	WarehouseCode string `json:"warehouse_code" binding:"required, uuid"`
}

type UpdatedEmployee struct {
	EmployeeName string  `json:"employee_name" binding:"min=3,max=23"`
	Password     *string `json:"password" binding:"min=6"`
	IDRole       int     `json:"id_role"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
