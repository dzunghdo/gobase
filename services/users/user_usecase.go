package users

import (
	"gobase/db"
	"gobase/models/repos"
	"gobase/services/users/dto"
)

type UserUseCase struct {
	userRepo repos.UserRepository
}

func NewUserUseCase() *UserUseCase {
	userRepo := repos.NewUserRepository(db.GetDB())
	return &UserUseCase{userRepo: *userRepo}
}

func (uc *UserUseCase) List() ([]dto.UserDTO, error) {
	users, err := uc.userRepo.List()
	if err != nil {
		return nil, err
	}
	var userDTOs []dto.UserDTO
	for _, user := range users {
		userDTOs = append(userDTOs, dto.ToUserDTO(user))
	}
	return userDTOs, nil
}
