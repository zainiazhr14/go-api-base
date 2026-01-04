package dto

type LoginWithEmailReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterWithEmailReq struct {
	LoginWithEmailReq
	Name string `json:"name" validate:"required"`
}
