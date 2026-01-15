package helper

import (
	"context"
	"net/http"

	"github.com/zainiazhr14/go-api/domain/constant"
	"github.com/zainiazhr14/go-api/domain/model"
	"github.com/zainiazhr14/go-api/pkg/response"
)

func GetAuthUser(w http.ResponseWriter, ctx context.Context) *model.User {
	user, ok := ctx.Value(constant.UserContextKey).(*model.User)

	if !ok {
		response.RespondError(w, 401, "Unauthorized")
		return nil
	}

	return user
}
