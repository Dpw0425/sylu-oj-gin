package utils

import (
	"github.com/golang-jwt/jwt"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/internal/app/entity"
	"sylu-oj-gin/internal/app/schema"
	"time"
)

var jwtKey = []byte(config.CONFIG.Jwt.SigningKey)

func ReleaseToken(eu entity.User) (string, error) {
	// token 结构生成
	claims := &schema.Claims{
		// 使用 ID、Username 作为有效载荷
		UID: int(eu.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + int64(config.CONFIG.Jwt.Expired), // 签名过期时间
			NotBefore: time.Now().Unix() - 1000,                             // 签名生效时间
			Issuer:    config.CONFIG.Jwt.Issuer,                             // 签名发行人
		},
	}

	// 将 Claims 加密存储为 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, *schema.Claims, error) {
	claims := &schema.Claims{}
	// 解码
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
