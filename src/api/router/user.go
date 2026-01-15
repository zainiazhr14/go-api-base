package router

import (
	"github.com/zainiazhr14/go-api/api/handler"
	"github.com/zainiazhr14/go-api/api/middleware"
)

func RegisterUserRoutes(r *RouteGroup) {
	h := handler.NewUserHandler(r.cfg, r.DB)
	users := r.Group("/user")

	users.Use(middleware.Auth(r.cfg, r.DB))

	users.Get("/me", h.GetMe)
}
