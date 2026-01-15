package usecase

import (
	"errors"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/domain/model"
	"github.com/zainiazhr14/go-api/domain/repository"
)

type AuthUsecase struct {
	cfg            *config.Config
	db             *gorm.DB
	userRepository *repository.UserRepository
}

func NewAuthUsecase(cfg *config.Config, db *gorm.DB) *AuthUsecase {
	userRepository := repository.NewUserRepository(db)

	return &AuthUsecase{
		cfg:            cfg,
		db:             db,
		userRepository: userRepository,
	}
}

func (u *AuthUsecase) LoginByEmail(email string, password string) (*model.User, error) {
	user, err := u.userRepository.FindByEmail(email)

	if err != nil {
		return nil, errors.New("Email not registered")
	}

	return user, nil
}

func (u *AuthUsecase) RegisterUser(user *model.User) error {
	foundUser, err := u.userRepository.FindByEmail(user.Email)

	if err == nil && foundUser.Id != uuid.Nil {
		return errors.New("Email already registered")
	}

	return u.userRepository.Create(user)
}
