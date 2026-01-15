package usecase

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/domain/model"
	"github.com/zainiazhr14/go-api/domain/repository"
)

type UserUsecase struct {
	cfg            *config.Config
	db             *gorm.DB
	userRepository *repository.UserRepository
}

func NewUserUsecase(cfg *config.Config, db *gorm.DB) *UserUsecase {
	userRepository := repository.NewUserRepository(db)

	return &UserUsecase{
		cfg:            cfg,
		db:             db,
		userRepository: userRepository,
	}
}

func (u *UserUsecase) FindById(id string) (*model.User, error) {
	var user *model.User

	err := u.userRepository.FindById(user, id)
	if err != nil {
		return nil, errors.New("User not registered")
	}

	return user, nil
}
