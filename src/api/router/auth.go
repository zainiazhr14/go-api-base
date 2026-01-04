package router

import "github.com/zainiazhr14/go-api/api/handler"

func RegisterAuthRoutes(r *RouteGroup) {
	h := handler.NewAuthHandler(r.cfg, r.DB)
	auth := r.Group("/auth")

	auth.Post("/login", h.LoginWithEmail)
	auth.Post("/register", h.RegisterWithEmail)
}
