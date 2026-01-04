package repository

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/domain/model"
)

type UserRepository struct {
	*Repository[model.User]
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		&Repository[model.User]{db},
	}
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	email = strings.TrimSpace(strings.ToLower(email))

	fmt.Println(email)

	err := r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}
