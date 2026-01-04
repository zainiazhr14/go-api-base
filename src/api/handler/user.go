package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/usecase"
)

type UserHandler struct {
	cfg         *config.Config
	userUsecase *usecase.UserUsecase
}

var validate = validator.New()

func NewUserHandler(cfg *config.Config, db *gorm.DB) *UserHandler {
	userUsecase := usecase.NewUserUsecase(cfg, db)

	return &UserHandler{
		cfg:         cfg,
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {

}
