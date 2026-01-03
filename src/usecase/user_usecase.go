package usecase

import (
	"net/http"

	"github.com/zainiazhr14/go-api/config"
)

type UserUsecase struct {
	cfg *config.Config
}

func NewUserUsecase(cfg *config.Config) *UserUsecase {
	return &UserUsecase{
		cfg: cfg,
	}
}

func (u *UserUsecase) LoginByEmail(r *http.Request, email string, password string) (string, error) {
	// Business logic for user login
	return "", nil
}
