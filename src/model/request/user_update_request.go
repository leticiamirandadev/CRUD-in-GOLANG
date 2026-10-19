package request

type UserUpdateRequest struct {
	Name *string `json:"name" binding:"omitempty,min=4,max=100"`
	Age  *int8   `json:"age" binding:"omitempty,min=1,max=140"`
}
