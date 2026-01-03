package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/pkg/response"
	"github.com/zainiazhr14/go-api/usecase"
)

type UserHandler struct {
	cfg         *config.Config
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(cfg *config.Config, db *gorm.DB) *UserHandler {
	userUsecase := usecase.NewUserUsecase(cfg, db)

	return &UserHandler{
		cfg:         cfg,
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) LoginWithEmail(w http.ResponseWriter, r *http.Request) {
	user, err := h.userUsecase.LoginByEmail(r, "test@mailinator.com", "123")

	if err != nil { 
		response.RespondError(w, 400, err.Error())
		return
	}

	response.RespondJSON(w, 200, user)
}

