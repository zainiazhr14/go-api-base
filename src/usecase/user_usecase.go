package usecase

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/domain/model"
	"github.com/zainiazhr14/go-api/repository"
)

type UserUsecase struct {
	cfg *config.Config
	db *gorm.DB
	userRepository	*repository.UserRepository
}

func NewUserUsecase(cfg *config.Config, db *gorm.DB) *UserUsecase {
	userRepository := repository.UserRepository{}

	return &UserUsecase{
		cfg: cfg,
		db: db,
		userRepository: &userRepository,
	}
}

func (u *UserUsecase) LoginByEmail(r *http.Request, email string, password string) (*model.User, error) {
	user, err := u.userRepository.FindByEmail(u.db, email)

	if err != nil {
		return nil, err	
	}

	return user, nil
}
