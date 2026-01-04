package router

import "github.com/zainiazhr14/go-api/api/middleware"

// import "github.com/zainiazhr14/go-api/api/handler"

func RegisterUserRoutes(r *RouteGroup) {
	// h := handler.NewUserHandler(r.cfg, r.DB)
	users := r.Group("/users")

	users.Use(middleware.Auth(r.cfg))
}
