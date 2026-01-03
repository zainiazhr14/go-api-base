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

func (r *UserRepository) Fin(db *gorm.DB, id any) (*model.User, error) {
	var user model.User
	err := r.Repository.FindById(db, &user, id)
	return &user, err
}
