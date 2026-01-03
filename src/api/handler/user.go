package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/domain/model"
	"github.com/zainiazhr14/go-api/pkg/response"
	"github.com/zainiazhr14/go-api/usecase"
)

type UserHandler struct {
	cfg         *config.Config
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	userUsecase := usecase.NewUserUsecase(cfg)

	return &UserHandler{
		cfg:         cfg,
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) LoginWithEmail(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	// Login logic here
}

func GetAllUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var users []model.User
	if err := db.Find(&users).Error; err != nil {
		response.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.RespondJSON(w, http.StatusOK, users)
}

func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user := getUserOr404(db, vars["id"], w, r)
	if user == nil {
		return
	}

	response.RespondJSON(w, http.StatusOK, user)
}

func getUserOr404(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.User {
	var user model.User

	if err := db.First(&user, id).Error; err != nil {
		response.RespondError(w, http.StatusNotFound, err.Error())
		return nil
	}

	return &user
}
