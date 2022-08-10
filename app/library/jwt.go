package library

import (
	"github.com/all-skeleton/gin-skeleton/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecretApi = []byte(config.Jwt.JwtSecret)

type ApiClaims struct {
	Uid     int    `json:"u"`
	Version string `json:"v"`
	jwt.StandardClaims
}

func GenerateApiToken(claims ApiClaims) (string, error) {
	expireTime := time.Now().Add(48 * time.Hour)
	claims.Version = config.App.Version
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    "gin-skeleton",
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecretApi)
}

func ParseApiToken(token string) (*ApiClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &ApiClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretApi, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*ApiClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
