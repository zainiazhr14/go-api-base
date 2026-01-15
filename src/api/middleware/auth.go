package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/domain/constant"
	"github.com/zainiazhr14/go-api/domain/model"
	"github.com/zainiazhr14/go-api/pkg/response"
	pkgToken "github.com/zainiazhr14/go-api/pkg/token"
)

func Auth(cfg *config.Config, db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				response.RespondError(w, http.StatusUnauthorized, "Unauthorized")
				return
			}

			if !strings.HasPrefix(auth, "Bearer ") {
				response.RespondError(w, http.StatusUnauthorized, "Invalid token format")
				return
			}

			rawToken := strings.TrimPrefix(auth, "Bearer ")

			payload, err := pkgToken.VerifyPasetoToken(cfg, rawToken)
			if err != nil {
				response.RespondError(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			userID, err := payload.GetString("user-id")
			if err != nil || userID == "" {
				response.RespondError(w, http.StatusUnauthorized, "Invalid token payload")
				return
			}

			user := new(model.User)
			if err := db.First(user, "id = ?", userID).Error; err != nil {
				response.RespondError(w, http.StatusUnauthorized, "User not found")
				return
			}

			ctx := context.WithValue(r.Context(), constant.UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
