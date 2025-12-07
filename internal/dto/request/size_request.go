package request

type CreateSize struct {
	Name string `json:"name" binding:"required,min=3,max=23"`
}

type UpdatedSize struct {
	Name string `json:"name" binding:"required,min=3,max=23"`
}
