package jwt

import (
	"errors"
	"time"

	jwt1 "github.com/golang-jwt/jwt/v5"
)

var (
	jwtKey = "this is my jwt key"
	jwtAge = time.Hour
)

type myClaim struct {
	jwt1.RegisteredClaims
	ID string `json:"id"`
}

func Set(key string, maxAge time.Duration) {
	jwtKey = key
	if maxAge > 0 {
		jwtAge = maxAge
	}
}

func SignAToken(id string) (token string, err error) {
	t := jwt1.NewWithClaims(jwt1.SigningMethodHS256, &myClaim{
		RegisteredClaims: jwt1.RegisteredClaims{ExpiresAt: jwt1.NewNumericDate(time.Now().Add(jwtAge))},
		ID:               id,
	})
	token, err = t.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) (id string, err error) {
	var t *jwt1.Token
	t, err = jwt1.ParseWithClaims(token, &myClaim{}, func(token *jwt1.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		return "", err
	}
	mc, ok := t.Claims.(*myClaim)
	if !ok {
		return "", errors.New("token parse failed")
	}
	return mc.ID, nil
}
