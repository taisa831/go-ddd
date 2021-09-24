package response

import "github.com/taisa831/go-ddd/domain/model"

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserListResponse(users []*model.User) []UserResponse {
	rs := []UserResponse{}
	for _, u := range users {
		ur := UserResponse{
			ID:   u.ID,
			Name: u.Name,
		}
		rs = append(rs, ur)
	}
	return rs
}
