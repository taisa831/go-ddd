package request

type UserCreateRequest struct {
	Name string `json:"name" binding:"required"`
}

type UserUpdateRequest struct {
	Name    string `json:"name" binding:"required"`
	Address string `json:"address" binding:"required"`
}
