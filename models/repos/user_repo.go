package repos

import (
	"gobase/models/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	tx *gorm.DB
}

func NewUserRepository(tx *gorm.DB) *UserRepository {
	return &UserRepository{tx: tx}
}

func (repo *UserRepository) Create(user *entities.User) error {
	return repo.tx.Create(user).Error
}

func (repo *UserRepository) FindByUsername(username string) (*entities.User, error) {
	var user *entities.User
	err := repo.tx.Where("username = ?", username).Take(&user).Error
	return user, err
}

func (repo *UserRepository) List() ([]entities.User, error) {
	var users []entities.User
	err := repo.tx.Find(&users).Error
	return users, err
}
