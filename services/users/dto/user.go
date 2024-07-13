package dto

import "gobase/models/entities"

type UserDTO struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserDTO(id uint64, username string, email string) *UserDTO {
	return &UserDTO{
		ID:       id,
		Username: username,
		Email:    email,
	}
}

func ToUserDTO(u entities.User) UserDTO {
	return UserDTO{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
	}
}
