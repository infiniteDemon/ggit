package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"service-all/app/init/config"
	"time"
)

var jwtSecret = []byte(config.Config.System.JwtSecret)

type Claims struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(UserName, PassWord string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)

	claims := Claims{
		UserName,
		PassWord,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    config.Config.System.Issuer,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
