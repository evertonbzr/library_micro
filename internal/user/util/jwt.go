package util

import (
	"errors"
	"time"

	"github.com/evertonbzr/library_micro/cmd/module/user/config"
	"github.com/evertonbzr/library_micro/internal/user/model"
	m "github.com/evertonbzr/library_micro/pkg/model"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(user *model.User) (string, error) {
	expTime := time.Now().Add(8760 * time.Hour)

	if user == nil {
		return "", errors.New("user was not found")
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		m.ModuleClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expTime),
			},
		})

	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
