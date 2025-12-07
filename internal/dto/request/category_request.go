package request

type CreateCategory struct {
	Name string `json:"name" binding:"required,min=3,max=23"`
}

type UpdatedCategory struct {
	Name string `json:"name" binding:"required,min=3,max=23"`
}
