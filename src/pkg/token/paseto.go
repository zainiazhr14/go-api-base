package token

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/zainiazhr14/go-api/config"
	"github.com/zainiazhr14/go-api/domain/model"
)

func GeneratePasetoToken(cfg *config.Config, user *model.User) (string, error) {
	fmt.Println("Loaded paseto key =", cfg.PasetoLocalKey)
	keyStr := cfg.PasetoLocalKey
	key, err := paseto.V4SymmetricKeyFromHex(keyStr)

	if err != nil {
		return "", err
	}

	token := paseto.NewToken()

	token.SetIssuedAt(time.Now())
	token.SetNotBefore(time.Now())
	token.SetExpiration(time.Now().Add(24 * time.Hour))
	token.SetString("user-id", user.Id.String())

	encryptedToken := token.V4Encrypt(key, nil)

	return encryptedToken, nil
}

func VerifyPasetoToken(cfg *config.Config, tokenStr string) (*paseto.Token, error) {
	keyStr := cfg.PasetoLocalKey
	key, err := paseto.V4SymmetricKeyFromHex(keyStr)

	if err != nil {
		return nil, err
	}

	parser := paseto.NewParser()
	token, err := parser.ParseV4Local(key, tokenStr, nil)

	if err != nil {
		return nil, err
	}

	return token, nil
}
