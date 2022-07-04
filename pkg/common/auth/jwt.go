package auth

import (
	"blog/config"
	"blog/pkg/request"
	"github.com/golang-jwt/jwt"
	"time"
)

var jwtSecret = []byte(config.AuthConfig.JwtSecret)

type claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func CreateToken(param request.UserLogin) (string, error) {
	//设置token有效时间
	expireTime := time.Now().Add(24 * 7 * time.Hour)

	claims := &claims{
		Username: param.UserName,
		Password: param.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Mashiro",
			Subject:   "User Token",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return token, err
}

func ParseToken(t string) (*jwt.Token, *claims, error) {
	claims := &claims{}
	token, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return config.AuthConfig.JwtSecret, nil
	})
	return token, claims, err
}
