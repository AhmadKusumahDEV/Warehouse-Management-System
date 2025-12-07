package response

type EmployeeResponse struct {
	UserID        string `json:"user_id"`
	Name          string `json:"name"`
	Role          int    `json:"role"`
	Employee_code string `json:"employee_code"`
	WarehouseCode string `json:"warehouse_code"`
}
