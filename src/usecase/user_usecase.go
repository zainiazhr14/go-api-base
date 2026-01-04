package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/repository"
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
