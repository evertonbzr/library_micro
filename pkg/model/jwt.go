package model

import "github.com/golang-jwt/jwt/v5"

type ModuleClaims struct {
	jwt.RegisteredClaims
}
