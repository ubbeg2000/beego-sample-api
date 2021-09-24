package utils

import (
	"errors"
	"fmt"
	"time"

	"github.com/beego/beego/v2/core/config"
	"github.com/dgrijalva/jwt-go"
)

func CreateToken(subject int64) (tok string, err error) {
	secret := config.DefaultString("jwtsecret", "rahasiabetdemigod")
	expTime := config.DefaultInt64("jwtexp", 3600)

	mySigningKey := []byte(secret)

	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprint(subject),
		Audience:  "localhost",
		Issuer:    "sample-api",
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Unix() + expTime,
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(mySigningKey)
}

func ValidateToken(tok string) (claims *jwt.StandardClaims, err error) {
	token, err := jwt.ParseWithClaims(tok, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.DefaultString("jwtsecret", "rahasiabetdemigod")), nil
	})

	if err != nil {
		return nil, err
	}

	if c, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return c, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
