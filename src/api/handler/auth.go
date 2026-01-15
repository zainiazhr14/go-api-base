package handler

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/api/dto"
	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/domain/model"
	"github.com/zainiazhr14/go-api/pkg/response"
	"github.com/zainiazhr14/go-api/pkg/token"
	"github.com/zainiazhr14/go-api/usecase"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
	cfg         *config.Config
}

func NewAuthHandler(cfg *config.Config, db *gorm.DB) *AuthHandler {
	authUsecase := usecase.NewAuthUsecase(cfg, db)

	return &AuthHandler{
		authUsecase: authUsecase,
		cfg:         cfg,
	}
}

func (h *AuthHandler) LoginWithEmail(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginWithEmailReq

	_ = json.NewDecoder(r.Body).Decode(&req)

	if err := validate.Struct(req); err != nil {
		response.RespondError(w, 400, err.Error())
		return
	}

	user, err := h.authUsecase.LoginByEmail(req.Email, req.Password)
	if err != nil {
		response.RespondError(w, 400, err.Error())
		return
	}

	token, err := token.GeneratePasetoToken(h.cfg, user)
	if err != nil {
		response.RespondError(w, 500, "Failed to generate token")
		return
	}

	response.RespondJSON(w, 200, map[string]interface{}{"token": token})
}

func (h *AuthHandler) RegisterWithEmail(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterWithEmailReq

	_ = json.NewDecoder(r.Body).Decode(&req)

	if err := validate.Struct(req); err != nil {
		response.RespondError(w, 400, err.Error())
		return
	}

	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	err := h.authUsecase.RegisterUser(user)

	if err != nil {
		response.RespondError(w, 400, err.Error())
		return
	}

	token, err := token.GeneratePasetoToken(h.cfg, user)
	if err != nil {
		response.RespondError(w, 500, "Failed to generate token")
		return
	}

	response.RespondJSON(w, 200, map[string]interface{}{"access-token": token})
}
