package middleware

import (
	"context"
	"net/http"

	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/pkg/response"
	"github.com/zainiazhr14/go-api/pkg/token"
)

func Auth(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				response.RespondError(w, 401, "Unauthorized")
				return
			}

			token, err := token.VerifyPasetoToken(cfg, auth)

			if err != nil {
				response.RespondError(w, 401, "Unauthorized")
				return
			}

			userId, err := token.GetString("user-id")
			if err != nil || userId == "" {
				response.RespondError(w, 401, "Unauthorized")
				return
			}

			ctx := context.WithValue(r.Context(), userId, userId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
