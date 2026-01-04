package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/config"
)

type RouteGroup struct {
	router *mux.Router
	DB     *gorm.DB
	cfg    *config.Config
}

type DBHandler func(w http.ResponseWriter, r *http.Request)

func InitRoutes(root *mux.Router, db *gorm.DB, cfg *config.Config) {
	api := &RouteGroup{router: root, DB: db, cfg: cfg}

	v1 := api.Group("/api/v1")

	RegisterUserRoutes(v1)
	RegisterAuthRoutes(v1)
}

func (g *RouteGroup) Group(prefix string) *RouteGroup {
	return &RouteGroup{
		router: g.router.PathPrefix(prefix).Subrouter(),
		DB:     g.DB,
		cfg:    g.cfg,
	}
}

func (g *RouteGroup) handle(h DBHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r)
	}
}

func (g *RouteGroup) Use(mw func(http.Handler) http.Handler) {
	g.router.Use(mw)
}

func (g *RouteGroup) Get(path string, h DBHandler) {
	g.router.HandleFunc(path, g.handle(h)).Methods("GET")
}

func (g *RouteGroup) Post(path string, h DBHandler) {
	g.router.HandleFunc(path, g.handle(h)).Methods("POST")
}

func (g *RouteGroup) Put(path string, h DBHandler) {
	g.router.HandleFunc(path, g.handle(h)).Methods("PUT")
}

func (g *RouteGroup) Delete(path string, h DBHandler) {
	g.router.HandleFunc(path, g.handle(h)).Methods("DELETE")
}
