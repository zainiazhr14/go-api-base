package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/domain/model"
)

type UserRepository struct {
	*Repository[model.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByEmail(db *gorm.DB, email string) (*model.User, error) {
	var user model.User

	err := db.Where("email = ?", email).First(&user).Error

	return &user, err
}
