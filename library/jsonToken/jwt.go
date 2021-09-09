package jsonToken

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// JSON TOKEN CONFIG
type JwtConfig struct {
	jwtSecret []byte
	Expires   time.Duration
	Issuer    string
}

var jwtSecret = []byte("DGene")

type Claims struct {
	User        string `json:"user"`
	PackageName string `json:"PackageName"`
	jwt.StandardClaims
}

func (jwtConfig *JwtConfig) GenerateToken(User, PackageName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(jwtConfig.Expires * time.Hour)

	claims := Claims{
		User,
		PackageName,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    jwtConfig.Issuer,
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
