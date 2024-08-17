package util

import (
	"errors"
	"time"

	"github.com/evertonbzr/library_micro/cmd/module/user/config"
	"github.com/evertonbzr/library_micro/pkg/model"
	"github.com/golang-jwt/jwt/v5"
)

func HasJwtExpired(token *jwt.Token) error {
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return err
	}

	if exp == nil || exp.Before(time.Now()) {
		return errors.New("TokenHasExpired")
	}

	return nil
}

func GetDurationFromJWT(token *jwt.Token) (time.Duration, error) {
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return 0, err
	}
	return time.Since(exp.Time).Abs(), nil
}

func DecodeJWT(tokenString string) (*jwt.Token, *model.ModuleClaims, error) {
	claims := &model.ModuleClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("UnexpectedSigningMethod")
		}

		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		return nil, nil, errors.New("InvalidToken")
	}

	if err = HasJwtExpired(token); err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}
